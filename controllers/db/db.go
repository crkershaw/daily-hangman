package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Returned from database
type WordRow struct {
	Id_num        int
	Wordlist_name string
	Word_num      int
	Word          string
	Message       string
	Creation_date time.Time
}

// Tranformed to these structures
type WordDetail struct {
	Word    string
	Message string
}

type Wordlist struct {
	Wordlist map[string]WordDetail
}

type Wordlists struct {
	Wordlist_name string
	Wordlist      Wordlist
	Creation_date time.Time
}

func Database() {

	// Capture connection properties
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBADDRESS"),
		DBName:               "hangman",
		AllowNativePasswords: true,
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to database")
}

// map[string]map[string]string

func Get_wordlist_fromdb(wordlist_name string) (*Wordlist, error) {

	var words_raw []WordRow

	rows, err := db.Query("SELECT word_num, word, message FROM wordlist WHERE wordlist_name  = ?", wordlist_name)

	if err != nil {
		log.Fatalf("get_wordlist_fromdb %q: %v", wordlist_name, err)
		return nil, errors.New("Error in get_wordlist_fromdb")
	}

	defer rows.Close()

	for rows.Next() {
		var word WordRow
		if err := rows.Scan(&word.Word_num, &word.Word, &word.Message); err != nil {
			log.Fatalf("get_wordlist_fromdb %q: %v", wordlist_name, err)
			return nil, errors.New("Error in get_wordlist_fromdb")
		}

		words_raw = append(words_raw, word)
	}

	wordlist := Wordlist{Wordlist: map[string]WordDetail{}}

	for _, item := range words_raw {
		var word_detail WordDetail

		word_detail.Word = item.Word
		word_detail.Message = item.Message

		wordlist.Wordlist[strconv.Itoa(item.Word_num)] = word_detail

		// = Wordlist{Wordlist: map[string]WordDetail{
		// 	strconv.Itoa(item.Word_num): word_detail}}
	}

	fmt.Println("Word list returned from database:")
	fmt.Println(wordlist)

	return &wordlist, nil

}
