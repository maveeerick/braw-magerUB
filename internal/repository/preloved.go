package repository

import (
	"log"

	"braw-api/entity"
	"braw-api/model"
	"gorm.io/gorm"
)

type IPrelovedRepository interface {
	CreatePreloved(prelovedReq *entity.Preloved) (*entity.Preloved, error)
	UpdatePreloved(prelovedReq *model.UpdatePreloved, id string) (*entity.Preloved, error)
	DeletePreloved(id string) error
	GetPrelovedByID(id string) (*entity.Preloved, error)
	GetAllPreloved() ([]*entity.Preloved, error)
	GetPrelovedByUserID(id string) ([]*entity.Preloved, error)
}

type PrelovedRepository struct {
	db *gorm.DB
}

func NewPrelovedRepository(db *gorm.DB) IPrelovedRepository {
	return &PrelovedRepository{
		db: db,
	}
}

func (br *PrelovedRepository) CreatePreloved(prelovedReq *entity.Preloved) (*entity.Preloved, error) {
	if err := br.db.Debug().Create(&prelovedReq).Error; err != nil {
		return nil, err
	}
	return prelovedReq, nil
}

func (br *PrelovedRepository) UpdatePreloved(prelovedReq *model.UpdatePreloved, id string) (*entity.Preloved, error) {
	tx := br.db.Begin()

	var preloved *entity.Preloved
	if err := tx.Debug().Where("id_preloved = ?", id).First(&preloved).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	preloved = parseUpdateReqPreloved(preloved, prelovedReq)

	if err := tx.Debug().Where("id_preloved = ?", id).Save(&preloved).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return preloved, nil
}

func (br *PrelovedRepository) DeletePreloved(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id_preloved = ?", id).Delete(&entity.Preloved{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *PrelovedRepository) GetPrelovedByID(id string) (*entity.Preloved, error) {
	var preloved entity.Preloved
	if err := br.db.Debug().Where("id_preloved = ?", id).First(&preloved).Error; err != nil {
		return nil, err
	}
	return &preloved, nil
}

func (br *PrelovedRepository) GetPrelovedByUserID(id string) ([]*entity.Preloved, error) {
	var preloved []*entity.Preloved
	if err := br.db.Debug().Where("id_user = ?", id).Find(&preloved).Error; err != nil {
		return nil, err
	}
	return preloved, nil
}

func (br *PrelovedRepository) GetAllPreloved() ([]*entity.Preloved, error) {
	var preloveds []*entity.Preloved
	if err := br.db.Debug().Find(&preloveds).Error; err != nil {
		return nil, err
	}
	return preloveds, nil
}

func parseUpdateReqPreloved(preloved *entity.Preloved, prelovedReq *model.UpdatePreloved) *entity.Preloved {
	if prelovedReq.Title != "" {
		preloved.Title = prelovedReq.Title
	}
	if prelovedReq.Category != "" {
		preloved.Category = prelovedReq.Category
	}
	if prelovedReq.Price > 0 {
		preloved.Price = prelovedReq.Price
	}
	if prelovedReq.Condition != "" {
		preloved.Condition = prelovedReq.Condition
	}
	if prelovedReq.Description != "" {
		preloved.Description = prelovedReq.Description
	}
	return preloved
}

