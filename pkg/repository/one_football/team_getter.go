package team_getter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/fatih/color"
	structs "github.com/reza_tm/football/model"
)

type Response struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Data   structs.DataStruct `json:"data"`
}

var neededTeams = []string{"Germany", "England", "France", "Spain", "Manchester United", "Arsenal", "Chelsea", "Barcelona", "Real Madrid", "Bayern Munich"}

func getTeam(id int) (team structs.Team) {
	apiUrl := fmt.Sprintf("https://api-origin.onefootball.com/score-one-proxy/api/teams/en/%v.json", id)
	resp, _ := http.Get(apiUrl)
	var teamInfo Response
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &teamInfo)
	team = teamInfo.Data.Team
	return
}

func GetAllTeams() (teams []structs.Team) {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	teams = []structs.Team{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			team := getTeam(id)
			for _, v := range neededTeams {
				mtx.Lock()
				if v == team.Name {
					fmt.Println(team.Name)
					teams = append(teams, team)
				}
				mtx.Unlock()
			}
		}(i)
	}
	wg.Wait()
	return
}

func GetPlayersWithSort(teams []structs.Team) (players []structs.Player) {
	players = []structs.Player{}
	for _, v := range teams {
		players = append(players, v.Players...)
	}

	sort.SliceStable(players, func(i, j int) bool {
		iNum, _ := strconv.Atoi(players[i].ID)
		jNum, _ := strconv.Atoi(players[j].ID)
		return iNum < jNum
	})

	r := color.New(color.FgRed)
	b := color.New(color.FgBlue)
	r.Printf("+%s+\n", strings.Repeat("-", 78))
	for i, v := range players {
		b.Printf("| name : %-40s | age: %-7s | id: %-7s | \n", v.Name, v.Age, v.ID)
		if i != len(players)-1 {
			r.Printf("|%s|\n", strings.Repeat("-", 78))
		}

	}
	r.Printf("+%s+\n", strings.Repeat("-", 78))
	return
}
