package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/tidwall/gjson"
	"gopkg.in/jarcoal/httpmock.v1"
)

func Test(t *testing.T) {

	var testData []map[string]string
	var a, b string

	// test stuff
	testData = setupTest("test1.json", "05/01/2017")

	a = testData[0]["homeTeamName"]
	b = "Yankees     "
	assertTest(a, b, t)

	a = testData[1]["awayBox"]
	b = "0 0 0 0 1 1 0 3 0"
	assertTest(a, b, t)

	a = testData[2]["homeTeamTotalRuns"]
	b = " 7"
	assertTest(a, b, t)

	// test more of the same with different data
	testData = setupTest("test2.json", "06/01/2017")

	a = testData[0]["awayBox"]
	b = "0 0 0 0 0 0 0 0 0"
	assertTest(a, b, t)

	a = testData[0]["awayTeamTotalRuns"]
	b = " 0"
	assertTest(a, b, t)

	a = testData[2]["homeTeamName"]
	b = "Cardinals   "
	assertTest(a, b, t)

	// test single game (World Series)
	testData = setupTest("test3.json", "11/01/2017")

	if len(testData) != 1 {
		t.Errorf("FAIL: Expected 1 game, received %d .", len(testData))
	}

	a = testData[0]["awayTeamTotalRuns"]
	b = "99"
	assertTest(a, b, t)

}

func setupTest(fileName string, dateArg string) []map[string]string {

	date := getDate(dateArg)
	url := getURL(date)

	bytes, err := ioutil.ReadFile("testdata/" + fileName)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, string(bytes)))

	gameList := getDataFromMLB(dateArg)
	gameData := parseData(gameList)

	return gameData
}

func readTestFile(f string) []gjson.Result {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	return gjson.GetBytes(bytes, "data.games.game").Array()
}

func assertTest(a string, b string, t *testing.T) {
	if a != b {
		t.Errorf("FAIL: Expected '%s', received '%s'.", a, b)
	}
}
