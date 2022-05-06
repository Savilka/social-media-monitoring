package main

import (
	"github.com/Savilka/social-media-monitoring/internal/api"
	_ "net/http/pprof"
)

func main() {
	app := api.App{}
	app.InitRouter()
	app.Run()
}
