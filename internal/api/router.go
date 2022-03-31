package api

import "github.com/gin-gonic/gin"

type App struct {
	Router *gin.Engine
}

func (app *App) InitRoutes(r *gin.Engine) {
	r.POST("/groups", searchInGroups)
}
