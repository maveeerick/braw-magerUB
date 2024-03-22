package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/pkg/supabase"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

type IPrelovedImagesService interface {
	CreatePrelovedImage(prelovedId uuid.UUID, file *multipart.FileHeader) (*entity.PrelovedPhotos, error)
	GetImageByID(id string) (*entity.PrelovedPhotos, error)
	GetImagesByPrelovedID(id string) ([]*entity.PrelovedPhotos, error)
	DeletePrelovedImage(id string) error
}

type PrelovedImagesService struct {
	br       repository.IPrelovedImagesRepository
	pr       repository.IPrelovedRepository
	supabase supabase.Interface
}

func NewPrelovedImagesService(br repository.IPrelovedImagesRepository, pr repository.IPrelovedRepository, supabase supabase.Interface) IPrelovedImagesService {
	return &PrelovedImagesService{
		br:       br,
		pr:       pr,
		supabase: supabase,
	}
}

func (bs *PrelovedImagesService) CreatePrelovedImage(prelovedId uuid.UUID, file *multipart.FileHeader) (*entity.PrelovedPhotos, error) {
	preloved, err := bs.pr.GetPrelovedByID(prelovedId.String())
	if err != nil {
		return nil, err
	}
	if preloved == nil {
		return nil, errors.New("no preloved id found")
	}

	url, err := bs.supabase.Upload(file)
	if err != nil {
		return nil, err
	}

	image := &entity.PrelovedPhotos{
		IdPrelovedPhoto: uuid.New(),
		IdPreloved:      prelovedId,
		PhotoLink:       url,
	}

	imageRes, err := bs.br.CreatePrelovedImage(image)
	if err != nil {
		return nil, err
	}

	return imageRes, nil
}

func (bs *PrelovedImagesService) GetImageByID(id string) (*entity.PrelovedPhotos, error) {
	preloved, err := bs.br.GetImageByID(id)
	if err != nil {
		return nil, err
	}
	return preloved, nil
}

func (bs *PrelovedImagesService) GetImagesByPrelovedID(id string) ([]*entity.PrelovedPhotos, error) {
	preloved, err := bs.br.GetImagesByPrelovedID(id)
	if err != nil {
		return nil, err
	}
	return preloved, nil
}

func (bs *PrelovedImagesService) DeletePrelovedImage(id string) error {
	image, errFind := bs.br.GetImageByID(id)
	if errFind != nil {
		return errors.New("can't find image id")
	}

	if err := bs.supabase.Delete(image.PhotoLink); err != nil {
		return err
	}

	if err := bs.br.DeletePrelovedImage(id); err != nil {
		return err
	}

	return nil
}
