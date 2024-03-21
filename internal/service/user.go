package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/model"
	"braw-api/pkg/bcrypt"
	"braw-api/pkg/jwt"
	"braw-api/pkg/supabase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserRegister) error
	GetUser(param model.UserParam) (entity.User, error)
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	UploadPhoto(ctx *gin.Context, param model.UserUploadPhoto) error
	UpdateUserData(userDataReq *model.UpdateUserData, id string) (*entity.User, error)
}

type UserService struct {
	ur       repository.IUserRepository
	bcrypt   bcrypt.Interface
	jwtAuth  jwt.Interface
	supabase supabase.Interface
	br repository.IUserRepository
}

//type User

func NewUserService(br repository.IUserRepository, userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) IUserService {
	return &UserService{
		ur:       userRepository,
		bcrypt:   bcrypt,
		jwtAuth:  jwtAuth,
		supabase: supabase,
		br: br,
	}
}

func (u *UserService) Register(param model.UserRegister) error {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	param.IdUser = uuid.New()
	param.Password = hashPassword

	user := entity.User{
		IdUser:       param.IdUser,
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
		Username: param.Username,
		Role:     2,
	}

	_, err = u.ur.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}

	user, err := u.ur.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = u.bcrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwtAuth.CreateJWTToken(user.IdUser)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.ur.GetUser(param)
}


func (u *UserService) UploadPhoto(ctx *gin.Context, param model.UserUploadPhoto) error {
	user, err := u.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		return err
	}

	if user.PhotoLink != "" {
		err := u.supabase.Delete(user.PhotoLink)
		if err != nil {
			return err
		}
	}

	link, err := u.supabase.Upload(param.Photo)
	if err != nil {
		return err
	}

	err = u.ur.UpdateUser(entity.User{
		PhotoLink: link,
	}, model.UserParam{
		IdUser: user.IdUser,
	})
	if err != nil {
		return err
	}

	return nil
}

func (bs *UserService) UpdateUserData(userDataReq *model.UpdateUserData, id string) (*entity.User, error) {
	userData, err := bs.br.UpdateUserData(userDataReq, id)
	if err != nil {
		return nil, err
	}

	return userData, nil
}