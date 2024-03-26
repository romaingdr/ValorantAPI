package controller

import (
	"ValorantAPI/backend"
	"ValorantAPI/templates"
	"fmt"
	"net/http"
)

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}

func AgentsPage(w http.ResponseWriter, r *http.Request) {

	var agents = backend.GetAgents()

	var groupedAgents [][]backend.Character
	for i := 0; i < len(agents.Data); i += 4 {
		end := i + 4
		if end > len(agents.Data) {
			end = len(agents.Data)
		}
		groupedAgents = append(groupedAgents, agents.Data[i:end])
	}

	templates.Temp.ExecuteTemplate(w, "agents", groupedAgents)
}

func AgentPage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	var agents = backend.GetAgents()

	var agent []backend.Character
	for _, character := range agents.Data {
		if character.DisplayName == name {
			agent = append(agent, character)
			break
		}
	}

	templates.Temp.ExecuteTemplate(w, "agent", agent[0])
}

func MapsPage(w http.ResponseWriter, r *http.Request) {

	maps := backend.GetMaps()

	var groupedMaps [][]backend.Map
	for i := 0; i < len(maps.Data); i += 4 {
		end := i + 4
		if end > len(maps.Data) {
			end = len(maps.Data)
		}
		groupedMaps = append(groupedMaps, maps.Data[i:end])
	}

	templates.Temp.ExecuteTemplate(w, "maps", groupedMaps)
}

func MapPage(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	mapDatas := backend.GetMap(id)

	fmt.Println(mapDatas)

	templates.Temp.ExecuteTemplate(w, "map", mapDatas.Datas)
}
