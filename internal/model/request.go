package model

type Request struct {
	Text  string   `json:"text" binding:"required"`
	Links []string `json:"links" binding:"required"`
}
