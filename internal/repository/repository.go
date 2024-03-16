package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	PrelovedRepository IPrelovedRepository
	JastipRepository IJastipRepository
	JasantarRepository IJasantarRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	prelovedRepository := NewPrelovedRepository(db)
	jastipRepository := NewJastipRepository(db)
	jasantarRepository := NewJasantarRepository(db)

	return &Repository{
		UserRepository: userRepository,
		PrelovedRepository: prelovedRepository,
		JastipRepository: jastipRepository,
		JasantarRepository: jasantarRepository,
	
	}
}