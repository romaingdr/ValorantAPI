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

func WeaponsPage(w http.ResponseWriter, r *http.Request) {

	weapons := backend.GetWeapons()

	var groupedWeapons [][]backend.Item
	for i := 0; i < len(weapons.Data); i += 4 {
		end := i + 4
		if end > len(weapons.Data) {
			end = len(weapons.Data)
		}
		groupedWeapons = append(groupedWeapons, weapons.Data[i:end])
	}

	templates.Temp.ExecuteTemplate(w, "weapons", groupedWeapons)
}

func WeaponPage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	var weapons = backend.GetWeapons()

	var weaponSelected []backend.Item
	for _, weapon := range weapons.Data {
		if weapon.DisplayName == name {
			weaponSelected = append(weaponSelected, weapon)
			break
		}
	}

	templates.Temp.ExecuteTemplate(w, "weapon", weaponSelected[0])
}

func GamemodsPage(w http.ResponseWriter, r *http.Request) {
	gamesMods := backend.GetGameMods()

	var groupedGameModes [][]backend.GameMode
	for i := 0; i < len(gamesMods.Data); i += 4 {
		end := i + 4
		if end > len(gamesMods.Data) {
			end = len(gamesMods.Data)
		}
		groupedGameModes = append(groupedGameModes, gamesMods.Data[i:end])
	}

	templates.Temp.ExecuteTemplate(w, "gamemods", groupedGameModes)
}

func SkinPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	skin := backend.GetSkin(id)

	templates.Temp.ExecuteTemplate(w, "skin", skin)
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "error", nil)
}
