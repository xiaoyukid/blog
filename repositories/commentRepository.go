package repositories

import (
	models "blog/model"
	"blog/model/bo"
	"blog/model/dto"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) (models.Comment, error)
	GetComments(query bo.CommentQuery) dto.CommentResult
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (c *commentRepository) CreateComment(comment *models.Comment) (models.Comment, error) {
	return *comment, c.db.Create(&comment).Error
}

func (c *commentRepository) GetComments(query bo.CommentQuery) dto.CommentResult {
	var comments []models.Comment
	db := *c.db
	if query.PostId > 0 {
		db.Where("post_id = ?", query.PostId)
	}
	if query.CommentContent != "" {
		db.Where("comment_content = ?", query.CommentContent)
	}
	db.Offset((query.Page - 1) * query.Size).Limit(query.Size).Find(&comments)
	var result dto.CommentResult
	result.Page = query.Page
	result.Size = query.Size
	result.List = comments
	db.Count(&result.Total)
	return result
}
