package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sergen-x/sergen-x-api/internal/api"
)

func Start() error {
	router := gin.Default()

	apiGroup := router.Group("/api")
	api.NewHandler(apiGroup)
	router.Run(":8080")

	return nil
}
