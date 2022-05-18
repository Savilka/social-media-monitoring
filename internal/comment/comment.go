package comment

import (
	"encoding/json"
	"fmt"
	"github.com/Savilka/social-media-monitoring/internal/model"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func SearchInComments(wg *sync.WaitGroup, ch chan []model.Comment, ownerId int, query string) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.vk.com/method/wall.get?owner_id=%d&count=%d&access_token=%s&v=%f",
		ownerId, 100, os.Getenv("ACCESS_TOKEN"), 5.131)

	res, err := http.Get(url)
	if err != nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = res.Body.Close()
	if err != nil {
		return
	}

	var response struct {
		Response model.PostResponse `json:"response"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	var satisfyingComments []model.Comment

	var commentResponse struct {
		Response model.CommentsResponse `json:"response"`
	}
	for _, item := range response.Response.Items {
		url := fmt.Sprintf("https://api.vk.com/method/wall.getComments?owner_id=%d&count=%d&post_id=%d&access_token=%s&v=%f",
			ownerId, 100, item.Id, os.Getenv("ACCESS_TOKEN"), 5.131)

		commentsRes, err := http.Get(url)
		if err != nil {
			return
		}

		body, err = io.ReadAll(commentsRes.Body)
		if err != nil {
			return
		}

		err = res.Body.Close()
		if err != nil {
			return
		}

		err = json.Unmarshal(body, &commentResponse)
		if err != nil {
			return
		}

		for _, comment := range commentResponse.Response.Items {
			if strings.Contains(comment.Text, query) {
				comment.Link = fmt.Sprintf("vk.com/wall-%d_%d?reply=%d", -ownerId, item.Id, comment.Id)
				satisfyingComments = append(satisfyingComments, comment)
			}
		}
	}

	ch <- satisfyingComments
}
