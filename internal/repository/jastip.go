package repository

import (
	"log"

	"braw-api/entity"
	"braw-api/model"
	//"braw-api/pkg/utils"
	"gorm.io/gorm"
)

type IJastipRepository interface {
	CreateJastip(jastipReq *entity.Jastip) (*entity.Jastip, error)
	UpdateJastip(jastipReq *model.UpdateJastip, id string) (*entity.Jastip, error)
	DeleteJastip(id string) error
	GetJastipByID(id string) (*entity.Jastip, error)
	GetAllJastip() ([]*entity.Jastip, error)
	GetJastipByUserID(id string) ([]*entity.Jastip, error)
}

type JastipRepository struct {
	db *gorm.DB
}

func NewJastipRepository(db *gorm.DB) IJastipRepository {
	return &JastipRepository{
		db: db,
	}
}

func (br *JastipRepository) CreateJastip(jastipReq *entity.Jastip) (*entity.Jastip, error) {
	if err := br.db.Debug().Create(&jastipReq).Error; err != nil {
		return nil, err
	}
	return jastipReq, nil
}

func (br *JastipRepository) UpdateJastip(jastipReq *model.UpdateJastip, id string) (*entity.Jastip, error) {
	tx := br.db.Begin()

	var jastip *entity.Jastip
	if err := tx.Debug().Where("id_jastip = ?", id).First(&jastip).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	jastip = parseUpdateReqJastip(jastip, jastipReq)

	if err := tx.Debug().Where("id_jastip = ?", id).Save(&jastip).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return jastip, nil
}

func (br *JastipRepository) DeleteJastip(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id_jastip = ?", id).Delete(&entity.Jastip{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *JastipRepository) GetJastipByID(id string) (*entity.Jastip, error) {
	var jastip entity.Jastip
	if err := br.db.Debug().Where("id_jastip = ?", id).Find(&jastip).Error; err != nil {
		return nil, err
	}
	return &jastip, nil
}

func (br *JastipRepository) GetJastipByUserID(id string) ([]*entity.Jastip, error) {
	var jastip []*entity.Jastip
	if err := br.db.Debug().Where("id_user = ?", id).Find(&jastip).Error; err != nil {
		return nil, err
	}
	return jastip, nil
}

func (br *JastipRepository) GetAllJastip() ([]*entity.Jastip, error) {
	var jastips []*entity.Jastip
	if err := br.db.Debug().Find(&jastips).Error; err != nil {
		return nil, err
	}
	return jastips, nil
}

func parseUpdateReqJastip(jastip *entity.Jastip, jastipReq *model.UpdateJastip) *entity.Jastip {
	if jastipReq.Title != "" {
		jastip.Title = jastipReq.Title
	}
	if jastipReq.Category != "" {
		jastip.Category = jastipReq.Category
	}
	if jastipReq.Price > 0 {
		jastip.Price = jastipReq.Price
	}
	if jastipReq.OpenDay != "" {
		jastip.OpenDay = jastipReq.OpenDay
	}
	if jastipReq.CloseOrder != "" {
		jastip.CloseOrder = jastipReq.CloseOrder//utils.ParseTime(jastipReq.CloseOrder)

	}

	return jastip
}
