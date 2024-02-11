package usecase

import "github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/models"

type IStorage interface {
	CreateItem(item *models.Item) (*int64, error)
	GetItems() ([]models.Item, error)
	UpdateItem(item *models.Item) error
}

type IService interface {
	CreateItem(item *models.Item) (*int64, error)
	GetItems() ([]models.Item, error)
	UpdateItem(item *models.Item) error
}
