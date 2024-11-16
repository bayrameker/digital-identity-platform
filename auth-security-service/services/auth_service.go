package services

import (
    "auth-security-service/models"
    "auth-security-service/repositories"
    "auth-security-service/utils"
    "errors"
    "github.com/jinzhu/gorm"
)

type AuthService struct {
    UserRepo repositories.UserRepository
}

func NewAuthService(db *gorm.DB) AuthService {
    return AuthService{
        UserRepo: repositories.NewUserRepository(db),
    }
}

func (s *AuthService) AuthenticateUser(username, password string) (string, error) {
    user, err := s.UserRepo.FindByUsername(username)
    if err != nil {
        return "", err
    }

    if !utils.CheckPasswordHash(password, user.PasswordHash) {
        return "", errors.New("invalid credentials")
    }

    token, err := utils.GenerateJWT(user)
    if err != nil {
        return "", err
    }

    return token, nil
}
