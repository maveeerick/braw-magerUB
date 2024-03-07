package main

import (
	"braw-api/internal/handler/rest"
	"braw-api/internal/repository"
	"braw-api/internal/service"
	"braw-api/pkg/bcrypt"
	"braw-api/pkg/config"
	"braw-api/pkg/database/mysql"
	"braw-api/pkg/jwt"
	"braw-api/pkg/middleware"
	"braw-api/pkg/supabase"
)

func main() {
	config.LoadEnv()

	jwtAuth := jwt.Init()

	bcrypt := bcrypt.Init()

	supabase := supabase.Init()

	db := mysql.ConnectDatabase()

	repository := repository.NewRepository(db)

	service := service.NewService(service.InitParam{Repository: repository, JwtAuth: jwtAuth, Bcrypt: bcrypt, Supabase: supabase})

	middleware := middleware.Init(jwtAuth, service)

	rest := rest.NewRest(service, middleware)

	mysql.Migration(db)

	mysql.SeedData(db)

	rest.MountEndpoint()

	rest.Serve()
}