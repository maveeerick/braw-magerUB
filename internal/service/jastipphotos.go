package service

import (
	"braw-api/entity"
	"braw-api/internal/repository"
	"braw-api/pkg/supabase"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

type IJastipImagesService interface {
	CreateJastipImage(jastipId uuid.UUID, file *multipart.FileHeader) (*entity.JastipPhotos, error)
	GetImageByID(id string) (*entity.JastipPhotos, error)
	GetImagesByJastipID(id string) ([]*entity.JastipPhotos, error)
	DeleteJastipImage(id string) error
}

type JastipImagesService struct {
	br       repository.IJastipImagesRepository
	pr       repository.IJastipRepository
	supabase supabase.Interface
}

func NewJastipImagesService(br repository.IJastipImagesRepository, pr repository.IJastipRepository, supabase supabase.Interface) IJastipImagesService {
	return &JastipImagesService{
		br:       br,
		pr:       pr,
		supabase: supabase,
	}
}

func (bs *JastipImagesService) CreateJastipImage(jastipId uuid.UUID, file *multipart.FileHeader) (*entity.JastipPhotos, error) {
	jastip, err := bs.pr.GetJastipByID(jastipId.String())
	if err != nil {
		return nil, err
	}
	if jastip == nil {
		return nil, errors.New("no jastip id found")
	}

	url, err := bs.supabase.Upload(file)
	if err != nil {
		return nil, err
	}

	image := &entity.JastipPhotos{
		IdJastipPhoto: uuid.New(),
		IdJastip:      jastipId,
		PhotoLink:     url,
	}

	imageRes, err := bs.br.CreateJastipImage(image)
	if err != nil {
		return nil, err
	}

	return imageRes, nil
}

func (bs *JastipImagesService) GetImageByID(id string) (*entity.JastipPhotos, error) {
	jastip, err := bs.br.GetImageByID(id)
	if err != nil {
		return nil, err
	}
	return jastip, nil
}

func (bs *JastipImagesService) GetImagesByJastipID(id string) ([]*entity.JastipPhotos, error) {
	jastip, err := bs.br.GetImagesByJastipID(id)
	if err != nil {
		return nil, err
	}
	return jastip, nil
}

func (bs *JastipImagesService) DeleteJastipImage(id string) error {
	image, errFind := bs.br.GetImageByID(id)
	if errFind != nil {
		return errors.New("can't find image id")
	}

	if err := bs.supabase.Delete(image.PhotoLink); err != nil {
		return err
	}

	if err := bs.br.DeleteJastipImage(id); err != nil {
		return err
	}

	return nil
}
