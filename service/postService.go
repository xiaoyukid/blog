package service

import (
	models "blog/model"
	"blog/model/bo"
	"blog/model/dto"
	"blog/repositories"
)

type PostService interface {
	CreatePost(post *models.Post) (models.Post, error)
	GetPosts(query bo.PostQuery) dto.PostResult
	GetPostById(id int) models.Post
	UpdatePost(post *models.Post) error
	DeletePost(ids []int) error
}

type postService struct {
	postRepository repositories.PostRepository
}

func NewPostService(postRepository repositories.PostRepository) PostService {
	return &postService{postRepository: postRepository}
}

func (p *postService) CreatePost(post *models.Post) (models.Post, error) {
	return p.postRepository.CreatePost(post)
}

func (p postService) UpdatePost(post *models.Post) error {
	return p.postRepository.UpdatePost(post)
}
func (p postService) DeletePost(ids []int) error {
	return p.postRepository.DeletePost(ids)
}

func (p postService) GetPosts(query bo.PostQuery) dto.PostResult {
	return p.postRepository.GetPosts(query)
}
func (p postService) GetPostById(id int) models.Post {
	return p.postRepository.GetPostById(id)
}
