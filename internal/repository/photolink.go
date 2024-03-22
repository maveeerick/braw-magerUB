package repository

import (
	"log"

	"braw-api/entity"

	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPhotolinkRepository interface {
	CreatePrelovedPhotos(prelovedPhotosReq *entity.PrelovedPhotos) (*entity.PrelovedPhotos, error)
	DeletePrelovedPhotos(id string) error
	GetPrelovedPhotosByID(id string) (*entity.PrelovedPhotos, error)
	CreateJastipPhotos(jastipPhotosReq *entity.JastipPhotos) (*entity.JastipPhotos, error)
	DeleteJastipPhotos(id string) error
	GetJastipPhotosByID(id string) (*entity.JastipPhotos, error)
	CreateJasantarPhotos(jasantarPhotosReq *entity.JasantarPhotos) (*entity.JasantarPhotos, error)
	DeleteJasantarPhotos(id string) error
	GetJasantarPhotosByID(id string) (*entity.JasantarPhotos, error)
	CreateKomunitasbrawPhotos(komunitasbrawPhotosReq *entity.KomunitasbrawPhotos) (*entity.KomunitasbrawPhotos, error)
	DeleteKomunitasbrawPhotos(id string) error
	GetKomunitasbrawPhotosByID(id string) (*entity.KomunitasbrawPhotos, error)
}

type PhotolinkRepository struct {
	db *gorm.DB
}

func NewPhotolinkRepository(db *gorm.DB) IPhotolinkRepository {
	return &PhotolinkRepository{
		db: db,
	}
}

func (br *PhotolinkRepository) CreatePrelovedPhotos(prelovedPhotosReq *entity.PrelovedPhotos) (*entity.PrelovedPhotos, error) {
	if err := br.db.Debug().Create(&prelovedPhotosReq).Error; err != nil {
		return nil, err
	}
	return prelovedPhotosReq, nil
}


func (br *PhotolinkRepository) DeletePrelovedPhotos(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.PrelovedPhotos{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *PhotolinkRepository) GetPrelovedPhotosByID(id string) (*entity.PrelovedPhotos, error) {
	var prelovedPhotos entity.PrelovedPhotos
	if err := br.db.Debug().Where("id = ?", id).Find(&prelovedPhotos).Error; err != nil {
		return nil, err
	}
	return &prelovedPhotos, nil
}

func (br *PhotolinkRepository) CreateJastipPhotos(jastipPhotosReq *entity.JastipPhotos) (*entity.JastipPhotos, error) {
	if err := br.db.Debug().Create(&jastipPhotosReq).Error; err != nil {
		return nil, err
	}
	return jastipPhotosReq, nil
}


func (br *PhotolinkRepository) DeleteJastipPhotos(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.JastipPhotos{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *PhotolinkRepository) GetJastipPhotosByID(id string) (*entity.JastipPhotos, error) {
	var jastipPhotos entity.JastipPhotos
	if err := br.db.Debug().Where("id = ?", id).Find(&jastipPhotos).Error; err != nil {
		return nil, err
	}
	return &jastipPhotos, nil
}

func (br *PhotolinkRepository) CreateJasantarPhotos(jasantarPhotosReq *entity.JasantarPhotos) (*entity.JasantarPhotos, error) {
	if err := br.db.Debug().Create(&jasantarPhotosReq).Error; err != nil {
		return nil, err
	}
	return jasantarPhotosReq, nil
}


func (br *PhotolinkRepository) DeleteJasantarPhotos(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.JasantarPhotos{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *PhotolinkRepository) GetJasantarPhotosByID(id string) (*entity.JasantarPhotos, error) {
	var jasantarPhotos entity.JasantarPhotos
	if err := br.db.Debug().Where("id = ?", id).Find(&jasantarPhotos).Error; err != nil {
		return nil, err
	}
	return &jasantarPhotos, nil
}

func (br *PhotolinkRepository) CreateKomunitasbrawPhotos(komunitasbrawPhotosReq *entity.KomunitasbrawPhotos) (*entity.KomunitasbrawPhotos, error) {
	if err := br.db.Debug().Create(&komunitasbrawPhotosReq).Error; err != nil {
		return nil, err
	}
	return komunitasbrawPhotosReq, nil
}


func (br *PhotolinkRepository) DeleteKomunitasbrawPhotos(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.KomunitasbrawPhotos{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *PhotolinkRepository) GetKomunitasbrawPhotosByID(id string) (*entity.KomunitasbrawPhotos, error) {
	var komunitasbrawPhotos entity.KomunitasbrawPhotos
	if err := br.db.Debug().Where("id = ?", id).Find(&komunitasbrawPhotos).Error; err != nil {
		return nil, err
	}
	return &komunitasbrawPhotos, nil
}