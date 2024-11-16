package controllers

import (
	"github.com/gin-gonic/gin"
)

func EnableMFA(c *gin.Context) {
	userID := c.GetUint("user_id")
	// MFA etkinleştirme işlemleri
}

func VerifyMFA(c *gin.Context) {
	// MFA doğrulama işlemleri
}
