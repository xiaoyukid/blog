package bo

type CommentQuery struct {
	PostId         uint   `query:"postId" binding:"required"`
	Page           int    `query:"page" binding:"required"`
	Size           int    `query:"size" binding:"required"`
	UserId         uint   `query:"userId"`
	CommentContent string `json:"commentContent"`
}
