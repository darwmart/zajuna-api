package services

import (
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	return s.repo.FindByFilters(filters, page, limit)
}

func (s *UserService) DeleteUsers(userIDs []int) error {
	return s.repo.DeleteUsers(userIDs)
}

func (s *UserService) UpdateUsers(users []models.User) (int64, error) {
	return s.repo.UpdateUsers(users)
}
