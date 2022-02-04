package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

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
	router.GET("/hangman/api/nextwrdtime", hangman_api_nextword)
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

func hangman_api_nextword(c *gin.Context) {

	time_available := base_date.AddDate(0, 0, hangman_dayword(time.Now())+1)
	time_available_json := map[string]int{
		"year":   time_available.Year(),
		"month":  int(time_available.Month()),
		"day":    time_available.Day(),
		"hour":   time_available.Hour(),
		"minute": time_available.Minute(),
		"second": time_available.Second(),
	}
	c.IndentedJSON(http.StatusOK, time_available_json)
}

func hangman_api_lengthcheck(c *gin.Context) {
	id := c.Param("id")

	for _, a := range wordlists {
		if a.Id == id {
			// fmt.Println(a)
			var word = a.Words[hangman_dayword(time.Now())]
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

			var word = strings.ToUpper(a.Words[hangman_dayword(time.Now())])

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

func hangman_dayword(date time.Time) int {
	days_between_dates := date.Sub(base_date).Hours() / 24
	// days_between_dates := 3 // Used for debugging
	return int(math.Floor(days_between_dates))
}

type wordlist struct {
	Id    string         `json:"id"`
	Words map[int]string `json:"words"`
}

// var today_date = time.Date(2022, time.Month(2), 5, 5, 0, 0, 0, time.UTC)
var base_date = time.Date(2022, time.Month(2), 4, 0, 0, 0, 0, time.UTC)

var words_abc123 = map[int]string{
	0: "doggo", 1: "foxy", 2: "wiggly", 3: "blasty", 4: "tufferina", 5: "england", 6: "apartment", 7: "brooklyn", 8: "london",
}

var wordlists = []wordlist{
	{Id: "abc123", Words: words_abc123},
}
