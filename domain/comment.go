package domain

type Comment struct {
	UserID  string
	Content string
}

type CommentDto struct {
	ID      string
	Comment Comment
}

type CommentUrlParameter struct {
	ID string `uri:"commentId"`
}
