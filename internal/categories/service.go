package categories

import (
	"context"
	"ecom_go_api/internal/models"

	"gorm.io/gorm"
)

type Service interface {
	ListCategories(ctx context.Context) ([]models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
}

type svc struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &svc{db: db}
}

func (s *svc) ListCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	if err := s.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *svc) CreateCategory(ctx context.Context, category *models.Category) error {
	return s.db.WithContext(ctx).Create(category).Error
}
