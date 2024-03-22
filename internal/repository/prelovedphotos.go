package repository

import (
	"log"

	"braw-api/entity"

	"gorm.io/gorm"
)

type IPrelovedImagesRepository interface {
	CreatePrelovedImage(prelovedReq *entity.PrelovedPhotos) (*entity.PrelovedPhotos, error)
	DeletePrelovedImage(id string) error
	GetImageByID(id string) (*entity.PrelovedPhotos, error)
	GetImagesByPrelovedID(id string) ([]*entity.PrelovedPhotos, error)
}

type PrelovedImagesRepository struct {
	db *gorm.DB
}

func NewPrelovedImagesRepository(db *gorm.DB) IPrelovedImagesRepository {
	return &PrelovedImagesRepository{
		db: db,
	}
}

func (br *PrelovedImagesRepository) CreatePrelovedImage(prelovedReq *entity.PrelovedPhotos) (*entity.PrelovedPhotos, error) {
	if err := br.db.Debug().Create(&prelovedReq).Error; err != nil {
		return nil, err
	}
	return prelovedReq, nil
}

func (br *PrelovedImagesRepository) DeletePrelovedImage(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.PrelovedPhotos{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *PrelovedImagesRepository) GetImageByID(id string) (*entity.PrelovedPhotos, error) {
	var image entity.PrelovedPhotos
	if err := br.db.Debug().Where("id = ?", id).First(&image).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (br *PrelovedImagesRepository) GetImagesByPrelovedID(id string) ([]*entity.PrelovedPhotos, error) {
	var images []*entity.PrelovedPhotos
	if err := br.db.Debug().Where("preloved_id = ?", id).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}