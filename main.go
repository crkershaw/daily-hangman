package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	config "github.com/crkershaw/hangman/configs"
	addwords "github.com/crkershaw/hangman/controllers/addwords"
	db "github.com/crkershaw/hangman/controllers/db"
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

	if fileExists("public/env-vars.json") {
		set_os_vars("public/env-vars.json")
	}

	if config.ConfigSource == "db" {
		db.Database()
	}

	router := gin.Default()

	router.Static("/public/css", "./public/css")
	router.Static("/public/images", "./public/images")
	router.Static("/public/js", "./public/js")

	router.LoadHTMLGlob("templates/*")

	hangman.Routes(router)  // Loading the routes from the hangman module
	addwords.Routes(router) // Loading the routes from the addwords module

	router.Run(":80")

	log.Fatal(router.Run())
}

// Adding S3 bucket url and database access details to environment variables
// Note these are not added to Github - if using with your own data, change ConfigSource variable in configs/config.go

type Os_Vars struct {
	S3_FILE_URL string `json:"S3_FILE_URL"`
	DBUSER      string `json:"DBUSER"`
	DBPASS      string `json:"DBPASS"`
	DBADDRESS   string `json:"DBADDRESS"`
}

func set_os_vars(config_file_path string) {

	jsonFile, err := os.Open(config_file_path)

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var osvars Os_Vars

	// Unmarshalling byteArray containing the jsonFile's content into 'osvars'
	json.Unmarshal(byteValue, &osvars)

	os.Setenv("S3_FILE_URL", osvars.S3_FILE_URL)
	os.Setenv("DBUSER", osvars.DBUSER)
	os.Setenv("DBPASS", osvars.DBPASS)
	os.Setenv("DBADDRESS", osvars.DBADDRESS)

}
