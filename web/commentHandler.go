package web

import (
	"blog/common/constModel"
	models "blog/model"
	"blog/model/bo"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (c *CommentHandler) GetComments(context *gin.Context) {
	var query bo.CommentQuery
	if err := context.ShouldBind(&query); err != nil {
		context.Error(constModel.ErrAuthInvalid.WithDetails("请输入分页参数： %v" + err.Error()))
		return
	}
	userID := context.GetFloat64("userId")
	if userID == 0 {
		context.Error(constModel.ErrAuthInvalid.WithDetails("请登录后操作！"))
		return
	}
	commentResult := c.commentService.GetComments(query)
	context.JSON(http.StatusOK, commentResult)
}

func (c *CommentHandler) CreateComment(context *gin.Context) {
	var comment models.Comment

	userID := context.GetUint("userId")
	comment.UserID = userID

	createComment, err := c.commentService.CreateComment(&comment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, createComment)
}
