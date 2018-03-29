package main

import (
	"fmt"

	"github.com/apcera/termtables"
)

func displayData(gameData []map[string]string) {

	table := termtables.CreateTable()

	for count, item := range gameData {
		table.AddRow(item["awayTeamName"], item["awayBox"], item["awayTeamTotalRuns"])
		table.AddRow(item["homeTeamName"], item["homeBox"], item["homeTeamTotalRuns"])
		if count < len(gameData)-1 {
			table.AddSeparator()
		}
	}

	fmt.Println(table.Render())

}

func displayHelp() {
	fmt.Println("mlbcli               defaults to 'today'")
	fmt.Println("mlbcli today         today's games")
	fmt.Println("mlbcli tomorrow      tomorrow's games")
	fmt.Println("mlbcli yesterday     yesterday's games")
	fmt.Println("mlbcli Friday        upcoming Friday's games")
	fmt.Println("mlbcli 7/11          defaults to current year")
	fmt.Println("mlbcli 10/27/2011    specific date")
}
