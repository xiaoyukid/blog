package web

import (
	"blog/common/constModel"
	models "blog/model"
	"blog/model/bo"
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetPosts(context *gin.Context) {
	var query bo.PostQuery
	if err := context.ShouldBindJSON(&query); err != nil {
		context.Error(constModel.ErrAuthInvalid.WithDetails("请输入分页参数： %v" + err.Error()))
		//context.JSON(http.StatusBadRequest, gin.H{"error": "请输入分页参数！"})
		return
	}
	//先查询userid
	userID := context.GetFloat64("userId")
	//如果userid为0，则返回错误
	if userID == 0 {
		context.Error(constModel.ErrAuthInvalid.WithDetails("请登录后操作！"))
		return
	}
	query.UserID = int(userID)
	posts := h.postService.GetPosts(query)
	context.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPostById(c *gin.Context) {
	id := c.Query("postID")
	//string转换成int类型
	num, _ := strconv.Atoi(id)
	post := h.postService.GetPostById(num)
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//获取用户id
	userID := c.GetFloat64("userId")
	post.UserID = uint(int(userID))
	createdPost, err := h.postService.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdPost)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//这需要判断标题不为空
	if post.Title == "" || post.Content == "" {
		c.Error(constModel.ErrAuthInvalid.WithDetails("请输入文章标题或内容！"))
		return
	}
	//获取用户id
	userID := c.GetFloat64("userId")
	if userID < 1 {
		c.Error(constModel.ErrAuthInvalid.WithDetails("请登录！"))
		return
	}
	//获取文章id
	if post.ID < 1 {
		c.Error(constModel.ErrAuthInvalid.WithDetails("请输入文章id！"))
		return
	}
	post.UserID = uint(userID)
	if err := h.postService.UpdatePost(&post); err != nil {
		c.Error(constModel.ErrInternalServer.WithDetails("更新失败！"))
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "更新成功！"})
	}
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	type deleteIds struct {
		Ids []int `json:"ids" binding:"required"`
	}
	var postIds deleteIds
	if err := c.ShouldBindJSON(&postIds); err != nil {
		c.Error(constModel.ErrAuthInvalid.WithDetails("请输入文章id！"))
		return
	}
	//获取用户id，只有自己的文章才能删除
	userId := c.GetFloat64("userId")
	if userId < 1 {
		c.Error(constModel.ErrAuthInvalid.WithDetails("请登录！"))
		return
	}
	if err := h.postService.DeletePost(postIds.Ids); err != nil {
		c.Error(constModel.ErrInternalServer.WithDetails(err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "删除成功！"})
	}
}
