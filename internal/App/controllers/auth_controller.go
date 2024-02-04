package controllers

import (
	"auth/internal/App/models"
	"auth/internal/App/services"
	"auth/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {

}

func (c *AuthController) Register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := io.ReadAll(r.Body)
		var user models.User
		var message models.Message
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = json.Unmarshal(resp, &user); err != nil {
			message := models.Message{
				Error: "Bad Request" + err.Error(),
			}
			utils.RespondWithJSON(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		gorm := utils.OrmInit("auth.db")
		if err := gorm.Insert(user); err != nil {
			if shouldReturn := CheckExist(err, message, w); shouldReturn {
				return
			}
			message := models.Message{
				Error: fmt.Sprintf("Status Internal Server Error: %v", err.Error()),
			}
			utils.RespondWithJSON(w, message, http.StatusInternalServerError)
			return
		}
		message.Message = fmt.Sprintf("User %s registered successfully", user.Username)
		utils.RespondWithJSON(w, message, http.StatusOK)
	})
}

func CheckExist(err error, message models.Message, w http.ResponseWriter) bool {
	if err.Error() == "UNIQUE constraint failed: User.Email" {
		message.Error = "Email already in use"
		utils.RespondWithJSON(w, message, http.StatusBadRequest)
		return true
	} else if err.Error() == "UNIQUE constraint failed: User.Password" {
		message.Error = "Password already in use"
		utils.RespondWithJSON(w, message, http.StatusBadRequest)
		return true
	}
	return false
}
