package app

import (
	"log"

	"github.com/OpitiCalvin/novelsTradeIn/pkg/api"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	userService api.UserService
	bookService api.BookService
}

func NewServer(router *gin.Engine, userService api.UserService, bookService api.BookService) *Server {
	return &Server{
		router:      router,
		userService: userService,
		bookService: bookService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on routerL %v", err)
		return err
	}

	return nil
}
