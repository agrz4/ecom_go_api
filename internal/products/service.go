package products

import (
	"context"
	"ecom_go_api/internal/models"

	"gorm.io/gorm"
)

type Service interface {
	ListProducts(ctx context.Context) ([]models.Product, error)
}

type svc struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &svc{db: db}
}

func (s *svc) ListProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	if err := s.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
