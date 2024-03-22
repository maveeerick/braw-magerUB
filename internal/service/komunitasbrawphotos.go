package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/pkg/supabase"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

type IKomunitasbrawImagesService interface {
	CreateKomunitasbrawImage(komunitasbrawId uuid.UUID, file *multipart.FileHeader) (*entity.KomunitasbrawPhotos, error)
	GetImageByID(id string) (*entity.KomunitasbrawPhotos, error)
	GetImagesByKomunitasbrawID(id string) ([]*entity.KomunitasbrawPhotos, error)
	DeleteKomunitasbrawImage(id string) error
}

type KomunitasbrawImagesService struct {
	br       repository.IKomunitasbrawImagesRepository
	pr       repository.IKomunitasbrawRepository
	supabase supabase.Interface
}

func NewKomunitasbrawImagesService(br repository.IKomunitasbrawImagesRepository, pr repository.IKomunitasbrawRepository, supabase supabase.Interface) IKomunitasbrawImagesService {
	return &KomunitasbrawImagesService{
		br:       br,
		pr:       pr,
		supabase: supabase,
	}
}

func (bs *KomunitasbrawImagesService) CreateKomunitasbrawImage(komunitasbrawId uuid.UUID, file *multipart.FileHeader) (*entity.KomunitasbrawPhotos, error) {
	komunitasbraw, err := bs.pr.GetKomunitasbrawByID(komunitasbrawId.String())
	if err != nil {
		return nil, err
	}
	if komunitasbraw == nil {
		return nil, errors.New("no Komunitasbraw id found")
	}

	url, err := bs.supabase.Upload(file)
	if err != nil {
		return nil, err
	}

	image := &entity.KomunitasbrawPhotos{
		IdKomunitasbrawphoto: uuid.New(),
		IdKomunitasbraw:      komunitasbrawId,
		PhotoLink:            url,
	}

	imageRes, err := bs.br.CreateKomunitasbrawImage(image)
	if err != nil {
		return nil, err
	}

	return imageRes, nil
}

func (bs *KomunitasbrawImagesService) GetImageByID(id string) (*entity.KomunitasbrawPhotos, error) {
	komunitasbraw, err := bs.br.GetImageByID(id)
	if err != nil {
		return nil, err
	}
	return komunitasbraw, nil
}

func (bs *KomunitasbrawImagesService) GetImagesByKomunitasbrawID(id string) ([]*entity.KomunitasbrawPhotos, error) {
	komunitasbraw, err := bs.br.GetImagesByKomunitasbrawID(id)
	if err != nil {
		return nil, err
	}
	return komunitasbraw, nil
}

func (bs *KomunitasbrawImagesService) DeleteKomunitasbrawImage(id string) error {
	image, errFind := bs.br.GetImageByID(id)
	if errFind != nil {
		return errors.New("can't find image id")
	}

	if err := bs.supabase.Delete(image.PhotoLink); err != nil {
		return err
	}

	if err := bs.br.DeleteKomunitasbrawImage(id); err != nil {
		return err
	}

	return nil
}
