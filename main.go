package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Healthy")
}

func main() {
	router := gin.Default()
	router.GET("health", getHealth)

	router.Run("localhost:8080")
}
