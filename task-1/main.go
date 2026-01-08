package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	TotalData  int     `json:"total_data"`
	TotalPages int     `json:"total_pages"`
	Data       []Match `json:"data"`
}

type Match struct {
	Competition string `json:"competition"`
	Year        int    `json:"year"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1goals  string `json:"team1goals"`
	Team2goals  string `json:"team2goals"`
}

func getUrl(team string, name string, year int, page int) string {
	return fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?year=%v&%v=%v&page=%v", year, team, name, page)
}

func sumTotalGoals(team string, name string, year int) int {
	pageNumber := 1
	totalGoals := 0

	url := getUrl(team, name, year, pageNumber)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	var result Response
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return 0
	}

	for i := 1; i <= result.TotalPages; i++ {
		urlTeam1 := getUrl(team, name, year, i)
		response, err := http.Get(urlTeam1)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Error reading body: %v", err)
		}

		var result Response
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return 0
		}

		for _, match := range result.Data {
			if team == "team1" {
				goals, _ := strconv.Atoi(match.Team1goals)
				totalGoals += goals
			} else {
				goals, _ := strconv.Atoi(match.Team2goals)
				totalGoals += goals
			}
		}
	}

	return totalGoals
}

func main() {
	const FirstTeam = "team1"
	const SecondTeam = "team2"

	var teamName string
	var year int

	fmt.Println("Enter team name and year:")
	fmt.Println("Example: Barcelona 2011")
	_, err := fmt.Scan(&teamName, &year)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	totalGoalsTeam1 := sumTotalGoals(FirstTeam, teamName, year)
	totalGoalsTeam2 := sumTotalGoals(SecondTeam, teamName, year)

	fmt.Println(totalGoalsTeam1 + totalGoalsTeam2)

}
