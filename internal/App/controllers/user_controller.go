package controllers

import (
	"auth/internal/App/models"
	"auth/internal/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (c *AuthController) GetUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		var id int
		var err error
		if len(params["id"]) != 0 {
			if id, err = strconv.Atoi(params["id"][0]); err != nil {
				message := models.Message{
					Error: "Bad Request",
				}
				utils.RespondWithJSON(w, message, http.StatusBadRequest)
				return
			}
		}

		if len(params["id"]) != 0 {
			c.gorm.Custom.Where("id", id)
		}

		user := c.gorm.Scan(models.User{}, "Id", "Username", "Email", "CreatedAt").([]struct {
			Id        int
			Username  string
			Email     string
			CreatedAt time.Time
		})

		c.gorm.Custom.Clear()

		if len(user) == 0 {
			message := models.Message{
				Error: "User not found",
			}
			utils.RespondWithJSON(w, message, http.StatusNotFound)
			return
		}

		utils.RespondWithJSON(w, user, http.StatusOK)
	})
}

func (c *AuthController) DeleteUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		if len(params["username"]) == 0 {
				message := models.Message{
					Error: "Bad Request",
				}
				utils.RespondWithJSON(w, message, http.StatusBadRequest)
				return	
		}
		name:= params["username"][0]
		if err := c.gorm.Delete(models.User{}, "Username", name); err != nil {
			message := models.Message{
				Error: "Bad request" + err.Error(),
			}
			utils.RespondWithJSON(w, message, http.StatusBadRequest)
		}

		var message models.Message
		message.Message = fmt.Sprintf("User %s deleted", name)
		utils.RespondWithJSON(w, message, http.StatusOK)
	})
}

func (c *AuthController) UpdateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}
