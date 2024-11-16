package middleware

import (
    "com.digitalidentity.datamanagementservice/config"
    "com.digitalidentity.datamanagementservice/utils"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header required")
            c.Abort()
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header format must be Bearer {token}")
            c.Abort()
            return
        }

        tokenString := parts[1]

        claims := &utils.Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(config.AppConfig.JWTSecret), nil
        })

        if err != nil || !token.Valid {
            utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
            c.Abort()
            return
        }

        c.Set("user_id", claims.UserID)
        c.Next()
    }
}
