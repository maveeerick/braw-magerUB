package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository                IUserRepository
	PrelovedRepository            IPrelovedRepository
	JastipRepository              IJastipRepository
	JasantarRepository            IJasantarRepository
	KomunitasbrawRepository       IKomunitasbrawRepository
	PrelovedImagesRepository      IPrelovedImagesRepository
	KomunitasbrawImagesRepository IKomunitasbrawImagesRepository
	JasantarImagesRepository      IJasantarImagesRepository
	JastipImagesRepository        IJastipImagesRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	prelovedRepository := NewPrelovedRepository(db)
	jastipRepository := NewJastipRepository(db)
	jasantarRepository := NewJasantarRepository(db)
	komunitasbrawRepository := NewKomunitasbrawRepository(db)
	prelovedImagesRepository := NewPrelovedImagesRepository(db)
	komunitasbrawImagesRepository := NewKomunitasbrawImagesRepository(db)
	jastipImagesRepository := NewJastipImagesRepository(db)
	jasantarImagesRepository := NewJasantarImagesRepository(db)

	return &Repository{
		UserRepository:                userRepository,
		PrelovedRepository:            prelovedRepository,
		JastipRepository:              jastipRepository,
		JasantarRepository:            jasantarRepository,
		KomunitasbrawRepository:       komunitasbrawRepository,
		PrelovedImagesRepository:      prelovedImagesRepository,
		KomunitasbrawImagesRepository: komunitasbrawImagesRepository,
		JastipImagesRepository:        jastipImagesRepository,
		JasantarImagesRepository:      jasantarImagesRepository,
	}
}
