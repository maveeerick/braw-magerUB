package service

import (
	"braw-api/internal/repository"
	"braw-api/pkg/bcrypt"
	"braw-api/pkg/jwt"
	"braw-api/pkg/supabase"
)

type Service struct {
	UserService IUserService
	//BookService IBookService
	//RentService IRentService
}

type InitParam struct {
	Repository *repository.Repository
	JwtAuth    jwt.Interface
	Bcrypt     bcrypt.Interface
	Supabase   supabase.Interface
}

func NewService(param InitParam) *Service {
	userService := NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth, param.Supabase)
	//bookService := NewBookService(param.Repository.BookRepository)
	//rentService := NewRentService(param.Repository.RentRepository, param.Repository.UserRepository, param.Repository.BookRepository)

	return &Service{
		UserService: userService,
		//BookService: bookService,
		//RentService: rentService,
	}
}