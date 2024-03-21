package repository

import (
	"braw-api/entity"
	"braw-api/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	//GetUserWithRent(param model.UserParam) (entity.User, error)
	UpdateUser(user entity.User, param model.UserParam) error
	UpdateUserData(UserReq *model.UpdateUserData, id string) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := u.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}


func (u *UserRepository) UpdateUser(user entity.User, param model.UserParam) error {
	err := u.db.Debug().Model(&entity.User{}).Where(param).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (br *UserRepository) UpdateUserData(userDataReq *model.UpdateUserData, id string) (*entity.User, error) {
	tx := br.db.Begin()

	var userData *entity.User
	if err := tx.Debug().Where("id_user = ?", id).First(&userData).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	userData = parseUpdateReqUserData(userData, userDataReq)

	if err := tx.Debug().Where("id_user = ?", id).Save(&userData).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return userData, nil
}

func parseUpdateReqUserData(userData *entity.User, userDataReq *model.UpdateUserData) *entity.User {
	if userDataReq.Username != "" {
		userData.Username = userDataReq.Username
	}
	if userDataReq.Email != "" {
		userData.Email = userDataReq.Email
	}
	if userDataReq.Password != "" {
		userData.Password = userDataReq.Password
	}
	if userDataReq.Contact != "" {
		userData.Contact = userDataReq.Contact
	}
	if userDataReq.PhotoLink != "" {
		userData.PhotoLink = userDataReq.PhotoLink
	}
	if userDataReq.Alamat != "" {
		userData.Alamat = userDataReq.Alamat
	}
	return userData
}