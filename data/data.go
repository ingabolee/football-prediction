package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	automate()
}

func automate() {
	var teamIds [10][2]interface{}
	var standings [20][2]interface{}
	var goals []map[string]interface{}
	var results interface{}
	var matchweek int

	for {
		teamIdsChan := make(chan [10][2]interface{})
		matchweekChan := make(chan int)
		standingsChan := make(chan [20][2]interface{})

		go dataCollectorOne(teamIdsChan, matchweekChan)
		go dataCollectorTwo(standingsChan)
		_, goals = dataCollectorThree()
		teamIds = <-teamIdsChan
		standings = <-standingsChan
		matchweek = <-matchweekChan

		log.Println("Collecting data for matchweek ", matchweek)
		log.Println("Waiting for matchweek results...")
		var sleepTime int = 120
		time.Sleep(time.Duration(sleepTime) * time.Second)

		results, _ = dataCollectorThree()

		table := [...]map[string]interface{}{
			{"team_id": 1, "team_name": "Manchester Blue"},
			{"team_id": 2, "team_name": "Manchester Reds"},
			{"team_id": 3, "team_name": "Liverpool"},
			{"team_id": 4, "team_name": "London Blues"},
			{"team_id": 5, "team_name": "Tottenham"},
			{"team_id": 6, "team_name": "London Reds"},
			{"team_id": 7, "team_name": "Burnley"},
			{"team_id": 8, "team_name": "Leicester"},
			{"team_id": 9, "team_name": "Everton"},
			{"team_id": 10, "team_name": "LEEDS"},
			{"team_id": 11, "team_name": "WEST BROM"},
			{"team_id": 12, "team_name": "West Ham"},
			{"team_id": 13, "team_name": "Newcastle"},
			{"team_id": 14, "team_name": "Brighton"},
			{"team_id": 15, "team_name": "Palace"},
			{"team_id": 16, "team_name": "FULHAM"},
			{"team_id": 17, "team_name": "ASTON V"},
			{"team_id": 18, "team_name": "Southampton"},
			{"team_id": 19, "team_name": "Wolves"},
			{"team_id": 20, "team_name": "SHEFFIELD U"},
		}

		for _, team_id := range teamIds {
			teamIdAway := team_id[0]
			teamIdHome := team_id[1]
			var home_team, away_team string

			for _, team := range table {
				if fmt.Sprint(team["team_id"]) == fmt.Sprint(teamIdHome) {
					home_team = fmt.Sprint(team["team_name"])
				} else if fmt.Sprint(team["team_id"]) == fmt.Sprint(teamIdAway) {
					away_team = fmt.Sprint(team["team_name"].(string))
				}
			}
			var printable [22]string

			printable[0] = fmt.Sprint(teamIdHome)

			printable[1] = fmt.Sprint(teamIdAway)

			for index, last_5 := range standings {
				if teamIdHome == last_5[0] {
					home_team_position := index + 1

					printable[2] = fmt.Sprint(home_team_position)
					homeForm :=
						strings.Split(
							strings.ReplaceAll(
								strings.ReplaceAll(
									strings.ReplaceAll(
										strings.ReplaceAll(fmt.Sprint(last_5[1]), "W", "1"), "D", "0.5"), "L", "0"), " ", ""), "-")
					printable[4] = fmt.Sprint(matchweek)

					var homeLast1 string
					var homeLast2 string
					var homeLast3 string
					var homeLast4 string
					var homeLast5 string

					var lenHomeForm int = len(homeForm)
					if lenHomeForm == 6 {
						homeLast1 = homeForm[4]
						homeLast2 = homeForm[3]
						homeLast3 = homeForm[2]
						homeLast4 = homeForm[1]
						homeLast5 = homeForm[0]

					} else if lenHomeForm == 5 {
						homeLast1 = homeForm[3]
						homeLast2 = homeForm[2]
						homeLast3 = homeForm[1]
						homeLast4 = homeForm[0]
						homeLast5 = "0"

					} else if lenHomeForm == 4 {
						homeLast1 = homeForm[2]
						homeLast2 = homeForm[1]
						homeLast3 = homeForm[0]
						homeLast4 = "0"
						homeLast5 = "0"
					} else if lenHomeForm == 3 {
						homeLast1 = homeForm[1]
						homeLast2 = homeForm[0]
						homeLast3 = "0"
						homeLast4 = "0"
						homeLast5 = "0"
					} else if lenHomeForm == 2 {
						homeLast1 = homeForm[0]
						homeLast2 = "0"
						homeLast3 = "0"
						homeLast4 = "0"
						homeLast5 = "0"
					}
					printable[9] = homeLast1
					printable[10] = homeLast2
					printable[11] = homeLast3
					printable[12] = homeLast4
					printable[13] = homeLast5

				} else if teamIdAway == last_5[0] {
					away_team_position := index + 1
					printable[3] = fmt.Sprint(away_team_position)
					awayForm :=
						strings.Split(
							strings.ReplaceAll(
								strings.ReplaceAll(
									strings.ReplaceAll(
										strings.ReplaceAll(fmt.Sprint(last_5[1]), "W", "1"), "D", "0.5"), "L", "0"), " ", ""), "-")

					var awayLast1 string
					var awayLast2 string
					var awayLast3 string
					var awayLast4 string
					var awayLast5 string

					var lenAwayForm int = len(awayForm)

					if lenAwayForm == 6 {
						awayLast1 = awayForm[4]
						awayLast2 = awayForm[3]
						awayLast3 = awayForm[2]
						awayLast4 = awayForm[1]
						awayLast5 = awayForm[0]

					} else if lenAwayForm == 5 {
						awayLast1 = awayForm[3]
						awayLast2 = awayForm[2]
						awayLast3 = awayForm[1]
						awayLast4 = awayForm[0]
						awayLast5 = "0"
					} else if lenAwayForm == 4 {
						awayLast1 = awayForm[2]
						awayLast2 = awayForm[1]
						awayLast3 = awayForm[0]
						awayLast4 = "0"
						awayLast5 = "0"
					} else if lenAwayForm == 3 {
						awayLast1 = awayForm[1]
						awayLast2 = awayForm[0]
						awayLast3 = "0"
						awayLast4 = "0"
						awayLast5 = "0"
					} else if lenAwayForm == 2 {
						awayLast1 = awayForm[0]
						awayLast2 = "0"
						awayLast3 = "0"
						awayLast4 = "0"
						awayLast5 = "0"
					}

					printable[14] = awayLast1
					printable[15] = awayLast2
					printable[16] = awayLast3
					printable[17] = awayLast4
					printable[18] = awayLast5
				}
			}

			for _, gfga := range goals {
				if fmt.Sprint(gfga["team_id"]) == fmt.Sprint(teamIdHome) {
					homeGA := gfga["GA"]
					printable[6] = fmt.Sprint(homeGA)
					homeGF := gfga["GF"]
					printable[5] = fmt.Sprint(homeGF)
					// fmt.Printf("results type: %T\n", results)

				} else if fmt.Sprint(gfga["team_id"]) == fmt.Sprint(teamIdAway) {
					awayGA := gfga["GA"]
					printable[8] = fmt.Sprint(awayGA)
					awayGF := gfga["GF"]
					printable[7] = fmt.Sprint(awayGF)

				}

			}

			for _, result := range results.([]interface{}) {
				results_away_team := fmt.Sprint(result.(map[string]interface{})["away_team"])
				results_home_team := fmt.Sprint(result.(map[string]interface{})["home_team"])

				if results_away_team == away_team && results_home_team == home_team {
					homeScore, err := strconv.Atoi(interface{}(strings.Split(fmt.Sprint(result.(map[string]interface{})["result"]), ":")[0]).(string))
					handleError(err)
					awayScore, err := strconv.Atoi(interface{}(strings.Split(fmt.Sprint(result.(map[string]interface{})["result"]), ":")[1]).(string))
					handleError(err)
					printable[19] = fmt.Sprint(homeScore)
					printable[20] = fmt.Sprint(awayScore)

					if homeScore+awayScore > 3 { // over 3.5
						printable[21] = "1"
					} else {
						printable[21] = "0"
					}
				}
			}

			writeToCSV(printable[:])
		}

	}
}

func dataCollectorOne(teamIdsChan chan [10][2]interface{}, matchweekChan chan int) {
	/*
		Collects:
		hometeamid, awayteamid, matchweek
	*/

	gamesChan := make(chan [10][2]interface{})
	roundIdChan := make(chan int)

	go getGames(gamesChan)     //hometeamid, awayteamid
	go getRoundId(roundIdChan) //matchweek

	a := <-gamesChan
	b := <-roundIdChan
	teamIdsChan <- a
	matchweekChan <- b
}

func dataCollectorTwo(standings chan [20][2]interface{}) {
	/*
		Collects:
		hometeam league position, awayteam league position
		hometeam form, awayteam form
	*/
	standingsChan := make(chan [20][2]interface{})

	go getStandings(standingsChan) //teamId, //team form

	a := <-standingsChan

	standings <- a

}

func dataCollectorThree() (interface{}, []map[string]interface{}) {
	/*
		Collects:
		hometeam GF, hometeam GA, awayteam GF, awayteam GA, results
	*/

	goalsChan := make(chan []map[string]interface{})
	resultsChan := make(chan interface{})

	go getData(goalsChan, resultsChan) //GA, GF, team_id, team_name

	return <-resultsChan, <-goalsChan

}

func handleError(err error) {
	if err != nil {
		log.Fatalln("AN ERROR OCCURED: ", err)
		return
	}
}

func getGames(c chan [10][2]interface{}) {
	url := "https://odibets.com/api/fv"

	//Converting our map payload to json
	data, _ := json.Marshal(map[string]string{
		"competition_id": "1",
		"tab":            "",
		"period":         "",
		"sub_type_id":    "",
	})

	// Converting payload to bytes
	payload := bytes.NewBuffer(data)

	// Our Http post request
	response, err := http.Post(url, "application/json", payload)
	handleError(err)

	games := handleGetGamesResponse(response)

	c <- games

}

func handleGetGamesResponse(response *http.Response) [10][2]interface{} {
	// Reading the response in bytes
	responseBody, err := ioutil.ReadAll(response.Body)
	handleError(err)

	// Declaring our struct type where the response data will be unmarshalled.
	// Its called 'BigData' because a very big json object will be unmarshalled into it.
	type BigData struct {
		Status_code        float64
		Status_description string
		Data               map[string]interface{}
	}

	// Initializing/ Instanciating the struct BigData
	m := BigData{}

	// Unmarshalling the response data into the struct variable
	err = json.Unmarshal(responseBody, &m)

	handleError(err)

	dat := m.Data["matches"].([]interface{})

	var games [10][2]interface{}

	for i := range dat {
		var data [2]interface{}
		_dat := dat[i].(map[string]interface{})

		data[0] = _dat["away_id"]
		data[1] = _dat["home_id"]

		games[i] = data
	}

	return games

}

func getRoundId(c chan int) {

	url := "https://odibets.com/api/fv"

	//Converting our map payload to json
	data, _ := json.Marshal(map[string]string{
		"competition_id": "1",
		"tab":            "results",
		"period":         "2021-12-24 11:55:00",
		"sub_type_id":    "",
	})

	// Converting payload to bytes
	payload := bytes.NewBuffer(data)

	// Our Http post request
	response, err := http.Post(url, "application/json", payload)
	handleError(err)

	// Close the connection after everything completes
	defer response.Body.Close()

	// Reading the response in bytes
	responseBody, err := ioutil.ReadAll(response.Body)
	handleError(err)

	// Declaring our struct type where the response data will be unmarshalled.
	// Its called 'BigData' because a very big json object will be unmarshalled into it.
	type BigData struct {
		Status_code        float64
		Status_description string
		Data               map[string]interface{}
	}

	// Initializing/ Instanciating the struct BigData
	m := BigData{}

	// Unmarshalling the response data into the struct variable
	err = json.Unmarshal(responseBody, &m)
	handleError(err)

	// This section of code simply checks which round it is. Rounds run from
	// 1 through 38. Data is fetched when the round is 38.
	dat := m.Data["results"].([]interface{})
	_dat := dat[0].(map[string]interface{})
	__dat := fmt.Sprint(_dat["round_id"])
	round_id, err := strconv.Atoi(__dat)
	handleError(err)

	c <- round_id
}

func getStandings(c chan [20][2]interface{}) {
	url := "https://odibets.com/api/fv"

	//Converting our map payload to json
	data, _ := json.Marshal(map[string]string{
		"competition_id": "1",
		"tab":            "standings",
		"period":         "",
		"sub_type_id":    "",
	})

	// Converting payload to bytes
	payload := bytes.NewBuffer(data)

	// Our Http post request
	response, err := http.Post(url, "application/json", payload)
	handleError(err)

	standings := handleGetStandingsResponse(response)

	c <- standings

}

func handleGetStandingsResponse(response *http.Response) [20][2]interface{} {

	// Reading the response in bytes
	responseBody, err := ioutil.ReadAll(response.Body)
	handleError(err)

	// Declaring our struct type where the response data will be unmarshalled.
	// Its called 'BigData' because a very big json object will be unmarshalled into it.
	type BigData struct {
		Status_code        float64
		Status_description string
		Data               map[string]interface{}
	}

	// Initializing/ Instanciating the struct BigData
	m := BigData{}

	// Unmarshalling the response data into the struct variable
	err = json.Unmarshal(responseBody, &m)

	handleError(err)

	var standings [20][2]interface{}

	dat := m.Data["standings"].([]interface{})
	for i := range dat {
		var data [2]interface{}
		_dat := dat[i].(map[string]interface{})
		data[0] = _dat["team_id"]
		data[1] = _dat["team_form"]

		standings[i] = data
	}

	return standings

}

func getData(c chan []map[string]interface{}, ce chan interface{}) {
	url := "https://odibets.com/api/fv"

	//Converting our map payload to json
	data, _ := json.Marshal(map[string]string{
		"competition_id": "1",
		"tab":            "results",
		"period":         "",
		"sub_type_id":    "",
	})

	// Converting payload to bytes
	payload := bytes.NewBuffer(data)

	// Our Http post request
	response, err := http.Post(url, "application/json", payload)
	handleError(err)

	// Close the connection after everything completes
	defer response.Body.Close()

	// Reading the response in bytes
	responseBody, err := ioutil.ReadAll(response.Body)
	handleError(err)

	// Declaring our struct type where the response data will be unmarshalled.
	// Its called 'BigData' because a very big json object will be unmarshalled into it.
	type BigData struct {
		Status_code        float64
		Status_description string
		Data               map[string]interface{}
	}

	// Initializing/ Instanciating the struct BigData
	m := BigData{}

	// Unmarshalling the response data into the struct variable
	err = json.Unmarshal(responseBody, &m)
	handleError(err)

	// This section of code simply checks which round it is. Rounds run from
	// 1 through 38. Data is fetched when the round is 38.
	dat := m.Data["results"].([]interface{})
	_dat := dat[0].(map[string]interface{})
	__dat := fmt.Sprint(_dat["round_id"])
	round_id, err := strconv.Atoi(__dat)
	handleError(err)

	table := [...]map[string]interface{}{
		{"team_id": 1, "team_name": "Manchester Blue", "GF": 0, "GA": 0},
		{"team_id": 2, "team_name": "Manchester Reds", "GF": 0, "GA": 0},
		{"team_id": 3, "team_name": "Liverpool", "GF": 0, "GA": 0},
		{"team_id": 4, "team_name": "London Blues", "GF": 0, "GA": 0},
		{"team_id": 5, "team_name": "Tottenham", "GF": 0, "GA": 0},
		{"team_id": 6, "team_name": "London Reds", "GF": 0, "GA": 0},
		{"team_id": 7, "team_name": "Burnley", "GF": 0, "GA": 0},
		{"team_id": 8, "team_name": "Leicester", "GF": 0, "GA": 0},
		{"team_id": 9, "team_name": "Everton", "GF": 0, "GA": 0},
		{"team_id": 10, "team_name": "LEEDS", "GF": 0, "GA": 0},
		{"team_id": 11, "team_name": "WEST BROM", "GF": 0, "GA": 0},
		{"team_id": 12, "team_name": "West Ham", "GF": 0, "GA": 0},
		{"team_id": 13, "team_name": "Newcastle", "GF": 0, "GA": 0},
		{"team_id": 14, "team_name": "Brighton", "GF": 0, "GA": 0},
		{"team_id": 15, "team_name": "Palace", "GF": 0, "GA": 0},
		{"team_id": 16, "team_name": "FULHAM", "GF": 0, "GA": 0},
		{"team_id": 17, "team_name": "ASTON V", "GF": 0, "GA": 0},
		{"team_id": 18, "team_name": "Southampton", "GF": 0, "GA": 0},
		{"team_id": 19, "team_name": "Wolves", "GF": 0, "GA": 0},
		{"team_id": 20, "team_name": "SHEFFIELD U", "GF": 0, "GA": 0},
	}

	// When the round is 38, fetch data, sort it and write it to the matches.csv file
	if round_id > 0 {

		// This outer loop iterates through all the 38 matchweeks results in descending order. This is because
		// matchweek 1 resuls appear as the last element in the data, so starting from the bottom ensures that
		// the table is updated from matchweek 1.

		for j := round_id - 1; j >= 0; j-- {
			dat1 := dat[j].(map[string]interface{})

			if j == 0 {
				ce <- dat1["matches"] //write latest results into channel
			}

			// This inner loop iterates through all individual matches extracting the relevant data such as:
			// -> home team
			// -> aawy team
			// -> results

			for i := 0; i < 10; i++ {

				dat2 := dat1["matches"].([]interface{})

				stringData := dat2[i].(map[string]interface{})
				away_team := stringData["away_team"]
				home_team := stringData["home_team"]
				result := stringData["result"]

				// Splitting the results string so as to obtain the scores for individual teams that is
				// scores for the home team 'home_score' and those for the awayteam 'away_score'
				splitResult := strings.Split(fmt.Sprintf("%s", result), ":")

				// Converting the home_score string into an integer for manipulation
				home_score, err := strconv.Atoi(splitResult[0])
				handleError(err)

				// Converting the away_score string into an integer for manipulation
				away_score, err := strconv.Atoi(splitResult[1])
				handleError(err)

				for _, value := range table {
					if value["team_name"] == home_team {
						value["GF"] = value["GF"].(int) + home_score
						value["GA"] = value["GA"].(int) + away_score
					} else if value["team_name"] == away_team {
						value["GF"] = value["GF"].(int) + away_score
						value["GA"] = value["GA"].(int) + home_score
					}

				}

			}

		}

	}
	c <- table[:]
}

func writeToCSV(s []string) {

	// The file is opened in append mode
	f, err := os.OpenFile("final.csv", os.O_RDWR|os.O_APPEND, os.ModeAppend)
	handleError(err)

	// Writer variable
	w := csv.NewWriter(f)

	err = w.Write(s)
	handleError(err)

	w.Flush()
	err = w.Error()
	handleError(err)

	f.Close()
	handleError(err)

}
