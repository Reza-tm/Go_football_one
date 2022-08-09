package main

import team_getter "github.com/reza_tm/football/pkg/repository/one_football"

func main() {
	teams := team_getter.GetAllTeams()
	team_getter.GetPlayersWithSort(teams)
}
