package gpIssuer

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"os"
	"time"
)

var api = struct {
	clientId string
	secret string
	productNumber int
	afinityGroup int
	branch int
	scope string
	tokenUrl string
	baseUrl string
}{}

var productBaseUrl string

var client *http.Client
func GPInit() {
	// gp demands static ip for communication.
	// we hosted a proxy for dev
	// todo api debería ser agnóstica de su ip
	if value, _ := os.LookupEnv("ENV"); value != "PROD" {
		if !viper.IsSet("httpProxy.url") {
			panic("Proxy is not set")
		}
		proxyUrl, err := url.Parse(viper.GetString("httpProxy.url"))
		if err != nil {
			panic(err)
		}
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	} else {
		client = &http.Client{}
	}
	// todo GP is available only between 8am and 10pm. podría solucionarlo con un thread sleeping
	// we need to set the api auth variables
	api.clientId = 		viper.GetString("gpIssuer.clientId")
	api.secret = 		viper.GetString("gpIssuer.secret")
	api.productNumber = viper.GetInt("gpIssuer.productNumber")
	api.afinityGroup = 	viper.GetInt("gpIssuer.afinityGroup")
	api.branch = 		viper.GetInt("gpIssuer.branch")
	api.scope = 		viper.GetString("gpIssuer.scope")
	api.tokenUrl = 		viper.GetString("gpIssuer.tokenUrl")
	api.baseUrl = 		viper.GetString("gpIssuer.baseUrl")
	// many request use default constat fields that depend on the env variables picked up by viper
	// so we fill the values now
	fillDefaults()
	// we generate de base url
	productBaseUrl = fmt.Sprintf("%s/Productos/%d", api.baseUrl, api.productNumber)
	getToken()
}

func getToken() {
	// setup login variables
	values := url.Values{}
	values.Set("client_id", api.clientId)
	values.Set("client_secret", api.secret)
	values.Set("grant_type","client_credentials")
	values.Set("scope", api.scope)
	// get token
	resp, err := client.PostForm(api.tokenUrl, values)
	if err != nil || resp.StatusCode == 404{
		panic(err)
	}
	var authData = struct {
		AccessToken string `json:"Access_token"`
		Expires   	int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&authData)
	if err != nil {
		panic(err)
	}
	// we copy the data to a global structure
	auth.AccessToken 	= authData.AccessToken
	auth.TokenType 		= authData.TokenType
	auth.ExpiresAt		= time.Now().Add(time.Duration(authData.Expires) * time.Second)
}
