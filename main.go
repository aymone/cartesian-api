package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aymone/cartesian-api/models"
	"github.com/aymone/cartesian-api/services"
	"github.com/aymone/cartesian-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	jsonFile, err := os.Open("data/points.json")
	if err != nil {
		log.Printf("Error opening data points file: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var points []*models.Points

	json.Unmarshal(byteValue, &points)

	service := services.NewPointsService(points)

	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define handlers
	r.GET("/api/points", func(c *gin.Context) {
		q := c.Request.URL.Query()

		ok, x := utils.ParseInt(q.Get("x"))
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
				Status: http.StatusBadRequest,
				Err:    "query parameter 'x' is required",
			})
			return
		}

		ok, y := utils.ParseInt(q.Get("y"))
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
				Status: http.StatusBadRequest,
				Err:    "query parameter 'y' is required",
			})
			return
		}

		ok, distance := utils.ParseInt(q.Get("distance"))
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
				Status: http.StatusBadRequest,
				Err:    "query parameter 'distance' is required",
			})
			return
		}

		points := service.GetPoints(&models.Points{
			X:        x,
			Y:        y,
			Distance: distance,
		})

		c.JSON(http.StatusOK, &points)
	})

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
