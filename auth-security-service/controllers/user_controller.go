package controllers

import (
    "auth-security-service/models"
    "auth-security-service/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type UserController struct {
    UserService services.UserService
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    var userData struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
        Roles    []uint `json:"roles"`
    }
    if err := c.ShouldBindJSON(&userData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    err := ctrl.UserService.CreateUser(userData.Username, userData.Password, userData.Roles)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (ctrl *UserController) GetUser(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := ctrl.UserService.GetUser(uint(userID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
    // Güncelleme işlemleri
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
    // Silme işlemleri
}
