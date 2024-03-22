package repository

import (
	"log"

	"braw-api/entity"

	"gorm.io/gorm"
)

type IKomunitasbrawImagesRepository interface {
	CreateKomunitasbrawImage(komunitasbrawReq *entity.KomunitasbrawPhotos) (*entity.KomunitasbrawPhotos, error)
	DeleteKomunitasbrawImage(id string) error
	GetImageByID(id string) (*entity.KomunitasbrawPhotos, error)
	GetImagesByKomunitasbrawID(id string) ([]*entity.KomunitasbrawPhotos, error)
}

type KomunitasbrawImagesRepository struct {
	db *gorm.DB
}

func NewKomunitasbrawImagesRepository(db *gorm.DB) IKomunitasbrawImagesRepository {
	return &KomunitasbrawImagesRepository{
		db: db,
	}
}

func (br *KomunitasbrawImagesRepository) CreateKomunitasbrawImage(komunitasbrawReq *entity.KomunitasbrawPhotos) (*entity.KomunitasbrawPhotos, error) {
	if err := br.db.Debug().Create(&komunitasbrawReq).Error; err != nil {
		return nil, err
	}
	return komunitasbrawReq, nil
}

func (br *KomunitasbrawImagesRepository) DeleteKomunitasbrawImage(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.KomunitasbrawPhotos{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *KomunitasbrawImagesRepository) GetImageByID(id string) (*entity.KomunitasbrawPhotos, error) {
	var image entity.KomunitasbrawPhotos
	if err := br.db.Debug().Where("id = ?", id).First(&image).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (br *KomunitasbrawImagesRepository) GetImagesByKomunitasbrawID(id string) ([]*entity.KomunitasbrawPhotos, error) {
	var images []*entity.KomunitasbrawPhotos
	if err := br.db.Debug().Where("komunitasbraw_id = ?", id).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}