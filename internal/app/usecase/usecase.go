package usecase

import (
	"errors"

	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/logger"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/models"
)

type Service struct {
	storage IStorage
	log     *logger.Logger
}

func NewService(storage IStorage, log *logger.Logger) *Service {
	return &Service{
		storage: storage,
		log:     log,
	}
}

func (s *Service) CreateItem(item *models.Item) (*int64, error) {
	err := validate(item.Title, item.Description)
	if err != nil {
		return nil, err
	}

	id, err := s.storage.CreateItem(item)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *Service) GetItems() ([]models.Item, error) {
	items, err := s.storage.GetItems()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateItem(item *models.Item) error {
	err := s.storage.UpdateItem(item)
	if err != nil {
		return err
	}
	return nil
}

func validate(fields ...string) error {
	for i := range fields {
		if fields[i] == "" {
			return errors.New("required field is empty")
		}
	}
	return nil
}
