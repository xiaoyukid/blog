package repositories

import (
	models "blog/model"
	"blog/model/bo"
	"blog/model/dto"
	"fmt"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetPosts(query bo.PostQuery) dto.PostResult
	CreatePost(post *models.Post) (models.Post, error)
	GetPostById(id int) models.Post
	UpdatePost(post *models.Post) error
	DeletePost(ids []int) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) GetPosts(query bo.PostQuery) dto.PostResult {
	db := p.db
	var posts []models.Post
	if query.Title != "" {
		db = db.Where("title like ?", "%"+query.Title+"%")
	}
	//判断内容是否有内容，有的话拼接上参数查询
	if query.Content != "" {
		db = db.Where("content like ?", "%"+query.Content+"%")
	}
	//这里需要返回总条数，方便分页
	var total int64
	//countDB := db.Session(&gorm.Session{})
	if err := db.Debug().Model(&models.Post{}).Where("user_id = ?", query.UserID).Count(&total).Error; err != nil {
		fmt.Println("count error:", err)
	}
	db.Debug().Where("user_id = ?", query.UserID).Limit(query.Size).Offset((query.Page - 1) * query.Size).Find(&posts)
	//context.JSON(http.StatusOK, gin.H{"data": posts})
	result := dto.PostResult{}
	result.Total = total
	result.List = posts
	result.Page = query.Page
	result.Size = query.Size
	return result
}

func (p *postRepository) CreatePost(post *models.Post) (models.Post, error) {
	return *post, p.db.Create(post).Error
}

func (p *postRepository) GetPostById(id int) models.Post {
	var post models.Post
	p.db.First(&post, id)
	return post
}

func (p *postRepository) UpdatePost(post *models.Post) error {
	return p.db.Updates(post).Error
}

func (p *postRepository) DeletePost(ids []int) error {
	return p.db.Delete(&models.Post{}, ids).Error
}
