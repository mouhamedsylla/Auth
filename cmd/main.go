package main

import (
	"auth/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	server := server.NewServer()
	server.StartServer("8080")
}
