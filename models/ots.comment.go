package models

type Comment struct {
	ProdId    string // prodid as partition key
	CommentId int64  //row key for each comment
	Content   string //comments text
}
