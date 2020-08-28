package main

import (
	"{{cookiecutter.project_slug}}/cmd/{{cookiecutter.app_slug}}/app"
	"log"
)

func main() {

	cmd := app.NewAPIServerCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
