package main

import "github.com/Savilka/social-media-monitoring/internal/api"

func main() {
	app := api.App{}
	app.InitRouter()
	app.Run()
}
