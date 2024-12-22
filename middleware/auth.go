package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
  }

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {	
		sessionToken, err := ctx.Cookie("session_token")
		if err != nil {
			if ctx.Request.Header.Get("Content-Type") == "application/json"{
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			ctx.Redirect(http.StatusSeeOther,"/login")
			return
			
			
		}
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(sessionToken, claims,func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil || !token.Valid{
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"invalid token"})
			
			return
		}
		ctx.Set("email", claims.Email)
		ctx.Next()			
	
	})
}

