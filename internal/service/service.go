package service

import (
	"braw-api/internal/repository"
	"braw-api/pkg/bcrypt"
	"braw-api/pkg/jwt"
	"braw-api/pkg/supabase"
)

type Service struct {
	UserService IUserService
	PrelovedService IPrelovedService
	JastipService IJastipService
	JasantarService IJasantarService
	KomunitasbrawService IKomunitasbrawService
	PhotolinkService IPhotolinkService
}

type InitParam struct {
	Repository *repository.Repository
	JwtAuth    jwt.Interface
	Bcrypt     bcrypt.Interface
	Supabase   supabase.Interface
}

func NewService(param InitParam) *Service {
	userService := NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth, param.Supabase)
	prelovedService := NewPrelovedService(param.Repository.PrelovedRepository)
	jastipService := NewJastipService(param.Repository.JastipRepository)
	jasantarService := NewJasantarService(param.Repository.JasantarRepository)
	komunitasbrawService := NewKomunitasbrawService(param.Repository.KomunitasbrawRepository)
	PhotolinkService := NewPhotolinkService(param.Repository.PhotolinkRepository)

	return &Service{
		UserService: userService,
		PrelovedService: prelovedService,
		JastipService: jastipService,
		JasantarService: jasantarService,
		KomunitasbrawService: komunitasbrawService,
		PhotolinkService: PhotolinkService,
	}
}