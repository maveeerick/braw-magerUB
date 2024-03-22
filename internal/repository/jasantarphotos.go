package repository

import (
	"log"

	"braw-api/entity"

	"gorm.io/gorm"
)

type IJasantarImagesRepository interface {
	CreateJasantarImage(jasantarReq *entity.JasantarPhotos) (*entity.JasantarPhotos, error)
	DeleteJasantarImage(id string) error
	GetImageByID(id string) (*entity.JasantarPhotos, error)
	GetImagesByJasantarID(id string) ([]*entity.JasantarPhotos, error)
}

type JasantarImagesRepository struct {
	db *gorm.DB
}

func NewJasantarImagesRepository(db *gorm.DB) IJasantarImagesRepository {
	return &JasantarImagesRepository{
		db: db,
	}
}

func (br *JasantarImagesRepository) CreateJasantarImage(jasantarReq *entity.JasantarPhotos) (*entity.JasantarPhotos, error) {
	if err := br.db.Debug().Create(&jasantarReq).Error; err != nil {
		return nil, err
	}
	return jasantarReq, nil
}

func (br *JasantarImagesRepository) DeleteJasantarImage(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.JasantarPhotos{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *JasantarImagesRepository) GetImageByID(id string) (*entity.JasantarPhotos, error) {
	var image entity.JasantarPhotos
	if err := br.db.Debug().Where("id = ?", id).First(&image).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (br *JasantarImagesRepository) GetImagesByJasantarID(id string) ([]*entity.JasantarPhotos, error) {
	var images []*entity.JasantarPhotos
	if err := br.db.Debug().Where("jasantar_id = ?", id).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}