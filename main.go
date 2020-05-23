package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/aymone/cartesian-api/handlers"
	"github.com/aymone/cartesian-api/models"
	"github.com/aymone/cartesian-api/router"
	"github.com/aymone/cartesian-api/services"
)

const portEnv string = "PORT"
const portDefault string = "8080"
const jsonFilePath string = "data/points.json"

func main() {
	port := os.Getenv(portEnv)
	if port == "" {
		port = portDefault
		log.Printf("Defaulting to port %s", port)
	}

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		log.Printf("Error opening data points file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var points []*models.Points
	json.Unmarshal(byteValue, &points)

	s := services.NewPointsService(points)
	h := handlers.New(s)
	r := router.Init(h)

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
