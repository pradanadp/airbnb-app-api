package middlewares

import (
	"be-api/app/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.SECRET_JWT),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId int) (string, error){
	claims := jwt.MapClaims{}
	claims["authorized"] =true
	claims["userId"] = userId
	claims["exp"]= time.Now().Add(time.Hour*24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtracTokenUserId(e echo.Context) (int){
	user := e.Get("user").(*jwt.Token)
	if user.Valid{
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}