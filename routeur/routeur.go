package routeur

import (
	"ValorantAPI/controller"
	"fmt"
	"log"
	"net/http"
)

func Initserv() {

	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	// Accueil
	http.HandleFunc("/accueil", controller.AccueilPage)

	// Agents
	http.HandleFunc("/agents", controller.AgentsPage)
	http.HandleFunc("/agent", controller.AgentPage)

	// Maps
	http.HandleFunc("/maps", controller.MapsPage)
	http.HandleFunc("/map", controller.MapPage)

	// Weapons
	http.HandleFunc("/weapons", controller.WeaponsPage)
	http.HandleFunc("/weapon", controller.WeaponPage)

	// Skins
	http.HandleFunc("/skin", controller.SkinPage)

	// GamesMods
	http.HandleFunc("/gamemods", controller.GamemodsPage)

	// Favoris
	http.HandleFunc("/addfav", controller.AddFavPage)
	http.HandleFunc("/favorites", controller.FavoritesPage)

	// Recherche
	http.HandleFunc("/search", controller.SearchPage)

	// Erreur 404
	http.HandleFunc("/", controller.ErrorPage)

	// Démarrage du serveur
	log.Println("Serveur lancé")
	fmt.Println("http://localhost:8081/accueil")
	http.ListenAndServe(":8081", nil)
	log.Fatal()
}
