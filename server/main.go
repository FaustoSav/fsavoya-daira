package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HistoryItem struct {
	ID        string `json:"id"`
	Operation string `json:"operation"`
	Result    string `json:"result"`
}

var HistoryItems = []HistoryItem{}

var lastID int

func getNextID() string {
	lastID++
	return strconv.Itoa(lastID)
}

func getOperations(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, HistoryItems)
}

func addOperation(context *gin.Context) {
	var newOperation HistoryItem
	if err := context.BindJSON(&newOperation); err != nil {
		return
	}

	newOperation.ID = getNextID()
	HistoryItems = append(HistoryItems, newOperation)

	context.IndentedJSON(http.StatusCreated, newOperation)
}

func main() {
	lastID = len(HistoryItems) // Establecer lastID como el último ID actual
	router := gin.Default()

	// Middleware CORS para permitir todas las solicitudes desde cualquier origen
	router.Use(corsMiddleware())

	router.GET("/history", getOperations)
	router.POST("/history", addOperation)
	router.Run("localhost:9090")
}

// Middleware CORS para permitir todas las solicitudes desde cualquier origen
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Permitir solicitudes solo desde http://localhost:5173
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")

		// Configurar las credenciales para el modo 'include'
		c.Header("Access-Control-Allow-Credentials", "true")

		// Manejar solicitudes OPTIONS
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
