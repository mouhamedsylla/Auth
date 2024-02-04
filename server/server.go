package server

import (
	"auth/internal/App/controllers"
	"auth/internal/App/services"
	"auth/server/middleware"
	"auth/server/router"
	"fmt"
	"net/http"
)

type Server struct {
	Router *router.Router
}

func NewServer() *Server {
	return &Server{
		Router: router.NewRouter(),
	}
}

func (s *Server) ConfigureRoutes() {
	var auth services.AuthService

	s.Router.Use(middleware.CORSMiddleware)
	controller := controllers.NewAuthController(auth)
	s.Router.Method(http.MethodPost).Handler("/register", controller.Register())
}

func (s *Server) StartServer(port string) error {
	s.ConfigureRoutes()
	fmt.Printf("http://localhost:%v", port)
	return http.ListenAndServe(":"+port, s.Router)
}
