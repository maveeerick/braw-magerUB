package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/model"
	"braw-api/pkg/utils"

	"github.com/google/uuid"
)

type IJastipService interface {
	CreateJastip(jastipReq *model.CreateJastip) (*entity.Jastip, error)
	GetJastipByID(id string) (*entity.Jastip, error)
	DeleteJastip(id string) error
	UpdateJastip(jastipReq *model.UpdateJastip, id string) (*entity.Jastip, error)
	GetAllJastip() ([]*entity.Jastip, error)
}

type JastipService struct {
	br repository.IJastipRepository
}

func NewJastipService(br repository.IJastipRepository) IJastipService {
	return &JastipService{
		br: br,
	}
}

func (bs *JastipService) CreateJastip(jastipReq *model.CreateJastip) (*entity.Jastip, error) {
	closed, _ := utils.ParseTime(jastipReq.Close_Order)
	jastipParse := &entity.Jastip{
		ID_Jastip:   uuid.New(),
		ID_User:     uuid.New(),
		Title:       jastipReq.Title,
		Category:    jastipReq.Category,
		Price:       jastipReq.Price,
		Open_Day:    jastipReq.Open_Day,
		Close_Order: closed,
		//Stock:       jastipReq.Stock,
	}

	jastip, err := bs.br.CreateJastip(jastipParse)
	if err != nil {
		return nil, err
	}

	return jastip, nil
}

func (bs *JastipService) GetJastipByID(id string) (*entity.Jastip, error) {
	jastip, err := bs.br.GetJastipByID(id)
	if err != nil {
		return nil, err
	}
	return jastip, nil
}

func (bs *JastipService) DeleteJastip(id string) error {
	err := bs.br.DeleteJastip(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *JastipService) UpdateJastip(jastipReq *model.UpdateJastip, id string) (*entity.Jastip, error) {
	jastip, err := bs.br.UpdateJastip(jastipReq, id)
	if err != nil {
		return nil, err
	}

	return jastip, nil
}

func (bs *JastipService) GetAllJastip() ([]*entity.Jastip, error) {

	jastips, err := bs.br.GetAllJastip()
	if err != nil {
		return nil, err
	}

	return jastips, nil
}
