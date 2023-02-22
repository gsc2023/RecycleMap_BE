package domain

type Comment struct {
	UID        string
	Content    string
	LocationID string
}

type CommentDto struct {
	ID      string
	Comment Comment
}

type SaveCommentDto struct {
	Content string
}

type CommentUrlParameter struct {
	ID string `uri:"commentId"`
}
