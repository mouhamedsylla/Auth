package models

import "auth/orm"

type Message struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

// User représente un utilisateur dans l'application.
type User struct {
	orm.Model
	Username string `orm-go:"NOT NULL"`
	Email    string `orm-go:"NOT NULL UNIQUE"`
	Password string `orm-go:"NOT NULL UNIQUE"`
}

func NewUser(name, email, password string) User{
	return User{
		Username: name,
		Email: email,
		Password: password,
	}
}

type Update struct {
	ToSelect string
	Value1     string
	ToUpdate string
	Value2    string
}
