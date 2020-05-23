package handlers

import (
	"net/http"

	"github.com/aymone/cartesian-api/services"
	"github.com/aymone/cartesian-api/utils"
	"github.com/gin-gonic/gin"
)

// Handlers will handle requests from gin to services
type Handler interface {
	GetPoints(c *gin.Context)
}

type handler struct {
	pointsService services.PointsService
}

func New(s services.PointsService) *handler {
	return &handler{
		pointsService: s,
	}
}

func (h *handler) GetPoints(c *gin.Context) {
	q := c.Request.URL.Query()

	ok, x := utils.ParseInt(q.Get("x"))
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMessage{
			Status: http.StatusBadRequest,
			Err:    "query parameter 'x' is required",
		})
		return
	}

	ok, y := utils.ParseInt(q.Get("y"))
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMessage{
			Status: http.StatusBadRequest,
			Err:    "query parameter 'y' is required",
		})
		return
	}

	ok, distance := utils.ParseInt(q.Get("distance"))
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMessage{
			Status: http.StatusBadRequest,
			Err:    "query parameter 'distance' is required",
		})
		return
	}

	points := h.pointsService.GetPoints(x, y, distance)

	c.JSON(http.StatusOK, &points)
}
