package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAgents() AgentResponse {
	fmt.Println("Récupération des agents")
	resp, err := http.Get("https://valorant-api.com/v1/agents?language=fr-FR&isPlayableCharacter=true")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse AgentResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return apiResponse
}

func GetMaps() MapsData {
	fmt.Println("Récupération des cartes")
	resp, err := http.Get("https://valorant-api.com/v1/maps?language=fr-FR")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse MapsData
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return apiResponse
}

func GetMap(id string) UniqueMap {
	fmt.Println("Récupération de la carte")

	url := "https://valorant-api.com/v1/maps/" + id + "?language=fr-FR"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// Vérifiez si le statut est 200 OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-200 status code: %d", resp.StatusCode)
	}

	// Lisez les données de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var apiResponse UniqueMap
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return apiResponse
}

func GetWeapons() Weapons {
	fmt.Println("Récupération des armes")
	resp, err := http.Get("https://valorant-api.com/v1/weapons")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse Weapons
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return apiResponse
}

func GetGameMods() GameModeData {
	fmt.Println("Récupération des modes de jeu")
	resp, err := http.Get("https://valorant-api.com/v1/gamemodes")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse GameModeData
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return apiResponse
}
