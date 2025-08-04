package dto

import models "blog/model"

type CommentResult struct {
	Page  int
	Size  int
	Total int64
	List  []models.Comment
}
