package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"w4s/models"
)

// Login is the signIn method
func (controller *loginController) Login(c *gin.Context) string{
	db:= c.MustGet("db").(*gorm.DB)
	user := models.User{}
	err:=c.ShouldBindJSON(&user)
	if err!=nil{
		defer db.Close()
		return ""
	}
	isAuthenticated:=controller.loginService.Login(user.Email,user.Password)
	if isAuthenticated{
		defer db.Close()
		return controller.jwtService.GerenateToken(user.Email,true)
	}
	defer db.Close()
	return ""
}
