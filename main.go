package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/aymone/cartesian-api/handlers"
	"github.com/aymone/cartesian-api/models"
	"github.com/aymone/cartesian-api/services"
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

	s := services.NewPointsService(points)
	h := handlers.New(s)
	r := gin.New()

	// Define handlers
	r.GET("/api/points", h.GetPoints)

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
