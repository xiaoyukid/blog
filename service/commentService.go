package service

import (
	models "blog/model"
	"blog/model/bo"
	"blog/model/dto"
	"blog/repositories"
)

type CommentService interface {
	CreateComment(comment *models.Comment) (models.Comment, error)
	GetComments(query bo.CommentQuery) dto.CommentResult
}

type commentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(commentRepository repositories.CommentRepository) CommentService {
	return &commentService{commentRepository: commentRepository}
}

func (s *commentService) CreateComment(comment *models.Comment) (models.Comment, error) {
	return s.commentRepository.CreateComment(comment)
}

func (s *commentService) GetComments(query bo.CommentQuery) dto.CommentResult {
	return s.commentRepository.GetComments(query)
}
