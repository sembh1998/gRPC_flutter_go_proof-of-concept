package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/gRCP_flutter_go_proof-of-concept/server/src/bootstrap"
)

type Server struct {
}

func NewServer() (bootstrap.Bootstrap, error) {
	return &Server{}, nil
}

func (s *Server) Initialice() error {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := router.Group("/api")
	version1 := api.Group("/v1")
	storesGroup := version1.Group("/stores")

	storesGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + os.Getenv("SERVER_PORT"))
	return nil
}
