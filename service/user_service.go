package service

import "github.com/pdnguyen1503/base-go/repository"

type userService struct {
	UserRepo repository.UserRepository
}

type UserConfig struct {
	UserRepo repository.UserRepository
}

func NewUserService(c *UserConfig) UserService {
	return &userService{
		UserRepo: c.UserRepo,
	}
}

func (s *userService) GetUser() (string, error) {
	result, err := s.UserRepo.GetUser()
	if err != nil {
		return "", err
	}
	return result, nil
}
