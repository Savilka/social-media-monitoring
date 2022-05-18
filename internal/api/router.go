package api

import (
	"github.com/Savilka/social-media-monitoring/internal/api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

type App struct {
	Router *gin.Engine
}

func (app *App) InitRouter() {
	app.Router = gin.Default()
	app.Router.Use(cors.Default())
	app.Router.POST("/groups", handlers.WallHandler)
	app.Router.POST("/comment", handlers.CommentsHandler)
}

func (app *App) Run() {
	err := godotenv.Load("./conf/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
