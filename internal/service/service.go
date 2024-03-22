package service

import (
	"braw-api/internal/repository"
	"braw-api/pkg/bcrypt"
	"braw-api/pkg/jwt"
	"braw-api/pkg/supabase"
)

type Service struct {
	UserService                IUserService
	PrelovedService            IPrelovedService
	JastipService              IJastipService
	JasantarService            IJasantarService
	KomunitasbrawService       IKomunitasbrawService
	PrelovedImagesService      IPrelovedImagesService
	KomunitasbrawImagesService IKomunitasbrawImagesService
	JastipImagesService        IJastipImagesService
	JasantarImagesService      IJasantarImagesService
}

type InitParam struct {
	Repository *repository.Repository
	JwtAuth    jwt.Interface
	Bcrypt     bcrypt.Interface
	Supabase   supabase.Interface
}

func NewService(param InitParam) *Service {
	userService := NewUserService(param.Repository.UserRepository, param.Repository.UserRepository, param.Bcrypt, param.JwtAuth, param.Supabase)
	prelovedService := NewPrelovedService(param.Repository.PrelovedRepository)
	jastipService := NewJastipService(param.Repository.JastipRepository)
	jasantarService := NewJasantarService(param.Repository.JasantarRepository)
	komunitasbrawService := NewKomunitasbrawService(param.Repository.KomunitasbrawRepository)
	prelovedImagesService := NewPrelovedImagesService(param.Repository.PrelovedImagesRepository, param.Repository.PrelovedRepository, param.Supabase)
	jastipImagesService := NewJastipImagesService(param.Repository.JastipImagesRepository, param.Repository.JastipRepository, param.Supabase)
	komunitasbrawImagesService := NewKomunitasbrawImagesService(param.Repository.KomunitasbrawImagesRepository, param.Repository.KomunitasbrawRepository, param.Supabase)
	jasantarImagesService := NewJasantarImagesService(param.Repository.JasantarImagesRepository, param.Repository.JasantarRepository, param.Supabase)

	return &Service{
		UserService:                userService,
		PrelovedService:            prelovedService,
		JastipService:              jastipService,
		JasantarService:            jasantarService,
		KomunitasbrawService:       komunitasbrawService,
		PrelovedImagesService:      prelovedImagesService,
		KomunitasbrawImagesService: komunitasbrawImagesService,
		JastipImagesService:        jastipImagesService,
		JasantarImagesService:      jasantarImagesService,
	}
}
