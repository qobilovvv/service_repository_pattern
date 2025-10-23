package service

import (
	"github.com/qobilovvv/patterns/service_repo/models"
	"github.com/qobilovvv/patterns/service_repo/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(ID uint) (*models.User, error)
	CreateUser(user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(ID uint) (*models.User, error) {
	return s.repo.FindByID(ID)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Save(user)
}
