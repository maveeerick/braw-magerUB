package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/pkg/supabase"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

type IJasantarImagesService interface {
	CreateJasantarImage(jasantarId uuid.UUID, file *multipart.FileHeader) (*entity.JasantarPhotos, error)
	GetImageByID(id string) (*entity.JasantarPhotos, error)
	GetImagesByJasantarID(id string) ([]*entity.JasantarPhotos, error)
	DeleteJasantarImage(id string) error
}

type JasantarImagesService struct {
	br       repository.IJasantarImagesRepository
	pr       repository.IJasantarRepository
	supabase supabase.Interface
}

func NewJasantarImagesService(br repository.IJasantarImagesRepository, pr repository.IJasantarRepository, supabase supabase.Interface) IJasantarImagesService {
	return &JasantarImagesService{
		br:       br,
		pr:       pr,
		supabase: supabase,
	}
}

func (bs *JasantarImagesService) CreateJasantarImage(jasantarId uuid.UUID, file *multipart.FileHeader) (*entity.JasantarPhotos, error) {
	Jasantar, err := bs.pr.GetJasantarByID(jasantarId.String())
	if err != nil {
		return nil, err
	}
	if Jasantar == nil {
		return nil, errors.New("no Jasantar id found")
	}

	url, err := bs.supabase.Upload(file)
	if err != nil {
		return nil, err
	}

	image := &entity.JasantarPhotos{
		IdJasantarPhoto: uuid.New(),
		IdJasantar:      jasantarId,
		PhotoLink:       url,
	}

	imageRes, err := bs.br.CreateJasantarImage(image)
	if err != nil {
		return nil, err
	}

	return imageRes, nil
}

func (bs *JasantarImagesService) GetImageByID(id string) (*entity.JasantarPhotos, error) {
	Jasantar, err := bs.br.GetImageByID(id)
	if err != nil {
		return nil, err
	}
	return Jasantar, nil
}

func (bs *JasantarImagesService) GetImagesByJasantarID(id string) ([]*entity.JasantarPhotos, error) {
	Jasantar, err := bs.br.GetImagesByJasantarID(id)
	if err != nil {
		return nil, err
	}
	return Jasantar, nil
}

func (bs *JasantarImagesService) DeleteJasantarImage(id string) error {
	image, errFind := bs.br.GetImageByID(id)
	if errFind != nil {
		return errors.New("can't find image id")
	}

	if err := bs.supabase.Delete(image.PhotoLink); err != nil {
		return err
	}

	if err := bs.br.DeleteJasantarImage(id); err != nil {
		return err
	}

	return nil
}
