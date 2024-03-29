package controller

import (
	"ValorantAPI/backend"
	"ValorantAPI/templates"
	"fmt"
	"net/http"
	"strings"
)

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}

func AgentsPage(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")

	var agents = backend.GetAgents()

	fmt.Println(filter)

	var filteredAgents [][]backend.Character

	var filtered []backend.Character
	if filter == "all" {
		filtered = agents.Data
	} else {
		for _, agent := range agents.Data {
			if agent.Role.DisplayName == filter {
				filtered = append(filtered, agent)
			}
		}
	}

	var currentGroup []backend.Character
	for _, agent := range filtered {
		currentGroup = append(currentGroup, agent)
		if len(currentGroup) == 4 {
			filteredAgents = append(filteredAgents, currentGroup)
			currentGroup = nil
		}
	}
	if len(currentGroup) > 0 {
		filteredAgents = append(filteredAgents, currentGroup)
	}

	type allAgents struct {
		Agents [][]backend.Character
		Filter string
	}

	templates.Temp.ExecuteTemplate(w, "agents", allAgents{filteredAgents, filter})
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

func AddFavPage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("key")
	t := r.URL.Query().Get("type")

	var imageURL string
	var id string

	if t == "agent" {
		imageURL = r.URL.Query().Get("img")
		err := backend.AjouterElement(t, name, imageURL)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		id = r.URL.Query().Get("id")
		err := backend.AjouterElement(t, name, id)
		if err != nil {
			fmt.Println(err)
		}
	}

	if t == "agent" {
		http.Redirect(w, r, "/agent?name="+name, http.StatusSeeOther)
	} else if t == "map" {
		http.Redirect(w, r, "/map?id="+id, http.StatusSeeOther)
	}
}

func FavoritesPage(w http.ResponseWriter, r *http.Request) {
	data, err := backend.GetFavorites()
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	fmt.Println(data)
	templates.Temp.ExecuteTemplate(w, "favorites", data)
}

func SearchPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	query := r.FormValue("query")

	agents := backend.GetAgents()
	maps := backend.GetMaps()
	weapons := backend.GetWeapons()

	var filteredAgents []backend.Character
	var filteredMaps []backend.Map
	var filteredWeapons []backend.Item

	// Filtrer les agents
	for _, agent := range agents.Data {
		if strings.Contains(strings.ToLower(agent.DisplayName), strings.ToLower(query)) {
			filteredAgents = append(filteredAgents, agent)
		}
	}

	// Filtrer les cartes
	for _, m := range maps.Data {
		if strings.Contains(strings.ToLower(m.DisplayName), strings.ToLower(query)) {
			filteredMaps = append(filteredMaps, m)
		}
	}

	// Filtrer les armes
	for _, weapon := range weapons.Data {
		if strings.Contains(strings.ToLower(weapon.DisplayName), strings.ToLower(query)) {
			filteredWeapons = append(filteredWeapons, weapon)
		}
	}

	searchResults := struct {
		Agents  []backend.Character
		Maps    []backend.Map
		Weapons []backend.Item
	}{
		Agents:  filteredAgents,
		Maps:    filteredMaps,
		Weapons: filteredWeapons,
	}

	templates.Temp.ExecuteTemplate(w, "search", searchResults)
}
