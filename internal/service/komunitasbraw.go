package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/model"
	"github.com/google/uuid"
)

type IKomunitasbrawService interface {
	CreateKomunitasbraw(komunitasbrawReq *model.CreateKomunitasbraw) (*entity.Komunitasbraw, error)
	GetKomunitasbrawByID(id string) (*entity.Komunitasbraw, error)
	GetKomunitasbrawByUserID(id string) ([]*entity.Komunitasbraw, error)
	DeleteKomunitasbraw(id string) error
	UpdateKomunitasbraw(komunitasbrawReq *model.UpdateKomunitasbraw, id string) (*entity.Komunitasbraw, error)
	GetAllKomunitasbraw() ([]*entity.Komunitasbraw, error)
}

type KomunitasbrawService struct {
	br repository.IKomunitasbrawRepository
}

func NewKomunitasbrawService(br repository.IKomunitasbrawRepository) IKomunitasbrawService {
	return &KomunitasbrawService{
		br: br,
	}
}

func (bs *KomunitasbrawService) CreateKomunitasbraw(komunitasbrawReq *model.CreateKomunitasbraw) (*entity.Komunitasbraw, error) {
	komunitasbrawParse := &entity.Komunitasbraw{
		IdUser:          komunitasbrawReq.IdUser,
		IdKomunitasbraw: uuid.New(),
		Title:       komunitasbrawReq.Title,
		Category:       komunitasbrawReq.Category,
		Description:     komunitasbrawReq.Description,
	}

	komunitasbraw, err := bs.br.CreateKomunitasbraw(komunitasbrawParse)
	if err != nil {
		return nil, err
	}

	return komunitasbraw, nil
}

func (bs *KomunitasbrawService) GetKomunitasbrawByID(id string) (*entity.Komunitasbraw, error) {
	komunitasbraw, err := bs.br.GetKomunitasbrawByID(id)
	if err != nil {
		return nil, err
	}
	return komunitasbraw, nil
}

func (bs *KomunitasbrawService) GetKomunitasbrawByUserID(id string) ([]*entity.Komunitasbraw, error) {
	komunitasbraw, err := bs.br.GetKomunitasbrawByUserID(id)
	if err != nil {
		return nil, err
	}
	return komunitasbraw, nil
}

func (bs *KomunitasbrawService) DeleteKomunitasbraw(id string) error {
	err := bs.br.DeleteKomunitasbraw(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *KomunitasbrawService) UpdateKomunitasbraw(komunitasbrawReq *model.UpdateKomunitasbraw, id string) (*entity.Komunitasbraw, error) {
	komunitasbraw, err := bs.br.UpdateKomunitasbraw(komunitasbrawReq, id)
	if err != nil {
		return nil, err
	}

	return komunitasbraw, nil
}

func (bs *KomunitasbrawService) GetAllKomunitasbraw() ([]*entity.Komunitasbraw, error) {
	

	komunitasbraws, err := bs.br.GetAllKomunitasbraw()
	if err != nil {
		return nil, err
	}

	return komunitasbraws, nil
}