package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func getDataFromMLB(dateArg string) []gjson.Result {

	date := getDate(dateArg)
	url := getURL(date)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error connecting to MLB API.")
		os.Exit(1)
	}
	defer response.Body.Close()

	json, err := ioutil.ReadAll(response.Body)
	if err != nil || !gjson.Valid(string(json)) {
		fmt.Println("Error with data returned from MLB API.")
		os.Exit(1)
	}

	if !gjson.GetBytes(json, "data.games.game").Exists() {
		m := "Was there a game on this date?"
		switch dateArg {
		case "today":
			m = "Is there a game today?"
		case "tomorrow":
			m = "Is there a game tomorrow?"
		case "yesterday":
			m = "Is there a game tomorrow?"
		}

		// TODO

		fmt.Println("Problem with data returned from MLB API. " + m)
		os.Exit(1)
	}

	if !gjson.GetBytes(json, "data.games.game").IsArray() {
		return []gjson.Result{gjson.GetBytes(json, "data.games.game")}
	}
	return gjson.GetBytes(json, "data.games.game").Array()
}

func getDate(dateArg string) map[string]string {

	format := "01/02/2006" // golang date format reference: https://golang.org/src/time/format.go
	currentDate := time.Now().Local().Format(format)

	switch dateArg {
	case "today", "tod":
		dateArg = currentDate
	case "tomorrow", "tom":
		dateArg = time.Now().Local().AddDate(0, 0, 1).Format(format)
	case "yesterday", "yes":
		dateArg = time.Now().Local().AddDate(0, 0, -1).Format(format)
	}

	if len(dateArg) < 3 {
		dateArg = "xxx"
	}

	dayOfWeek := strings.ToLower(dateArg[:3])
	match, _ := regexp.MatchString("sun|mon|tue|wed|thu|fri|sat", dayOfWeek)
	if match {
		dateArg = getNextDate(dayOfWeek, format)
	}

	dateArg = strings.Replace(dateArg, "-", "/", -1)

	if !strings.Contains(dateArg, "/") {
		displayHelp()
		os.Exit(1)
	}

	date := strings.Split(dateArg, "/")

	if len(date) < 3 {
		y := strings.Split(currentDate, "/")[2]
		date = append(date, y)
	}

	m := "0" + date[0]
	d := "0" + date[1]
	y := "20" + date[2]
	month := m[len(m)-2:]
	day := d[len(d)-2:]
	year := y[len(y)-4:]

	timestamp := month + "/" + day + "/" + year
	_, error := time.Parse(format, timestamp)

	yy, _ := strconv.Atoi(year)
	cy, _ := strconv.Atoi(strings.Split(currentDate, "/")[2])

	if error != nil || yy < 2007 || yy > cy {
		displayHelp()
		fmt.Println("ERROR: Invalid date")
		os.Exit(1)
	}
	return map[string]string{"month": month, "day": day, "year": year}
}

func getNextDate(day string, format string) string {
	weekdays := map[string]int{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}
	currentDay := int(time.Now().Weekday())
	nextDay := weekdays[day] - currentDay
	if nextDay < 1 {
		nextDay += 7
	}
	nextDate := time.Now().Local().AddDate(0, 0, nextDay).Format(format)
	return nextDate
}

func getURL(date map[string]string) string {
	return "http://gd2.mlb.com/components/game/mlb/year_" + date["year"] + "/month_" + date["month"] + "/day_" + date["day"] + "/master_scoreboard.json"
}
