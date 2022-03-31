package model

type PostResponse struct {
	Count int
	Items []Post
}

type Post struct {
	Id   int
	Date int
	Text string
}
