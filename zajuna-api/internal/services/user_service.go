package services

import (
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(filters map[string]string, page, pageSize int) ([]models.User, int, error) {
	return s.repo.GetUsers(filters, page, pageSize)
}
