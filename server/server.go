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
	s.Router.Method(http.MethodGet).Handler("/getuser", controller.GetUser())
	s.Router.Method(http.MethodDelete).Handler("/delete", controller.DeleteUser())
	s.Router.Method(http.MethodPost).Handler("/update", controller.UpdateUser())
}

func (s *Server) StartServer(port string) error {
	s.ConfigureRoutes()
	fmt.Printf("http://localhost:%v", port)
	return http.ListenAndServe(":"+port, s.Router)
}
