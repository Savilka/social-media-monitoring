package api

import (
	"github.com/Savilka/social-media-monitoring/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	Router *gin.Engine
}

func (app *App) InitRouter() {
	app.Router = gin.Default()

	app.Router.POST("/groups", handlers.SearchInGroups)
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
