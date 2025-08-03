package dto

import models "blog/model"

type PostResult struct {
	Page  int
	Size  int
	Total int64
	List  []models.Post
}
