package handlers

import (
	"github.com/Savilka/social-media-monitoring/internal/groups"
	"github.com/Savilka/social-media-monitoring/internal/model"
	"github.com/Savilka/social-media-monitoring/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"sync"
)

func SearchInGroups(c *gin.Context) {
	var req model.Request
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reg, err := regexp.Compile("(http?://)?(m.)?(vk.com/)")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var posts []model.Post
	wg := sync.WaitGroup{}
	ch := make(chan []model.Post, len(req.Links))
	for _, link := range req.Links {
		screenName, err := utils.GetScreenName(link, reg)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		id := utils.GetId(screenName)
		if id == 0 {
			continue
		}

		wg.Add(1)
		go groups.SearchInGroup(&wg, -id, ch, req.Text)
	}
	wg.Wait()
	for i := 0; i < len(req.Links); i++ {
		posts = append(posts, <-ch...)
	}
	c.JSON(http.StatusOK, posts)
}
