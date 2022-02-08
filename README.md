# Hangman

## Front-end
Front-end is html, css and javascript using the React framework

## Back-end
Back-end is written in Go

## Structure
configs/config.go       Configuration setup - including settings that must be adjusted to run locally, such as where the wordlist is sourced from

# Running on docker

```
docker build -t hangman .

docker run --rm -p 80:80 hangman
```



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
