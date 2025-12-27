package products

import (
	"context"
	"ecom_go_api/internal/models"

	"gorm.io/gorm"
)

type Service interface {
	ListProducts(ctx context.Context, categoryID string) ([]models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) error
	UpdateProduct(ctx context.Context, id int64, product *models.Product) error
	DeleteProduct(ctx context.Context, id int64) error
}

type svc struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &svc{db: db}
}

func (s *svc) ListProducts(ctx context.Context, categoryID string) ([]models.Product, error) {
	var products []models.Product
	query := s.db.WithContext(ctx).Preload("Category")
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *svc) CreateProduct(ctx context.Context, product *models.Product) error {
	return s.db.WithContext(ctx).Create(product).Error
}

func (s *svc) UpdateProduct(ctx context.Context, id int64, product *models.Product) error {
	return s.db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (s *svc) DeleteProduct(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&models.Product{}, id).Error
}
