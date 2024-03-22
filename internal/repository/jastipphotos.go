package repository

import (
	"log"

	"braw-api/entity"

	"gorm.io/gorm"
)

type IJastipImagesRepository interface {
	CreateJastipImage(jastipReq *entity.JastipPhotos) (*entity.JastipPhotos, error)
	DeleteJastipImage(id string) error
	GetImageByID(id string) (*entity.JastipPhotos, error)
	GetImagesByJastipID(id string) ([]*entity.JastipPhotos, error)
}

type JastipImagesRepository struct {
	db *gorm.DB
}

func NewJastipImagesRepository(db *gorm.DB) IJastipImagesRepository {
	return &JastipImagesRepository{
		db: db,
	}
}

func (br *JastipImagesRepository) CreateJastipImage(jastipReq *entity.JastipPhotos) (*entity.JastipPhotos, error) {
	if err := br.db.Debug().Create(&jastipReq).Error; err != nil {
		return nil, err
	}
	return jastipReq, nil
}

func (br *JastipImagesRepository) DeleteJastipImage(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.JastipPhotos{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *JastipImagesRepository) GetImageByID(id string) (*entity.JastipPhotos, error) {
	var image entity.JastipPhotos
	if err := br.db.Debug().Where("id = ?", id).First(&image).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (br *JastipImagesRepository) GetImagesByJastipID(id string) ([]*entity.JastipPhotos, error) {
	var images []*entity.JastipPhotos
	if err := br.db.Debug().Where("jastip_id = ?", id).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}