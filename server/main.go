package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HistoryItem struct {
	ID        string `json:"id"`
	Operation string `json:"operation"`
	Result    string `json:"result"`
}

var HistoryItems = []HistoryItem{
	{ID: "1", Operation: "2+2+2+2", Result: "8"},
	{ID: "2", Operation: "2+2+2+2", Result: "8"},
	{ID: "3", Operation: "2+2+2+2", Result: "8"},
	{ID: "4", Operation: "2+2+2+2", Result: "8"},
	{ID: "5", Operation: "2+2+2+2", Result: "8"},
}

func getOperations (context *gin.Context){
context.IndentedJSON(http.StatusOK, HistoryItems)
}
func main() {
	router := gin.Default()
	router.GET("/history", getOperations)
	router.Run("localhost:9090")
}
