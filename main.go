package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/public/css", "./public/css")
	router.Static("/public/images", "./public/images")
	router.Static("/public/js", "./public/js")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", homepage)
	router.GET("/hangman", hangman)
	router.GET("/hangman/api/wrdlen/:id", hangman_api_lengthcheck)
	router.GET("/hangman/api/ltrchk/:id", hangman_api_lettercheck)

	router.Run(":80")

	log.Fatal(router.Run())
}

func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", []string{})
}

func hangman(c *gin.Context) {
	c.HTML(http.StatusOK, "hangman.html", []string{})
}

func hangman_api_lengthcheck(c *gin.Context) {
	id := c.Param("id")

	for _, a := range wordlists {
		if a.Id == id {
			// fmt.Println(a)
			var word = a.Words[word_index]
			c.IndentedJSON(http.StatusOK, len(word))
			return
		}
	}
}

func hangman_api_lettercheck(c *gin.Context) {
	id := c.Param("id")
	letter := strings.ToUpper(c.Query("letter"))

	fmt.Println("Received letter check api call, ID: " + string(id) + ", Letter: " + string(letter))

	for _, a := range wordlists {

		var answer_list = map[int]string{}

		if a.Id == id {
			var word = strings.ToUpper(a.Words[word_index])

			for index, value := range word {

				if string(value) == string(letter) {
					answer_list[index] = string(value)
				}
			}

		}

		c.IndentedJSON(http.StatusOK, answer_list)
		return
	}
}

type wordlist struct {
	Id    string   `json:"id"`
	Words []string `json:"words"`
}

var word_index = 2
var words = []string{"doggo", "foxy", "wiggly", "blasty", "tufferina", "england", "apartment"}

var wordlists = []wordlist{
	{Id: "abc123", Words: words},
}
