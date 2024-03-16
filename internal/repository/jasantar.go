package repository

import (
	"log"

	"braw-api/entity"
	"braw-api/model"
	"gorm.io/gorm"
)

type IJasantarRepository interface {
	CreateJasantar(jasantarReq *entity.Jasantar) (*entity.Jasantar, error)
	UpdateJasantar(jasantarReq *model.UpdateJasantar, id string) (*entity.Jasantar, error)
	DeleteJasantar(id string) error
	GetJasantarByID(id string) (*entity.Jasantar, error)
	GetAllJasantar() ([]*entity.Jasantar, error)
}

type JasantarRepository struct {
	db *gorm.DB
}

func NewJasantarRepository(db *gorm.DB) IJasantarRepository {
	return &JasantarRepository{
		db: db,
	}
}

func (br *JasantarRepository) CreateJasantar(jasantarReq *entity.Jasantar) (*entity.Jasantar, error) {
	if err := br.db.Debug().Create(&jasantarReq).Error; err != nil {
		return nil, err
	}
	return jasantarReq, nil
}

func (br *JasantarRepository) UpdateJasantar(jasantarReq *model.UpdateJasantar, id string) (*entity.Jasantar, error) {
	tx := br.db.Begin()

	var jasantar *entity.Jasantar
	if err := tx.Debug().Where("id = ?", id).First(&jasantar).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	jasantar = parseUpdateReq(jasantar, jasantarReq)

	if err := tx.Debug().Where("id = ?", id).Save(&jasantar).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return jasantar, nil
}

func (br *JasantarRepository) DeleteJasantar(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.Jasantar{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *JasantarRepository) GetJasantarByID(id string) (*entity.Jasantar, error) {
	var jasantar entity.Jasantar
	if err := br.db.Debug().Where("id = ?", id).Find(&jasantar).Error; err != nil {
		return nil, err
	}
	return &jasantar, nil
}

func (br *JasantarRepository) GetAllJasantar() ([]*entity.Jasantar, error) {
	var jasantars []*entity.Jasantar
	if err := br.db.Debug().Find(&jasantars).Error; err != nil {
		return nil, err
	}
	return jasantars, nil
}


func parseUpdateReq(jasantar *entity.Jasantar, jasantarReq *model.UpdateJasantar) *entity.Jasantar {
	if jasantarReq.Title != "" {
		jasantar.Title = jasantarReq.Title
	}
	if jasantarReq.Area != "" {
		jasantar.Area = jasantarReq.Area
	}
	if jasantarReq.Category != "" {
		jasantar.Category = jasantarReq.Category
	}
	
	return jasantar
}