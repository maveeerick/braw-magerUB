package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	PrelovedRepository IPrelovedRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	prelovedRepository := NewPrelovedRepository(db)

	return &Repository{
		UserRepository: userRepository,
		PrelovedRepository: prelovedRepository,
	
	}
}