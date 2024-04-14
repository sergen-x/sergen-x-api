package api

import (
	"github.com/gin-gonic/gin"
)

type handler struct{}

func NewHandler(router *gin.RouterGroup) *gin.RouterGroup {
	handler := &handler{}

	router.GET("", handler.Start())
	return router
}

func (h *handler) Start() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "OK")
	}
}
