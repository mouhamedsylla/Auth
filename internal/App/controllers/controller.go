package controllers

import (
	"auth/internal/App/services"
	"auth/internal/utils"
	"auth/orm"
)

type AuthController struct {
	authService services.AuthService
	gorm        *orm.ORM
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
		gorm:        utils.OrmInit("auth.db"),
	}
}
