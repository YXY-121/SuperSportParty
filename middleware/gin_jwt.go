package middleware

import (
	"apiproject/config"
	"apiproject/pkg"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MyJwt struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func CreateJwt(userId string) string {
	// Create the Claims
	mySigningKey := []byte(config.App.Server.TokenSecretKey)
	claims := MyJwt{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.App.Server.TokenExpireTime) * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		logrus.Errorln("create jwt fail,err is [%v]\n", err)
	}
	return tokenString

}
func ParseJwt(token string, c *gin.Context) error {
	g := pkg.Gin{C: c}

	//有token，解析验证token
	tokenClaims, err := jwt.ParseWithClaims(token, &MyJwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.Server.TokenSecretKey), nil
	})
	claims, ok := tokenClaims.Claims.(*MyJwt)
	if ok && tokenClaims.Valid {
		logrus.Debugf("parese token ok ,the userid is [%v],the token expireTimeAt [%v]\n", claims.UserId, claims.StandardClaims.ExpiresAt)
		return nil
	}

	logrus.Errorf("parse token fail,err is [%v]\n", err)
	g.Response(http.StatusBadRequest, pkg.ErrorTokenInvalid, nil)
	return err

}

func JwtToGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		g := pkg.Gin{C: c}
		token := c.GetHeader(config.App.Server.JwtHeader)
		//请求login 就放行
		if c.Request.URL.String() == "/login" {
			c.Next()
			return
		}

		//没有token，非法访问
		if token == "" {
			logrus.Errorln("token not exist,please login")
			g.Response(http.StatusBadRequest, pkg.ErrorTokenNoExist, nil)
			c.Abort()
			return
		}
		err := ParseJwt(token, c)
		if err != nil {
			c.Abort()
			return
		}
		c.Next()

	}

}
