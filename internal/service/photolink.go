package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/model"
	"github.com/google/uuid"
)

type IPhotolinkService interface {
	CreatePhotolink(photolinkReq *model.CreatePhotolink) (*entity.Photolink, error)
	GetPhotolinkByID(id string) (*entity.Photolink, error)
	DeletePhotolink(id string) error
	UpdatePhotolink(photolinkReq *model.UpdatePhotolink, id string) (*entity.Photolink, error)
	GetAllPhotolink() ([]*entity.Photolink, error)
}

type PhotolinkService struct {
	br repository.IPhotolinkRepository
}

func NewPhotolinkService(br repository.IPhotolinkRepository) IPhotolinkService {
	return &PhotolinkService{
		br: br,
	}
}

func (bs *PhotolinkService) CreatePhotolink(photolinkReq *model.CreatePhotolink) (*entity.Photolink, error) {
	photolinkParse := &entity.Photolink{
		IdPhotolink:          uuid.New(),
		PhotoLink:       photolinkReq.PhotoLink,
		// IdPreloved:     photolinkReq.IdPreloved,
		// IdJastip:        photolinkReq.IdJastip,
		// IdJasantar:       photolinkReq.IdJasantar,
		// IdKomunitasbraw: photolinkReq.IdKomunitasbraw,
	}

	photolink, err := bs.br.CreatePhotolink(photolinkParse)
	if err != nil {
		return nil, err
	}

	return photolink, nil
}

func (bs *PhotolinkService) GetPhotolinkByID(id string) (*entity.Photolink, error) {
	photolink, err := bs.br.GetPhotolinkByID(id)
	if err != nil {
		return nil, err
	}
	return photolink, nil
}

func (bs *PhotolinkService) DeletePhotolink(id string) error {
	err := bs.br.DeletePhotolink(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *PhotolinkService) UpdatePhotolink(photolinkReq *model.UpdatePhotolink, id string) (*entity.Photolink, error) {
	photolink, err := bs.br.UpdatePhotolink(photolinkReq, id)
	if err != nil {
		return nil, err
	}

	return photolink, nil
}

func (bs *PhotolinkService) GetAllPhotolink() ([]*entity.Photolink, error) {
	

	photolinks, err := bs.br.GetAllPhotolink()
	if err != nil {
		return nil, err
	}

	return photolinks, nil
}