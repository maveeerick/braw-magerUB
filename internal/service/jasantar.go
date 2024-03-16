package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/model"
	"github.com/google/uuid"
)

type IJasantarService interface {
	CreateJasantar(jasantarReq *model.CreateJasantar) (*entity.Jasantar, error)
	GetJasantarByID(id string) (*entity.Jasantar, error)
	DeleteJasantar(id string) error
	UpdateJasantar(jasantarReq *model.UpdateJasantar, id string) (*entity.Jasantar, error)
	GetAllJasantar() ([]*entity.Jasantar, error)
}

type JasantarService struct {
	br repository.IJasantarRepository
}

func NewJasantarService(br repository.IJasantarRepository) IJasantarService {
	return &JasantarService{
		br: br,
	}
}

func (bs *JasantarService) CreateJasantar(jasantarReq *model.CreateJasantar) (*entity.Jasantar, error) {
	jasantarParse := &entity.Jasantar{
		ID_Jasantar:          uuid.New(),
		ID_User: uuid.New(),
		Title:       jasantarReq.Title,
		Area:     jasantarReq.Area,
		Category:        jasantarReq.Category,
	}

	jasantar, err := bs.br.CreateJasantar(jasantarParse)
	if err != nil {
		return nil, err
	}

	return jasantar, nil
}

func (bs *JasantarService) GetJasantarByID(id string) (*entity.Jasantar, error) {
	jasantar, err := bs.br.GetJasantarByID(id)
	if err != nil {
		return nil, err
	}
	return jasantar, nil
}

func (bs *JasantarService) DeleteJasantar(id string) error {
	err := bs.br.DeleteJasantar(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *JasantarService) UpdateJasantar(jasantarReq *model.UpdateJasantar, id string) (*entity.Jasantar, error) {
	jasantar, err := bs.br.UpdateJasantar(jasantarReq, id)
	if err != nil {
		return nil, err
	}

	return jasantar, nil
}

func (bs *JasantarService) GetAllJasantar() ([]*entity.Jasantar, error) {
	

	jasantars, err := bs.br.GetAllJasantar()
	if err != nil {
		return nil, err
	}

	return jasantars, nil
}