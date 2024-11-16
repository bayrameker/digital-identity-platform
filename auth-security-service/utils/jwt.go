package utils

import (
    "auth-security-service/config"
    "auth-security-service/models"
    "github.com/dgrijalva/jwt-go"
    "time"
)

func GenerateJWT(user models.User) (string, error) {
    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["user_id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.AppConfig.JWTSecretKey))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.AppConfig.JWTSecretKey), nil
    })
    return token, err
}
