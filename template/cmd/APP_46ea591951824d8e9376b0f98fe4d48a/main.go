package main

import (
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/cmd/APP_46ea591951824d8e9376b0f98fe4d48a/app"
	"log"
)

func main() {

	cmd := app.NewAPIServerCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
