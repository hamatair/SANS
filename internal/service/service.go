package service

import "SANS/internal/repository"

type Service struct {
	UserService UserServiceInterface
}

type InitParam struct {
	Repository *repository.Repository
}

func NewService(param InitParam) *Service {
	return &Service{
		UserService: NewUserService(param.Repository.UserRepository),
	}
}
