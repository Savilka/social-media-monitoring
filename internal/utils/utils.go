package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func GetScreenName(url string, reg *regexp.Regexp) (string, error) {
	splitUrl := reg.Split(url, -1)
	if len(splitUrl) == 1 {
		return "", errors.New("bad url")
	} else {
		return splitUrl[1], nil
	}
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
