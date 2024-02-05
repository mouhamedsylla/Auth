package main

import (
	"auth/internal/App/services/jwt"
	"fmt"
)

func main() {
	var jwt jwt.JWT

	token := jwt.GenerateToken()

	fmt.Println(token)
}