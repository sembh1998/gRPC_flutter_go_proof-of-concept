package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/bootstrap"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/util"
)

type Server struct {
}

func NewServer(config util.Config) (bootstrap.Bootstrap, error) {
	return &Server{}, nil
}

func (s *Server) Initialice(config util.Config) error {
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

	router.Run(config.HTTPServerAddress)
	return nil
}
