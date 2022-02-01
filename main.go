package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", homepage)

	router.Run("localhost:8080")

	log.Fatal(router.Run())
}

func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", []string{})

}
