# Hangman
[http://hangman.charliekershaw.com/](http://hangman.charliekershaw.com/)

A simple game of hangman, which gives the player one word per day and rewards them with a message when they get it right.

Front-end is build in Javascript React/HTML/css, and the back-end is written in Go using the Gin web framework. Packaged up in Docker, and hosted on AWS.

### Fundamentals
1. Player is served one word per day, with a message appearing when they get it right
2. All data for what that word is is on the back-end
3. The back-end has API endpoints to give:
    * The length of the word
    * The positions (if any) of a requested letter in the word
    * The victory message for the word
    * The time until the next word is available
4. The raw data has four key variables:
    * Word list: the id of the wordlist itself ('default' is the id used for the base page localhost/, others ids are loaded at localhost/c/:id ) 
    * Word number: the day on which this word will appear (vs. the BaseDate defined in config/config.go). For example, word 2 will appear 2 days after the BaseDate
    * Word: the word itself, to be guessed in the game
    * Message: the message associated with the word, to be displayed once the word is guessed
5. The homepage laods the 'default' word list, but hangman with your custom wordlist can be viewed at localhost/c/:custom_wordlist_id

### Core Structure
| Location | Purpose |
|----------|---------|
|configs/config.go      | Module for configuration setup - including settings that must be adjusted to run locally, such as where the wordlist is sourced from |
|controllers/hangman    | Module containing the code for the hangman back-end API |
|public/                | Contains the web assets - css and javascript |
|wordlist/              | Contains a csv file that the words can be read from (configurable in config.go) |
|main.go                | The main server-running file

## Running it yourself
### Fork the repository
First, fork the repository so you have your own copy.

### Update the configs/config.go file
You must change:
1. The ConfigSource defaults to S3, where a csv of words and messages is stored in an AWS S3 bucket. However the location of this is saved in an environment variable not pushed to Github, so you need to point it to either a local csv or the hardcoded list of words and messages in the configs/config.go file.
2. You will likely want to change the BaseDate to be a more recent date. The application serves a word per day, and decides which word to serve by counting the days since the BaseDate, and serving the word that corresponds to that number in the wordlist. So if the BaseDate is too far in the past, you may not have a word tagged to the current date.

### Running locally
You can run locally by running the main.go file

### Running on docker
You can run through docker by cd-ing to the main folder, then running:
```
docker build -t hangman .
docker run --rm -p 80:80 hangman
```

# Accessing it
You can access the site (if running locally) by going to http://localhost

# Notes for adding new modules
1. Create a folder for the module (e.g. controllers/hangman)
2. Ensure the .go file in that declares the package name at the top
```
package hangman
```
3. cd to the folder, and run go mod init:
```
$ go mod init github.com/crkershaw/hangman/controllers/hangman
```
4. Import that module by referencing that url
```
--- main.go ---
import (
    github.com/crkershaw/hangman/controllers/hangman
)
```
5. Ensure that the module calling that module knows where to find it locally by adding this line to its go.mod
```
--- hangman/go.mod
replace github.com/crkershaw/hangman/controllers/hangman => /controllers/hangman
```

6. cd to the directory of the file calling the module, and run go mod tidy
```
cd C:\SoftDev\hangman\
go mod tidy
```
