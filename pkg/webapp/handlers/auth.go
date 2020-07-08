package handlers

import (
	"fmt"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"massimple.com/wallet-controller/pkg/service"
	. "massimple.com/wallet-controller/pkg/webapp/dtos"
	. "massimple.com/wallet-controller/pkg/webapp/utils"
	"net/http"
	"strconv"
	"time"
)

const IdentityKey = "acc_id"

type JwtUser struct {
	ID  string
	Disabled bool
}

type JwtUserInterface interface {
	getId()	uint
}

func (jw JwtUser) getId() uint {
	id, err := strconv.ParseUint(jw.ID,10, 64)
	if err != nil {
		return 0
	}
	return uint(id)
}

// @Summary Generate a jwt
// @Description Returns a jwt token to use as authentication
// @ID Get User
// @Produce  json
// @Param   login body  dtos.LoginDto true "user's phone number and the received sms code"
// @Success 200 {object} dtos.TokenDto
// @Failure 400 {object} string "The phone number does not match the code"
// @Failure 404 {object} string "" "id does not exist"
// @Router /auth/login [post]
//https://github.com/appleboy/gin-jwt
func AuthMiddleware() (*jwt.GinJWTMiddleware, error){
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("asdfasdfasdfasfdasecret key"), // todo env
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
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
			return getJwtAccount(login.PhoneNumber)
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*JwtUser); ok && !v.Disabled {
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

func getJwtAccount(phoneNumber string) (interface{}, error) {
	acc, err := service.GetAccount(phoneNumber)
	if err != nil {
		return nil, err
	}
	return JwtUser{ID: fmt.Sprint(acc.ID), Disabled: !acc.DisabledSince.IsZero()}, nil
}

// @Summary SMS auth
// @Description Sends an sms to the specified phonenumber
// @ID Get User
// @Param   login body  dtos.PhoneNumberDto true "user's phone number"
// @Success 200
// @Failure 400 {object} string "Something went wrong"
// @Router /auth/sms-code [post]
//https://github.com/appleboy/gin-jwt
func SendSmsHandler(c *gin.Context)  {
	var phoneNumber PhoneNumberDto
	if err := c.BindJSON(&phoneNumber); err != nil {
		Respond(c, http.StatusBadRequest, "You must provide a phone number", nil)
		return
	}
	if err := service.SendSms(phoneNumber.PhoneNumber); err != nil {
		Respond(c, http.StatusBadRequest, nil, err)
		return
	}
	Respond(c, http.StatusOK, nil, nil)
}