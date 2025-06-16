package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TemperatureData struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func main() {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Get temperature by location
	router.GET("/temperature", func(c *gin.Context) {
		location := c.Query("location")
		if location == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
			return
		}

		// Generate random temperature data based on location
		data := generateTemperatureData(location, "")

		c.JSON(http.StatusOK, data)
	})

	// Get temperature by sensor ID
	router.GET("/temperature/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Sensor ID is required"})
			return
		}

		// Generate random temperature data based on sensor ID
		data := generateTemperatureData("", id)

		c.JSON(http.StatusOK, data)
	})

	// Start server
	log.Println("Temperature API starting on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func generateTemperatureData(location, sensorID string) TemperatureData {
	// Generate a random temperature between 18 and 28 degrees Celsius
	value := 18.0 + float64(time.Now().UnixNano()%10) + float64(time.Now().UnixNano()%100)/100.0

	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	return TemperatureData{
		Value:       value,
		Unit:        "°C",
		Timestamp:   time.Now(),
		Location:    location,
		Status:      "active",
		SensorID:    sensorID,
		SensorType:  "temperature",
		Description: "Temperature sensor in " + location,
	}
}
