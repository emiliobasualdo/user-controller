package handlers

import (
	"errors"
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
const Realm = "text-realm" // todo change based on env

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
// @ID Get User
// @Produce  json
// @Param   login body  dtos.LoginDto true "user's phone number and the received sms code"
// @Success 200 {object} dtos.TokenDto
// @Failure 401 {object} string "Invalid phone and code combination"
// @Router /auth/login [post]
//https://github.com/appleboy/gin-jwt
func AuthMiddleware() (*jwt.GinJWTMiddleware, error){
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       Realm,
		Key:         []byte("asdfasdfasdfasfdasecret key"), // todo env
		Timeout:     time.Until(time.Now().AddDate(1,0,0)),
		MaxRefresh:  time.Until(time.Now().AddDate(1,1,0)),
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
			return getJwtAccount(login.PhoneNumber, login.Code)
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

func getJwtAccount(phoneNumber string, code string) (interface{}, error) {
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
	return JwtUser{ID: fmt.Sprint(acc.ID), Disabled: !acc.Disabled}, nil
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
	if err := service.SendSmsCode(phoneNumber.PhoneNumber); err != nil {
		Respond(c, http.StatusBadRequest, nil, err)
		return
	}
	Respond(c, http.StatusOK, nil, nil)
}