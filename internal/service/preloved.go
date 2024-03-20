package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/model"

	"github.com/google/uuid"
)

type IPrelovedService interface {
	CreatePreloved(prelovedReq *model.CreatePreloved) (*entity.Preloved, error)
	GetPrelovedByID(id string) (*entity.Preloved, error)
	GetPrelovedByUserID(id string) ([]*entity.Preloved, error)
	DeletePreloved(id string) error
	UpdatePreloved(prelovedReq *model.UpdatePreloved, id string) (*entity.Preloved, error)
	GetAllPreloved() ([]*entity.Preloved, error)
}

type PrelovedService struct {
	br repository.IPrelovedRepository
}

func NewPrelovedService(br repository.IPrelovedRepository) IPrelovedService {
	return &PrelovedService{
		br: br,
	}
}

func (bs *PrelovedService) CreatePreloved(prelovedReq *model.CreatePreloved) (*entity.Preloved, error) {
	prelovedParse := &entity.Preloved{
		IdPreloved:  uuid.New(),
		IdUser:      prelovedReq.IdUser,
		Title:       prelovedReq.Title,
		Category:    prelovedReq.Category,
		Price:       prelovedReq.Price,
		Condition:   prelovedReq.Condition,
		Description: prelovedReq.Description,
		//Stock:       prelovedReq.Stock,
	}

	preloved, err := bs.br.CreatePreloved(prelovedParse)
	if err != nil {
		return nil, err
	}

	return preloved, nil
}

func (bs *PrelovedService) GetPrelovedByID(id string) (*entity.Preloved, error) {
	preloved, err := bs.br.GetPrelovedByID(id)
	if err != nil {
		return nil, err
	}
	return preloved, nil
}

func (bs *PrelovedService) GetPrelovedByUserID(id string) ([]*entity.Preloved, error) {
	preloved, err := bs.br.GetPrelovedByUserID(id)
	if err != nil {
		return nil, err
	}
	return preloved, nil
}

func (bs *PrelovedService) DeletePreloved(id string) error {
	err := bs.br.DeletePreloved(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *PrelovedService) UpdatePreloved(prelovedReq *model.UpdatePreloved, id string) (*entity.Preloved, error) {
	preloved, err := bs.br.UpdatePreloved(prelovedReq, id)
	if err != nil {
		return nil, err
	}

	return preloved, nil
}

func (bs *PrelovedService) GetAllPreloved() ([]*entity.Preloved, error) {

	preloveds, err := bs.br.GetAllPreloved()
	if err != nil {
		return nil, err
	}

	return preloveds, nil
}
