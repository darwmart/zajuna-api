package services

import (
	"fmt"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	return s.repo.FindByFilters(filters, page, limit)
}

func (s *UserService) Login(username, password string) (string, error) {

	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("usuario no encontrado")
	}
	return "aaa", nil
}
