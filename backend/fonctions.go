package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func GetSkin(id string) WeaponSkinData {
	fmt.Println("Récupération du skins")

	url := "https://valorant-api.com/v1/weapons/skins/" + id

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-200 status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var apiResponse WeaponSkinData
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return apiResponse
}

func AjouterElement(t string, nom string, imageURL string) error {
	var data Data
	file, err := os.Open("favorites.json")
	defer file.Close()
	if err == nil {
		err = json.NewDecoder(file).Decode(&data)
		if err != nil {
			return err
		}
	}

	// Vérifier si l'élément existe déjà
	switch t {
	case "agent":
		for _, agent := range data.Agents {
			if agent.Nom == nom && agent.ImageURL == imageURL {
				fmt.Println("L'agent existe déjà dans la liste.")
				return nil
			}
		}
		data.Agents = append(data.Agents, Element{Nom: nom, ImageURL: imageURL})
	case "map":
		for _, elementMap := range data.Maps {
			if elementMap.Nom == nom && elementMap.Id == imageURL {
				fmt.Println("La carte existe déjà dans la liste.")
				return nil
			}
		}
		data.Maps = append(data.Maps, ElementMap{Nom: nom, Id: imageURL})
	default:
		return fmt.Errorf("Type non valide")
	}

	file, err = os.Create("favorites.json")
	defer file.Close()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func GetFavorites() (Data, error) {
	var data Data

	file, err := os.Open("favorites.json")
	if err != nil {
		return data, fmt.Errorf("erreur lors de l'ouverture du fichier : %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return data, fmt.Errorf("erreur lors de la lecture du fichier JSON : %v", err)
	}

	return data, nil
}
