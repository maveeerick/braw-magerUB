package repository

import (
	"log"

	"braw-api/entity"
	"braw-api/model"
	"gorm.io/gorm"
)

type IKomunitasbrawRepository interface {
	CreateKomunitasbraw(komunitasbrawReq *entity.Komunitasbraw) (*entity.Komunitasbraw, error)
	UpdateKomunitasbraw(komunitasbrawReq *model.UpdateKomunitasbraw, id string) (*entity.Komunitasbraw, error)
	DeleteKomunitasbraw(id string) error
	GetKomunitasbrawByID(id string) (*entity.Komunitasbraw, error)
	GetAllKomunitasbraw() ([]*entity.Komunitasbraw, error)
	GetKomunitasbrawByUserID(id string) ([]*entity.Komunitasbraw, error)
}

type KomunitasbrawRepository struct {
	db *gorm.DB
}

func NewKomunitasbrawRepository(db *gorm.DB) IKomunitasbrawRepository {
	return &KomunitasbrawRepository{
		db: db,
	}
}

func (br *KomunitasbrawRepository) CreateKomunitasbraw(komunitasbrawReq *entity.Komunitasbraw) (*entity.Komunitasbraw, error) {
	if err := br.db.Debug().Create(&komunitasbrawReq).Error; err != nil {
		return nil, err
	}
	return komunitasbrawReq, nil
}

func (br *KomunitasbrawRepository) UpdateKomunitasbraw(komunitasbrawReq *model.UpdateKomunitasbraw, id string) (*entity.Komunitasbraw, error) {
	tx := br.db.Begin()

	var komunitasbraw *entity.Komunitasbraw
	if err := tx.Debug().Where("id_komunitasbraw = ?", id).First(&komunitasbraw).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	komunitasbraw = parseUpdateReqKomunitasbraw(komunitasbraw, komunitasbrawReq)

	if err := tx.Debug().Where("id_komunitasbraw = ?", id).Save(&komunitasbraw).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return komunitasbraw, nil
}

func (br *KomunitasbrawRepository) DeleteKomunitasbraw(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id_komunitasbraw = ?", id).Delete(&entity.Komunitasbraw{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *KomunitasbrawRepository) GetKomunitasbrawByID(id string) (*entity.Komunitasbraw, error) {
	var komunitasbraw entity.Komunitasbraw
	if err := br.db.Debug().Where("id_komunitasbraw = ?", id).Find(&komunitasbraw).Error; err != nil {
		return nil, err
	}
	return &komunitasbraw, nil
}

func (br *KomunitasbrawRepository) GetKomunitasbrawByUserID(id string) ([]*entity.Komunitasbraw, error) {
	var komunitasbraw []*entity.Komunitasbraw
	if err := br.db.Debug().Where("id_user = ?", id).Find(&komunitasbraw).Error; err != nil {
		return nil, err
	}
	return komunitasbraw, nil
}

func (br *KomunitasbrawRepository) GetAllKomunitasbraw() ([]*entity.Komunitasbraw, error) {
	var komunitasbraws []*entity.Komunitasbraw
	if err := br.db.Debug().Find(&komunitasbraws).Error; err != nil {
		return nil, err
	}
	return komunitasbraws, nil
}

func parseUpdateReqKomunitasbraw(komunitasbraw *entity.Komunitasbraw, komunitasbrawReq *model.UpdateKomunitasbraw) *entity.Komunitasbraw {
	if komunitasbrawReq.Title != "" {
		komunitasbraw.Title = komunitasbrawReq.Title
	}
	if komunitasbrawReq.Category != "" {
		komunitasbraw.Category = komunitasbrawReq.Category
	}
	if komunitasbrawReq.Description != "" {
		komunitasbraw.Description = komunitasbrawReq.Description
	}

	return komunitasbraw
}