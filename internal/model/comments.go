package model

type CommentsResponse struct {
	Count int
	Items []Comment
}

type Comment struct {
	Link string
	Id   int
	Date int
	Text string
}
