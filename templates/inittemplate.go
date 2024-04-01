package templates

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

var templateFuncs = template.FuncMap{
	"add": func(a, b int) int { return a + b },
	"sub": func(a, b int) int { return a - b },
	"div": func(a, b int) int { return a / b },
}

func InitTemplate() {
	temp, err := template.New("").Funcs(templateFuncs).ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf("Erreur template: %v\n", err.Error())
		os.Exit(1)
	}
	Temp = temp
}
