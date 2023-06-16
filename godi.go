package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
//
// This program fetches data from the odibets website for analysis. It arranges the data
// neatly and writes it into a csv file.
//
*/

// session variable
var response http.Response

// Main function
func main() {
	login()
	getData()

}

// This function handles all errors
func handleError(err error) {
	if err != nil {
		fmt.Println("\n ")
		log.Fatal("AN ERROR OCCURED!!: ", err)
	}
}

func login() {
	// login url
	url1 := "https://odibets.com/api/va"

	// login post data
	data1, _ := json.Marshal(map[string]string{
		"msisdn": "your mobile number",
		"pwd":    "your password",
	})

	// Converting payload to bytes
	payload1 := bytes.NewBuffer(data1)

	// The Http post request
	response, err := http.Post(url1, "application/json", payload1)
	handleError(err)

	// Reading the response data
	responseBody1, err := ioutil.ReadAll(response.Body)
	handleError(err)

	// Declaring the map variable where the data will be Unmarshalled
	var s map[string]interface{}

	// Unmarshalling the response data into the map variable
	err = json.Unmarshal(responseBody1, &s)
	handleError(err)

	log.Println(s["status_description"])
}

/*
// The function bellow is an implemaentation of Quick Sort which sorts teams
// basing on the points gained after a match and the goal difference as well
*/
func quicksortTeams(table []map[string]interface{}) []map[string]interface{} {

	if len(table) < 2 {
		return table
	}

	left, right := 0, len(table)-1

	pivot := rand.Int() % len(table)

	table[pivot], table[right] = table[right], table[pivot]

	for i := range table {

		// Swap teams positions if points are less than the team on the right
		if table[i]["pts"].(int) < table[right]["pts"].(int) {
			table[left], table[i] = table[i], table[left]
			left++
		}

		// Swap teams positions if the teams have equal points but goal differences differ i.e.
		// if the goal difference is less than that of the team on the right
		if table[i]["pts"].(int) == table[right]["pts"].(int) {
			if table[i]["GD"].(int) < table[right]["GD"].(int) {
				table[left], table[i] = table[i], table[left]
				left++
			}
		}

	}

	table[left], table[right] = table[right], table[left]

	// This is a recursive algorithm
	// Do quick sort on the left array after splitting
	quicksortTeams(table[:left])

	// Do quick sort on the right array after splitting
	quicksortTeams(table[left+1:])

	return table
}

/*
// The qucksortTeams() function sorts teams and arranges them in aascending order as
// in the team with least points and goal difference appears at the top.
// This function basically reverses the sorted table so that its in descending
// order, i.e. the team with the highest points and goal differnce appaears at
// the top.
*/
func reverseTable(s []map[string]interface{}) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// This function writes data to the csv file known as data.csv
func writeToCSV(s []string) {

	// The file is opened in append mode
	f, err := os.OpenFile("matches.csv", os.O_RDWR|os.O_APPEND, os.ModeAppend)
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

func getData() {
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

	// When the round is 38, fetch data, sort it and write it to the matches.csv file
	if round_id == 38 {
		log.Println("Writing to matches.csv..")

		// Declaring our standings table array. Initially, the teams are arranged in
		// alphabetic order.
		table1 := [...]map[string]interface{}{
			{"team": "ASTON V", "GD": 0, "pts": 0},
			{"team": "Brighton", "GD": 0, "pts": 0},
			{"team": "Burnley", "GD": 0, "pts": 0},
			{"team": "Everton", "GD": 0, "pts": 0},
			{"team": "FULHAM", "GD": 0, "pts": 0},
			{"team": "LEEDS", "GD": 0, "pts": 0},
			{"team": "Leicester", "GD": 0, "pts": 0},
			{"team": "Liverpool", "GD": 0, "pts": 0},
			{"team": "London Blues", "GD": 0, "pts": 0},
			{"team": "London Reds", "GD": 0, "pts": 0},
			{"team": "Manchester Blue", "GD": 0, "pts": 0},
			{"team": "Manchester Reds", "GD": 0, "pts": 0},
			{"team": "Newcastle", "GD": 0, "pts": 0},
			{"team": "Palace", "GD": 0, "pts": 0},
			{"team": "SHEFFIELD U", "GD": 0, "pts": 0},
			{"team": "Southampton", "GD": 0, "pts": 0},
			{"team": "Tottenham", "GD": 0, "pts": 0},
			{"team": "WEST BROM", "GD": 0, "pts": 0},
			{"team": "West Ham", "GD": 0, "pts": 0},
			{"team": "Wolves", "GD": 0, "pts": 0},
		}

		// Declaring our table slice from the table array. It needs to be a slice since that's the datatype
		// of the formal parameters for the quicksortTeams() function where it will be passed.
		table := table1[:]

		team := table

		// Write the initial table to the csv file
		// for _, value := range table {
		// 	team := []string{value["team"].(string), fmt.Sprint(value["GD"].(int)), fmt.Sprint(value["pts"].(int))}
		// 	writeToCSV(team)
		// }

		over_one_point_five_accuracy_counter1 := [...]map[string]interface{}{
			{"position": 1, "counter": 0.0},
			{"position": 2, "counter": 0.0},
			{"position": 3, "counter": 0.0},
			{"position": 4, "counter": 0.0},
			{"position": 5, "counter": 0.0},
			{"position": 6, "counter": 0.0},
			{"position": 7, "counter": 0.0},
			{"position": 8, "counter": 0.0},
			{"position": 9, "counter": 0.0},
			{"position": 10, "counter": 0.0},
			{"position": 11, "counter": 0.0},
			{"position": 12, "counter": 0.0},
			{"position": 13, "counter": 0.0},
			{"position": 14, "counter": 0.0},
			{"position": 15, "counter": 0.0},
			{"position": 16, "counter": 0.0},
			{"position": 17, "counter": 0.0},
			{"position": 18, "counter": 0.0},
			{"position": 19, "counter": 0.0},
			{"position": 20, "counter": 0.0},
		}

		over_one_point_five_accuracy_counter := over_one_point_five_accuracy_counter1[:]

		no_gg_accuracy_counter1 := [...]map[string]interface{}{
			{"position": 1, "counter": 0.0, "goals": 0.0},
			{"position": 2, "counter": 0.0, "goals": 0.0},
			{"position": 3, "counter": 0.0, "goals": 0.0},
			{"position": 4, "counter": 0.0, "goals": 0.0},
			{"position": 5, "counter": 0.0, "goals": 0.0},
			{"position": 6, "counter": 0.0, "goals": 0.0},
			{"position": 7, "counter": 0.0, "goals": 0.0},
			{"position": 8, "counter": 0.0, "goals": 0.0},
			{"position": 9, "counter": 0.0, "goals": 0.0},
			{"position": 10, "counter": 0.0, "goals": 0.0},
			{"position": 11, "counter": 0.0, "goals": 0.0},
			{"position": 12, "counter": 0.0, "goals": 0.0},
			{"position": 13, "counter": 0.0, "goals": 0.0},
			{"position": 14, "counter": 0.0, "goals": 0.0},
			{"position": 15, "counter": 0.0, "goals": 0.0},
			{"position": 16, "counter": 0.0, "goals": 0.0},
			{"position": 17, "counter": 0.0, "goals": 0.0},
			{"position": 18, "counter": 0.0, "goals": 0.0},
			{"position": 19, "counter": 0.0, "goals": 0.0},
			{"position": 20, "counter": 0.0, "goals": 0.0},
		}

		no_gg_accuracy_counter := no_gg_accuracy_counter1[:]

		var total_goals_count int

		// This outer loop iterates through all the 38 matchweeks results in descending order. This is because
		// matchweek 1 resuls appear as the last element in the data, so starting from the bottom ensures that
		// the table is updated from matchweek 1.

		for j := 37; j >= 0; j-- {
			dat1 := dat[j].(map[string]interface{})

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

				// This section of code obtains goal difference values for the individual teams
				// by subtracting the score. Note that goal difference values can be negative numbers.
				homeGD := home_score - away_score
				awayGD := away_score - home_score

				over_or_under := home_score + away_score
				var over_one_point_five string
				if over_or_under > 1 {
					over_one_point_five = "over1.5"
				} else {
					over_one_point_five = "under1.5"
				}

				var gg_no_gg string
				if home_score > 0 && away_score > 0 {
					gg_no_gg = "gg"
				} else {
					gg_no_gg = "nogg"
				}

				var home_team_position int
				var away_team_position int

				for i, value := range team {
					if value["team"] == home_team {
						home_team_position = i + 1
					} else if value["team"] == away_team {
						away_team_position = i + 1
					}

				}

				for _, value := range over_one_point_five_accuracy_counter {
					if value["position"] == home_team_position {
						if over_one_point_five == "over1.5" {
							value["counter"] = value["counter"].(float64) + 1
						}
					} else if value["position"] == away_team_position {
						if over_one_point_five == "over1.5" {
							value["counter"] = value["counter"].(float64) + 1
						}
					}

				}

				for _, value := range no_gg_accuracy_counter {
					if value["position"] == home_team_position {
						if gg_no_gg == "gg" {
							value["counter"] = value["counter"].(float64) + 1
						}
						value["goals"] = value["goals"].(float64) + float64(home_score)
					} else if value["position"] == away_team_position {
						if gg_no_gg == "gg" {
							value["counter"] = value["counter"].(float64) + 1
						}
						value["goals"] = value["goals"].(float64) + float64(away_score)
					}

				}

				total_goals_count += home_score
				total_goals_count += away_score

				// This section updates the table by updating the values for points and goal difference
				// for individual teams
				for _, value := range table {
					if value["team"] == home_team {
						value["GD"] = value["GD"].(int) + homeGD

						if homeGD == awayGD {
							value["pts"] = value["pts"].(int) + 1
						} else if homeGD > awayGD {
							value["pts"] = value["pts"].(int) + 3
						}

					} else if value["team"] == away_team {
						value["GD"] = value["GD"].(int) + awayGD

						if awayGD == homeGD {
							value["pts"] = value["pts"].(int) + 1
						} else if awayGD > homeGD {
							value["pts"] = value["pts"].(int) + 3
						}
					}
				}

				// Declaring and initializing the array variable containing scores data to be written
				// to the csv file
				//team := []string{home_team.(string), strings.Replace(fmt.Sprint(result), ":", "-", 1), away_team.(string), fmt.Sprint(home_team_position), fmt.Sprint(away_team_position), over_one_point_five, gg_no_gg}
				// Writing the scores data to the csv file.
				//writeToCSV(team)

			}

			// After updating the scores for the individual teams, the table now needs to be sorted and re-arranged
			// which is exactly what the next section of code below does.
			table = quicksortTeams(table)
			reverseTable(table)

			// The updated table is written to the csv file with the loop below
			// for _, value := range table {
			// 	team := []string{value["team"].(string), fmt.Sprint(value["GD"].(int)), fmt.Sprint(value["pts"].(int))}
			// 	writeToCSV(team)
			// }

		}

		writeToCSV([]string{"\n"})
		writeToCSV([]string{"over1.5 accuracy percentage/position"})
		for _, value := range over_one_point_five_accuracy_counter {
			team_name := fmt.Sprint(value["position"])
			accuracy := fmt.Sprint((value["counter"].(float64) / 38) * 100)
			profit := fmt.Sprint("sh.", (value["counter"].(float64)*29)-((38-value["counter"].(float64))*100))

			accuracy_percentage := []string{team_name, accuracy, profit}
			writeToCSV(accuracy_percentage)
		}

		writeToCSV([]string{"\n"})
		writeToCSV([]string{"gg accuracy percentage/position"})
		for _, value := range no_gg_accuracy_counter {
			position := fmt.Sprint(value["position"])
			goals := fmt.Sprint(value["goals"])
			accuracy := fmt.Sprint((value["counter"].(float64) / 38) * 100)
			profit := fmt.Sprint("sh.", ((38-value["counter"].(float64))*100)-(value["counter"].(float64)*100))

			accuracy_percentage := []string{position, goals, accuracy, profit}
			writeToCSV(accuracy_percentage)
		}

		writeToCSV([]string{"\n"})
		total_goals := []string{"total_goals", fmt.Sprint(total_goals_count)}
		writeToCSV(total_goals)

		// Includes an empty entry into the csv file after every season.
		writeToCSV([]string{"\n\n"})

		log.Println("Written to matches.csv.")
		log.Println("Waiting for next round.")

		// Program waits for 120 seconds before the recurrsive call so as to prevent writing to the csv file
		// multiple times since the condition for the matchweek being 38 will still be true.
		time.Sleep(120 * time.Second)
		getData()

		// This condition checks whether the matchweek is less than 38. If true, it waits for 120 seconds
		// then makes the recurrsive call for the next matchweek.
	} else if round_id < 38 {
		sleepTime := 120
		log.Println("It's matchweek ", (round_id), ", Waiting for ", sleepTime, " seconds.")
		time.Sleep(time.Duration(sleepTime) * time.Second)
		getData()
	}

}
