package main

import (
	"os"

	"github.com/tidwall/match"
)

func main() {

	for _, i := range os.Args {
		if match.Match(i, "*help") || match.Match(i, "*-h") {
			displayHelp()
			os.Exit(0)
		}
	}

	var dateArg string

	if len(os.Args) < 2 {
		dateArg = "today"
	} else {
		dateArg = os.Args[1]
	}

	gameList := getDataFromMLB(dateArg)
	gameData := parseData(gameList)
	displayData(gameData)
	os.Exit(0)

}
