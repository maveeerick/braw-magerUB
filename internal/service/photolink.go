package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/model"
	"github.com/google/uuid"
)

type IPhotolinkService interface {
	CreatePrelovedPhotos(prelovedPhotosReq *model.CreatePrelovedPhotos) (*entity.PrelovedPhotos, error)
	GetPrelovedPhotosByID(id string) (*entity.PrelovedPhotos, error)
	DeletePrelovedPhotos(id string) error
	CreateJastipPhotos(jastipPhotosReq *model.CreateJastipPhotos) (*entity.JastipPhotos, error)
	GetJastipPhotosByID(id string) (*entity.JastipPhotos, error)
	DeleteJastipPhotos(id string) error
	CreateJasantarPhotos(jasantarPhotosReq *model.CreateJasantarPhotos) (*entity.JasantarPhotos, error)
	GetJasantarPhotosByID(id string) (*entity.JasantarPhotos, error)
	DeleteJasantarPhotos(id string) error
	CreateKomunitasbrawPhotos(komunitasbrawPhotosReq *model.CreateKomunitasbrawPhotos) (*entity.KomunitasbrawPhotos, error)
	GetKomunitasbrawPhotosByID(id string) (*entity.KomunitasbrawPhotos, error)
	DeleteKomunitasbrawPhotos(id string) error
}

type PhotolinkService struct {
	br repository.IPhotolinkRepository
}

func NewPhotolinkService(br repository.IPhotolinkRepository) IPhotolinkService {
	return &PhotolinkService{
		br: br,
	}
}

func (bs *PhotolinkService) CreatePrelovedPhotos(prelovedPhotosReq *model.CreatePrelovedPhotos) (*entity.PrelovedPhotos, error) {
	prelovedPhotosParse := &entity.PrelovedPhotos{
		IdPrelovedPhoto:          uuid.New(),
		PhotoLink:       prelovedPhotosReq.PhotoLink,
	}

	prelovedphotos, err := bs.br.CreatePrelovedPhotos(prelovedPhotosParse)
	if err != nil {
		return nil, err
	}

	return prelovedphotos, nil
}

func (bs *PhotolinkService) GetPrelovedPhotosByID(id string) (*entity.PrelovedPhotos, error) {
	prelovedPhotos, err := bs.br.GetPrelovedPhotosByID(id)
	if err != nil {
		return nil, err
	}
	return prelovedPhotos, nil
}

func (bs *PhotolinkService) DeletePrelovedPhotos(id string) error {
	err := bs.br.DeletePrelovedPhotos(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *PhotolinkService) CreateJastipPhotos(jastipPhotosReq *model.CreateJastipPhotos) (*entity.JastipPhotos, error) {
	jastipPhotosParse := &entity.JastipPhotos{
		IdJastipPhoto:          uuid.New(),
		PhotoLink:       jastipPhotosReq.PhotoLink,
	}

	jastipphotos, err := bs.br.CreateJastipPhotos(jastipPhotosParse)
	if err != nil {
		return nil, err
	}

	return jastipphotos, nil
}

func (bs *PhotolinkService) GetJastipPhotosByID(id string) (*entity.JastipPhotos, error) {
	jastipPhotos, err := bs.br.GetJastipPhotosByID(id)
	if err != nil {
		return nil, err
	}
	return jastipPhotos, nil
}

func (bs *PhotolinkService) DeleteJastipPhotos(id string) error {
	err := bs.br.DeleteJastipPhotos(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *PhotolinkService) CreateJasantarPhotos(jasantarPhotosReq *model.CreateJasantarPhotos) (*entity.JasantarPhotos, error) {
	jasantarPhotosParse := &entity.JasantarPhotos{
		IdJasantarPhoto:          uuid.New(),
		PhotoLink:       jasantarPhotosReq.PhotoLink,
	}

	jasantarphotos, err := bs.br.CreateJasantarPhotos(jasantarPhotosParse)
	if err != nil {
		return nil, err
	}

	return jasantarphotos, nil
}

func (bs *PhotolinkService) GetJasantarPhotosByID(id string) (*entity.JasantarPhotos, error) {
	jasantarPhotos, err := bs.br.GetJasantarPhotosByID(id)
	if err != nil {
		return nil, err
	}
	return jasantarPhotos, nil
}

func (bs *PhotolinkService) DeleteJasantarPhotos(id string) error {
	err := bs.br.DeleteJasantarPhotos(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *PhotolinkService) CreateKomunitasbrawPhotos(komunitasbrawPhotosReq *model.CreateKomunitasbrawPhotos) (*entity.KomunitasbrawPhotos, error) {
	komunitasbrawPhotosParse := &entity.KomunitasbrawPhotos{
		IdKomunitasbrawphoto:          uuid.New(),
		PhotoLink:       komunitasbrawPhotosReq.PhotoLink,
	}

	komunitasbrawphotos, err := bs.br.CreateKomunitasbrawPhotos(komunitasbrawPhotosParse)
	if err != nil {
		return nil, err
	}

	return komunitasbrawphotos, nil
}

func (bs *PhotolinkService) GetKomunitasbrawPhotosByID(id string) (*entity.KomunitasbrawPhotos, error) {
	komunitasbrawPhotos, err := bs.br.GetKomunitasbrawPhotosByID(id)
	if err != nil {
		return nil, err
	}
	return komunitasbrawPhotos, nil
}

func (bs *PhotolinkService) DeleteKomunitasbrawPhotos(id string) error {
	err := bs.br.DeleteKomunitasbrawPhotos(id)
	if err != nil {
		return err
	}
	return nil
}