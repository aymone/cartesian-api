package router

import (
	"github.com/aymone/cartesian-api/handlers"
	"github.com/gin-gonic/gin"
)

// URL base
const URL string = "/api/points"

// Init router
func Init(h handlers.Handler) *gin.Engine {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	// Define endpoints
	router.GET("/api/points", h.GetPoints)

	return router
}
