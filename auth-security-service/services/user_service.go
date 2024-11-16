package services

import (
    "auth-security-service/models"
    "auth-security-service/repositories"
    "auth-security-service/utils"
    "errors"
    "github.com/jinzhu/gorm"
)

type UserService struct {
    UserRepo repositories.UserRepository
    RoleRepo repositories.RoleRepository
}

func NewUserService(db *gorm.DB) UserService {
    return UserService{
        UserRepo: repositories.NewUserRepository(db),
        RoleRepo: repositories.NewRoleRepository(db),
    }
}

func (s *UserService) CreateUser(username, password string, roleIDs []uint) error {
    hashedPassword, err := utils.HashPassword(password)
    if err != nil {
        return err
    }

    roles, err := s.RoleRepo.FindRolesByIDs(roleIDs)
    if err != nil {
        return err
    }

    user := models.User{
        Username:     username,
        PasswordHash: hashedPassword,
        Roles:        roles,
    }

    return s.UserRepo.CreateUser(&user)
}

func (s *UserService) GetUser(userID uint) (*models.User, error) {
    return s.UserRepo.FindByID(userID)
}
