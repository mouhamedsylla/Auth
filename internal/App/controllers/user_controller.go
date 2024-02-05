package controllers

import (
	"auth/internal/App/models"
	"auth/internal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		name := params["username"][0]
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
		body, err := ioutil.ReadAll(r.Body)
		var message models.Message
		if err != nil {
			message.Error =  "Error reading request body"
			utils.RespondWithJSON(w, message, http.StatusBadRequest)
			return
		}

		var data models.Update
		err = json.Unmarshal(body, &data)
		fmt.Println(data)
		if err != nil {
			message.Error = "Error unmarshaling JSON"
			utils.RespondWithJSON(w, message, http.StatusBadRequest)
			return
		}

		modifer := c.gorm.SetModel(data.ToSelect, data.Value1, models.User{})
		modifer.UpdateField(data.Value2, data.ToUpdate).Update(c.gorm.Db)
		message.Message = fmt.Sprintf("User %s password updated", data.ToSelect)
		utils.RespondWithJSON(w, message, http.StatusOK)
	})
}
