package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func DelFavorite(targetType, targetName string) error {
	data, err := ioutil.ReadFile("favorites.json")
	if err != nil {
		return err
	}

	var jsonData map[string][]map[string]string
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	switch targetType {
	case "agent":
		for i, agent := range jsonData["agent"] {
			if agent["nom"] == targetName {
				jsonData["agent"] = append(jsonData["agent"][:i], jsonData["agent"][i+1:]...)
				break
			}
		}
	case "map":
		for i, m := range jsonData["map"] {
			if m["nom"] == targetName {
				jsonData["map"] = append(jsonData["map"][:i], jsonData["map"][i+1:]...)
				break
			}
		}
	default:
		return fmt.Errorf("Type cible non pris en charge: %s", targetType)
	}

	updatedData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("favorites.json", updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
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

func IsFavorite(targetType, targetName string) bool {
	data, err := ioutil.ReadFile("favorites.json")
	if err != nil {
		return false
	}

	var jsonData map[string][]map[string]string
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return false
	}

	switch targetType {
	case "agent":
		for _, agent := range jsonData["agent"] {
			if agent["nom"] == targetName {
				return true
			}
		}
	case "map":
		for _, m := range jsonData["map"] {
			if m["nom"] == targetName {
				return true
			}
		}
	default:
		return false
	}

	return false
}
