package model

type PostResponse struct {
	Count int
	Items []Post
}

type Post struct {
	Link string
	Id   int
	Date int
	Text string
}
