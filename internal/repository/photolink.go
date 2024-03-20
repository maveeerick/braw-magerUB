package repository

import (
	"log"

	"braw-api/entity"
	"braw-api/model"

	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPhotolinkRepository interface {
	CreatePhotolink(photolinkReq *entity.Photolink) (*entity.Photolink, error)
	UpdatePhotolink(photolinkReq *model.UpdatePhotolink, id string) (*entity.Photolink, error)
	DeletePhotolink(id string) error
	GetPhotolinkByID(id string) (*entity.Photolink, error)
	GetAllPhotolink() ([]*entity.Photolink, error)
}

type PhotolinkRepository struct {
	db *gorm.DB
}

func NewPhotolinkRepository(db *gorm.DB) IPhotolinkRepository {
	return &PhotolinkRepository{
		db: db,
	}
}

func (br *PhotolinkRepository) CreatePhotolink(photolinkReq *entity.Photolink) (*entity.Photolink, error) {
	if err := br.db.Debug().Create(&photolinkReq).Error; err != nil {
		return nil, err
	}
	return photolinkReq, nil
}

func (br *PhotolinkRepository) UpdatePhotolink(photolinkReq *model.UpdatePhotolink, id string) (*entity.Photolink, error) {
	tx := br.db.Begin()

	var photolink *entity.Photolink
	if err := tx.Debug().Where("id = ?", id).First(&photolink).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	photolink = parseUpdateReq(photolink, photolinkReq)

	if err := tx.Debug().Where("id = ?", id).Save(&photolink).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return photolink, nil
}

func (br *PhotolinkRepository) DeletePhotolink(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.Photolink{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *PhotolinkRepository) GetPhotolinkByID(id string) (*entity.Photolink, error) {
	var photolink entity.Photolink
	if err := br.db.Debug().Where("id = ?", id).Find(&photolink).Error; err != nil {
		return nil, err
	}
	return &photolink, nil
}

func (br *PhotolinkRepository) GetAllPhotolink() ([]*entity.Photolink, error) {
	var photolinks []*entity.Photolink
	if err := br.db.Debug().Find(&photolinks).Error; err != nil {
		return nil, err
	}
	return photolinks, nil
}

func parseUpdateReq(photolink *entity.Photolink, photolinkReq *model.UpdatePhotolink) *entity.Photolink {
	if photolinkReq.PhotoLink != "" {
		photolink.PhotoLink = photolinkReq.PhotoLink
	}
	// if photolinkReq.IdPreloved != uuid.Nil {
	// 	photolink.IdPreloved = photolinkReq.IdPreloved
	// }
	// if photolinkReq.IdPreloved != uuid.Nil {
	// 	photolink.IdPreloved = photolinkReq.IdPreloved
	// }
	// if photolinkReq.IdJasantar != uuid.Nil {
	// 	photolink.IdJastip = photolinkReq.IdJastip
	// }
	// if photolinkReq.IdJasantar != uuid.Nil {
	// 	photolink.IdJasantar = photolinkReq.IdJasantar
	// }
	// if photolinkReq.IdKomunitasbraw != uuid.Nil {
	// 	photolink.IdKomunitasbraw = photolinkReq.IdKomunitasbraw
	// }


	return photolink
}