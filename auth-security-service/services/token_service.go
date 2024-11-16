package services

import (
    "auth-security-service/config"
    "auth-security-service/models"
    "auth-security-service/repositories"
    "auth-security-service/utils"
    "errors"
    "github.com/dgrijalva/jwt-go"
    "time"
)

type TokenService struct {
    TokenRepo repositories.TokenRepository
}

func NewTokenService() TokenService {
    return TokenService{
        TokenRepo: repositories.NewTokenRepository(),
    }
}

func (s *TokenService) RefreshAccessToken(refreshToken string) (string, error) {
    token, err := utils.ValidateJWT(refreshToken)
    if err != nil || !token.Valid {
        return "", errors.New("Invalid refresh token")
    }

    claims := token.Claims.(jwt.MapClaims)
    if claims["token_type"] != "refresh" {
        return "", errors.New("Invalid token type")
    }

    userID := uint(claims["user_id"].(float64))

    newToken, err := utils.GenerateJWTWithClaims(map[string]interface{}{
        "authorized": true,
        "user_id":    userID,
        "exp":        time.Now().Add(time.Hour * 1).Unix(),
        "token_type": "access",
    })
    if err != nil {
        return "", err
    }

    return newToken, nil
}
