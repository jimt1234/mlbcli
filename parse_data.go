package main

import (
	"github.com/tidwall/gjson"
)

var s, ss, sss = " ", "  ", "                 " // spaces for formatting the table

func parseData(games []gjson.Result) []map[string]string {

	var gameData = []map[string]string{}

	for _, v := range games {

		var awayBox, homeBox string
		var awayTeamTotalRuns, homeTeamTotalRuns string

		gameStatus := gjson.Get(v.Raw, "status.status").String()

		switch gameStatus {
		case "In Progress", "Completed Early", "Game Over", "Final":
			awayBox, homeBox = boxScores(v)
			awayTeamTotalRuns, homeTeamTotalRuns = totalRuns(v)
		case "Preview", "Warmup", "Pre-Game":
			awayBox = gameStatus
			homeBox = (gjson.Get(v.Raw, "time_hm_lg").String() + s + gjson.Get(v.Raw, "hm_lg_ampm").String() + s + gjson.Get(v.Raw, "tz_hm_lg_gen").String() + sss)[:17]
			awayTeamTotalRuns, homeTeamTotalRuns = ss, ss
		case "Postponed", "Cancelled":
			awayBox = gameStatus
			homeBox = sss
			awayTeamTotalRuns, homeTeamTotalRuns = ss, ss
		default:
			awayBox = "Error: '" + gameStatus + "' ???"
			homeBox = sss
			awayTeamTotalRuns, homeTeamTotalRuns = ss, ss
		}

		awayTeamName := (gjson.Get(v.Raw, "away_team_name").String() + sss)[:12]
		homeTeamName := (gjson.Get(v.Raw, "home_team_name").String() + sss)[:12]

		gameData = append(gameData, map[string]string{
			"awayTeamName":      awayTeamName,
			"homeTeamName":      homeTeamName,
			"awayBox":           awayBox,
			"homeBox":           homeBox,
			"awayTeamTotalRuns": awayTeamTotalRuns,
			"homeTeamTotalRuns": homeTeamTotalRuns,
		})
	}
	return gameData
}

func boxScores(v gjson.Result) (string, string) {

	var awayBox, homeBox string

	inningRuns := gjson.Get(v.Raw, "linescore.inning").Array()

	for count, x := range inningRuns {
		awayInning := gjson.Get(x.Raw, "away").String()
		homeInning := gjson.Get(x.Raw, "home").String()
		if len(awayInning) < 1 {
			awayInning = s
		}
		if len(homeInning) < 1 {
			homeInning = s
		}

		awayBox += awayInning
		homeBox += homeInning

		if count < len(inningRuns)-1 {
			awayBox += s
			homeBox += s
		}
	}
	return awayBox, homeBox
}

func totalRuns(v gjson.Result) (string, string) {

	var awayTeamTotalRuns, homeTeamTotalRuns string

	awayTeamRuns := gjson.Get(v.Raw, "linescore.r.away").String()
	if len(awayTeamRuns) < 1 {
		awayTeamTotalRuns = ss
	} else {
		atr := ss + awayTeamRuns
		awayTeamTotalRuns = string(atr[len(atr)-2:])
	}
	homeTeamRuns := gjson.Get(v.Raw, "linescore.r.home").String()
	if len(homeTeamRuns) < 1 {
		homeTeamTotalRuns = ss
	} else {
		htr := ss + homeTeamRuns
		homeTeamTotalRuns = string(htr[len(htr)-2:])
	}

	return awayTeamTotalRuns, homeTeamTotalRuns

}
