package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    "com.digitalidentity.identityservice/config"
)

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(userID uint) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.AppConfig.JWTSecret), nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    return claims, nil
}
