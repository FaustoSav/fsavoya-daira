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
	router.GET("/history", getOperations)
	router.POST("/history", addOperation)
	router.Run("localhost:9090")
}
