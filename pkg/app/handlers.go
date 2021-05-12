package app

import (
	"log"
	"net/http"

	"github.com/OpitiCalvin/novelsTradeIn/pkg/api"

	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-type", "application/json")

		var newUser api.NewUserRequest

		err := c.ShouldBindJSON(&newUser)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.userService.New(newUser)

		if err != nil {
			log.Printf("Service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "New user record successfully created",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-type", "application/json")

		// query db for data
		users, err := s.userService.GetUsers()

		if err != nil {
			log.Printf("get users handler error: %v", err)
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
