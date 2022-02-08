package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	// Note S3 bucket url is not added to Github - if using with your own data, change config_source config variable beloq
	if config_source == "s3" && fileExists("env-vars.txt") {
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

	router.GET("/", hangman)
	router.GET("/c/:id", hangman)
	router.GET("/hangman/api/nextwrdtime", hangman_api_nextword)
	router.GET("/hangman/api/wrdlen/:id", hangman_api_lengthcheck)
	router.GET("/hangman/api/ltrchk/:id", hangman_api_lettercheck)
	router.GET("/hangman/api/msg/:id", hangman_api_message)

	router.Run(":80")

	log.Fatal(router.Run())
}

func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", []string{})
}

func hangman(c *gin.Context) {
	c.HTML(http.StatusOK, "hangman.html", []string{})
}

// Returns all words for chosen ID from csv file
func hangman_getwords(wordlist_id string) map[string]map[string]string {

	var records [][]string

	if config_source == "s3" {
		records = readCsvFile("s3", os.Getenv("s3-file-url"))
	} else {
		records = readCsvFile("local", "wordlist/wordlist_custom.csv")
	}

	header := []string{} // holds first row - header
	// body := []map[string]string{}
	words := map[string]map[string]string{}

	// Looping through csv line by line
	for lineNum, record := range records {

		// Setting headers
		if lineNum == 0 {
			for i := 0; i < len(record); i++ {
				header = append(header, strings.TrimSpace(
					// Excel appends this hidden start of document character when making a csv, so we are removing it
					strings.Replace(record[i], "\ufeff", "", -1)))
			}

			// Transforming lines into map
		} else {
			line := map[string]string{}
			for i := 0; i < len(record); i++ {
				line[header[i]] = record[i]
			}

			// Add to output if id matches the selected ID
			if line["id"] == wordlist_id {

				words[line["num"]] = map[string]string{
					"word":    line["word"],
					"message": line["message"],
				}

				// body = append(body, line)
			}
		}
	}

	return words
}

// Returns chosen word and message for given wordlist and day number
func hangman_getwordmessage(wordlist_id string, day int) (string, string) {

	var wordlist map[string]map[string]string

	if config_source == "hardcoded" {
		wordlist = wordlists["default"]
	} else if config_source == "csv" || config_source == "s3" {
		wordlist = hangman_getwords(wordlist_id)
	}

	day_str := strconv.Itoa(hangman_dayword(time.Now()))

	word := wordlist[day_str]["word"]
	message := wordlist[day_str]["message"]

	return word, message
}

// Returning time that next word will be available
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

// Returning length of the chosen word
func hangman_api_lengthcheck(c *gin.Context) {

	word, _ := hangman_getwordmessage(c.Param("id"), hangman_dayword(time.Now()))
	c.IndentedJSON(http.StatusOK, len(word))

}

// Given a letter, returns map of where in the word it is
func hangman_api_lettercheck(c *gin.Context) {
	id := c.Param("id")
	letter := strings.ToUpper(c.Query("letter"))

	fmt.Println("Received letter check api call, ID: " + string(id) + ", Letter: " + string(letter))

	var answer_list = map[int]string{}

	word, _ := hangman_getwordmessage(id, hangman_dayword(time.Now()))
	word_upper := strings.ToUpper(word)

	for index, value := range word_upper {

		if string(value) == string(letter) {
			answer_list[index] = string(value)
		}
	}

	c.IndentedJSON(http.StatusOK, answer_list)

}

// Returns the message associated with the word
// Note: stateless so may cause minor issues if they start before the word changes, and finish after
// May want to require the word in the request
func hangman_api_message(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)

	_, message := hangman_getwordmessage(id, hangman_dayword(time.Now()))
	c.IndentedJSON(http.StatusOK, message)
}

// Returns the 'day' that we are on, vs the base date, to be used to pick which word to give
func hangman_dayword(date time.Time) int {
	days_between_dates := date.Sub(base_date).Hours() / 24
	// days_between_dates := 3 // Used for debugging
	return int(math.Floor(days_between_dates))
}

func readCsvFile(source string, filePath string) [][]string {

	var csvReader *csv.Reader

	if source == "local" {
		f, err := os.Open(filePath)

		if err != nil {
			log.Fatal("Unable to read input file "+filePath, err)
		}

		defer f.Close()
		csvReader = csv.NewReader(f)
	} else {
		f, err := http.Get(filePath)

		if err != nil {
			log.Fatal("Unable to read input file "+filePath, err)
		}

		defer f.Body.Close()
		csvReader = csv.NewReader(f.Body)
	}

	// var records [][]string
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

// var today_date = time.Date(2022, time.Month(2), 5, 5, 0, 0, 0, time.UTC)
var base_date = time.Date(2022, time.Month(2), 4, 0, 0, 0, 0, time.UTC)

var words_default = map[string]map[string]string{
	"0": {
		"word":    "doggo",
		"message": "Doggos make life worth living",
	},
	"1": {
		"word":    "avocado",
		"message": "House prices are proportional to avocado consumption",
	},
	"2": {
		"word":    "brooklyn",
		"message": "Cooking raw with the Brooklyn boy",
	},
	"3": {
		"word":    "guatemala",
		"message": "We can get away-ay, maybe to Guatemala…",
	},
	"4": {
		"word":    "football",
		"message": "Sissoko…kicks it up towards Llorente…Dele Alli…they're slipping…they're sliding…IT'S IN LUCAS MOURA WITH THE HATTRICK GOAL",
	},
	"5": {
		"word":    "periodic",
		"message": "My friend, you would not tell with such high zest, To children ardent for some desperate glory, The old lie: Dulce et decorum est, Pro patria mori",
	},
	"6": {
		"word":    "hedonistic",
		"message": "The only way to get rid of temptation is to yield to it. Resist it, and your soul grows sick with longing for the things it has forbidden to itself",
	},
	"7": {
		"word":    "england",
		"message": "It's coming home…it's coming home…it's coming…FOOTBALL'S COMING HOME",
	},
	"8": {
		"word":    "climbing",
		"message": "Do not go gentle into that good night, Forearms should burn and rave at close of day; Climb, climb, to ever greater heights",
	},
}

var wordlists = map[string]map[string]map[string]string{
	"default": words_default,
}

var config_sources = [3]string{"csv", "hardcoded", "s3"}
var config_source string = config_sources[2]
