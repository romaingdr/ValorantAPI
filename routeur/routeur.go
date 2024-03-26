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

	// Démarrage du serveur
	log.Println("Serveur lancé")
	fmt.Println("http://localhost:8080/accueil")
	http.ListenAndServe(":8080", nil)
	log.Fatal()
}
