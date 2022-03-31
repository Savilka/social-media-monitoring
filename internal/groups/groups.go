package groups

import (
	"encoding/json"
	"fmt"
	"github.com/Savilka/social-media-monitoring/internal/model"
	"io"
	"net/http"
	"os"
	"sync"
)

func SearchInGroup(wg *sync.WaitGroup, ownerId int, ch chan []model.Post, query string) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.vk.com/method/wall.search?owner_id=%d&query=%s&access_token=%s&v=%f",
		ownerId, query, os.Getenv("ACCESS_TOKEN"), 5.131)

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

	ch <- response.Response.Items
}
