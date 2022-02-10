package hangman

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	config "github.com/crkershaw/hangman/configs"
	db "github.com/crkershaw/hangman/controllers/db"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {

	hangman := route.Group("/")
	hangman.GET("/", hangman_page)
	hangman.GET("/c/:id", hangman_page)
	hangman.GET("/hangman/api/nextwrdtime", hangman_api_nextword)
	hangman.GET("/hangman/api/wrdlen/:id", hangman_api_lengthcheck)
	hangman.GET("/hangman/api/ltrchk/:id", hangman_api_lettercheck)
	hangman.GET("/hangman/api/msg/:id", hangman_api_message)
}

func hangman_page(c *gin.Context) {
	c.HTML(http.StatusOK, "hangman.html", []string{})
}

func hangman_getwords_db(wordlist_id string) map[string]map[string]string {

	wordlist, _ := db.Get_wordlist_fromdb(wordlist_id) // Update to catch error
	var wordlist_deref = *wordlist

	wordlist_newformat := map[string]map[string]string{}

	for key, item := range wordlist_deref.Wordlist {

		wordlist_newformat[key] = map[string]string{
			"word":    item.Word,
			"message": item.Message,
		}
	}

	return wordlist_newformat

}

// Returns dictionary of all words for chosen ID from csv file
func hangman_getwords(wordlist_id string) map[string]map[string]string {

	var records [][]string

	if config.ConfigSource == "s3" {
		records = readCsvFile("s3", os.Getenv("S3_FILE_URL"))
	} else if config.ConfigSource == "csv" {
		records = readCsvFile("local", "wordlist/wordlist_custom.csv")
	}

	header := []string{} // holds first row - header
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
			}
		}
	}

	return words
}

// Returns chosen word and message for given wordlist and day number
func hangman_getwordmessage(wordlist_id string, day int) (string, string) {

	var wordlist map[string]map[string]string

	if config.ConfigSource == "hardcoded" {
		wordlist = config.WordLists["default"]
	} else if config.ConfigSource == "csv" || config.ConfigSource == "s3" {
		wordlist = hangman_getwords(wordlist_id)
	} else if config.ConfigSource == "db" {
		wordlist = hangman_getwords_db(wordlist_id)
	}

	day_str := strconv.Itoa(hangman_dayword(time.Now()))

	word := wordlist[day_str]["word"]
	message := wordlist[day_str]["message"]

	return word, message
}

// Returning time that next word will be available
func hangman_api_nextword(c *gin.Context) {

	time_available := config.BaseDate.AddDate(0, 0, hangman_dayword(time.Now())+1)
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
	days_between_dates := date.Sub(config.BaseDate).Hours() / 24
	// days_between_dates := 3 // Used for debugging - change to check a certain word
	return int(math.Floor(days_between_dates))
}

// Loads data from a local or cloud-hosted csv
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

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
