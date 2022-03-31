package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetScreenName(url string) (string, error) {
	var screenName string
	if url[0:7] == "vk.com/" {
		screenName = url[7:]
	} else {
		if url[0:15] == "https://vk.com/" {
			screenName = url[15:]
		} else {
			return "", errors.New("not valid url")
		}
	}
	return screenName, nil
}

type Response struct {
	Response ResponseObj `json:"response"`
}

type ResponseObj struct {
	ObjectId int    `json:"object_id"`
	Type     string `json:"type"`
}

func GetId(screenName string) int {
	url := fmt.Sprintf("https://api.vk.com/method/utils.resolveScreenName?screen_name=%s&access_token=%s&v=%f",
		screenName, os.Getenv("ACCESS_TOKEN"), 5.131)

	res, err := http.Get(url)
	if err != nil {
		return 0
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0
	}

	err = res.Body.Close()
	if err != nil {
		return 0
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0
	}

	return response.Response.ObjectId
}
