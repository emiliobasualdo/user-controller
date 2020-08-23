package handlers

import (
	"errors"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	. "massimple.com/wallet-controller/internal/dtos"
	"massimple.com/wallet-controller/internal/models"
	"massimple.com/wallet-controller/internal/service"
	. "massimple.com/wallet-controller/internal/webapp/utils"
	"net/http"
	"time"
)

const IdentityKey = "acc_id"

type JwtUser struct {
	ID  string
}

func (jw JwtUser) getId() models.ID {
	return models.ID(jw.ID)
}

func (jw JwtUser) getIdString() string {
	return jw.ID
}

type JwtUserInterface interface {
	getId()	models.ID
	getIdString() string
}

func AuthMiddlewareWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get(IdentityKey)
		if !exists {
			c.String(http.StatusBadRequest, "broken jwt")
			return
		}
		c.Keys[IdentityKey] = user
		c.Next()
	}
}

// @Summary Generate a jwt
// @Description Returns a jwt token to use as authentication
// @query Get User
// @Produce  json
// @Param   login body  dtos.LoginDto true "user's phone number and the received sms code"
// @Success 200 {object} dtos.TokenDto
// @Failure 401 "Invalid phone and code combination"
// @Router /auth/login [post]
//https://github.com/appleboy/gin-jwt
func AuthMiddleware() (*jwt.GinJWTMiddleware, error){
	jwtPrivate := viper.GetString("jwt.privateKeyFile")
	jwtPublic := viper.GetString("jwt.publicKeyFile")
	jwtRealm  := viper.GetString("jwt.realm")
	jwtSigningAlgorithm := viper.GetString("jwt.signingAlgorithm")
	return jwt.New(&jwt.GinJWTMiddleware{
		SigningAlgorithm: 	jwtSigningAlgorithm,
		PrivKeyFile: 		jwtPrivate,
		PubKeyFile: 		jwtPublic,
		Realm:       		jwtRealm,
		Timeout:     		time.Until(time.Now().AddDate(1,0,0)),
		MaxRefresh:	 		time.Until(time.Now().AddDate(1,1,0)),
		IdentityKey: 		IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(JwtUser); ok {
				return jwt.MapClaims{
					IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &JwtUser{
				ID: claims[IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login LoginDto
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			return getJwtAccount(login.PhoneNumber, login.Code)
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*JwtUser); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value.
		//This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}

func getJwtAccount(phoneNumber models.PhoneNumber, code string) (interface{}, error) {
	verified, err := service.CheckCode(phoneNumber, code)
	if !verified {
		return nil, errors.New("invalid code")
	}
	if err != nil {
		return nil, err
	}
	acc, err := service.GetAccount(phoneNumber)
	if err != nil {
		return nil, err
	}
	return JwtUser{ID: acc.ID.String()}, nil
}

type phoneNumberDto struct {
	PhoneNumber models.PhoneNumber `json:"phoneNumber" binding:"required" example:"005491133071114"`
}


// @Summary SMS auth
// @Description Sends an sms to the specified phoneNumber
// @query Get User
// @Param   login body  handlers.phoneNumberDto true "user's phone number"
// @Success 200
// @Failure 400 "Invalid phone number"
// @Failure 500 "Something went wrong"
// @Router /auth/sms-code [post]
//https://github.com/appleboy/gin-jwt
func SendSmsHandler(c *gin.Context)  {
	var dto phoneNumberDto
	if err := c.BindJSON(&dto); err != nil {
		Respond(c, http.StatusBadRequest, "You must provide a phone number", nil)
		return
	}
	if err := service.SendSmsCode(dto.PhoneNumber); err != nil {
		Respond(c, http.StatusBadRequest, nil, err)
		return
	}
	Respond(c, http.StatusOK, nil, nil)
}