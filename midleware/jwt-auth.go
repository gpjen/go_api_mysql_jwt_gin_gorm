package midleware

import (
	"go_api_mysql_jwt_gin_gorm/helper"
	"go_api_mysql_jwt_gin_gorm/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

func AuthorixeJWT(jwtService service.JWTservice) gin.HandlerFunc {
	return func(c *gin.Context) {
		autHeader := c.GetHeader("Authorization")
		if autHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, helper.ResponseFail("Failed to process request", "No auth token found"))
			return
		}
		token, err := jwtService.ValidateToken(autHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id] :", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ResponseFail("Token is not valid", err.Error()))
		}
	}
}
