package storage

import (
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/logger"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/models"
	"gorm.io/gorm"
)

type Storage struct {
	db  *gorm.DB
	log *logger.Logger
}

func NewStorage(db *gorm.DB, log *logger.Logger) *Storage {
	return &Storage{
		db:  db,
		log: log,
	}
}

func (s *Storage) CreateItem(item *models.Item) (*int64, error) {
	err := s.db.Create(item).Error
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	return item.ID, nil
}

func (s *Storage) GetItems() ([]models.Item, error) {
	var item []models.Item
	err := s.db.Model(&models.Item{}).Find(&item).Error
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	return item, nil
}

func (s *Storage) UpdateItem(item *models.Item) error {
	err := s.db.Where(item.ID).Updates(item).Error
	if err != nil {
		s.log.Error(err)
		return err
	}

	return nil
}
