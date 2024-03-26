package main

import (
	"ValorantAPI/routeur"
	"ValorantAPI/templates"
)

func main() {
	templates.InitTemplate()
	routeur.Initserv()
}
