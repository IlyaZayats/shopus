package services

import (
	"github.com/IlyaZayats/shopus/internal/entity"
	"github.com/IlyaZayats/shopus/internal/interfaces"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) (*UserService, error) {
	return &UserService{
		repo: repo,
	}, nil
}

func (s *UserService) Login(login, pwd string) (int, error) {
	return s.repo.Login(entity.User{Login: login, Pwd: pwd})
}
