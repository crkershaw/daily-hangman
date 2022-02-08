package config

import (
	"time"
)

// Where the data for the hangman words and messages should be pulled from

// hardcoded: from the WordsDefault dictionary below - recommended for first time running locally
// csv: from the csv file wordlist/wordlist_default
// s3: from the csv file in the s3 bucket defined in the environment variable set in env-vars.txt (not committed to Github)
var ConfigSources = [3]string{"csv", "hardcoded", "s3"}
var ConfigSource string = ConfigSources[2]

// Dates: BaseDate is the date the word of day number '0' appears on
// So if it is two days after the BaseDate, the user will be shown word '2'
var BaseDate = time.Date(2022, time.Month(2), 4, 0, 0, 0, 0, time.UTC)

// var TodayDate = time.Date(2022, time.Month(2), 5, 5, 0, 0, 0, time.UTC)

// Hardcoded wordlists - ideal for first time use getting it working
var WordsDefault = map[string]map[string]string{
	"0": { // Day number: number of days from BaseDate that this word shows on
		"word":    "doggo",                         // The hangman word itself
		"message": "Doggos make life worth living", // The message to display once one gets the word correct
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

var WordLists = map[string]map[string]map[string]string{
	"default": WordsDefault,
}
