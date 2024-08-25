package service

import "SANS/internal/repository"

type UserServiceInterface interface {
	//Daftar Fungsi
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: repository,
	}
}