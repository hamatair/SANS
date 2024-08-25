package repository

import "gorm.io/gorm"

type UserRepositoryInterface interface {
	//Daftar Fungsi
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}
