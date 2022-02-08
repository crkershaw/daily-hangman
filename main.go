package main

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	config "github.com/crkershaw/hangman/configs"
	"github.com/crkershaw/hangman/controllers/hangman"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {

	// Adding S3 bucket to environment variables
	// Note S3 bucket url is not added to Github - if using with your own data, change ConfigSource variable in configs/config.go
	if config.ConfigSource == "s3" && fileExists("env-vars.txt") {
		file, err := os.Open("env-vars.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			os.Setenv("s3-file-url", scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	router := gin.Default()

	router.Static("/public/css", "./public/css")
	router.Static("/public/images", "./public/images")
	router.Static("/public/js", "./public/js")

	router.LoadHTMLGlob("templates/*")

	hangman.Routes(router) // Loading the routes from the hangman module

	router.Run(":80")

	log.Fatal(router.Run())
}

func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", []string{})
}
