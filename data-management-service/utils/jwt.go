package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

type Claims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(userID string, secret string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
