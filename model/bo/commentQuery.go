package bo

type CommentQuery struct {
	PostId         uint   `query:"postId" binding:"required"`
	Page           int    `query:"page" binding:"required"`
	Size           int    `query:"size" binding:"required"`
	UserId         int    `query:"userId" binding:"required"`
	CommentContent string `json:"commentContent"`
}
