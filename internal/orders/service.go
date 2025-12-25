package orders

import (
	"context"
	"ecom_go_api/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has not enough stock")
)

type svc struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &svc{
		db: db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (models.Order, error) {
	// validate payload
	if tempOrder.CustomerID == 0 {
		return models.Order{}, fmt.Errorf("customer ID is required")
	}
	if len(tempOrder.Items) == 0 {
		return models.Order{}, fmt.Errorf("at least one item is required")
	}

	var order models.Order

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// create an order
		order = models.Order{CustomerID: tempOrder.CustomerID}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// look for the product if exist
		for _, item := range tempOrder.Items {
			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrProductNotFound
				}
				return err
			}

			if product.Quantity < item.Quantity {
				return ErrProductNoStock
			}

			// create order item
			orderItem := models.OrderItem{
				OrderID:    order.ID,
				ProductID:  item.ProductID,
				Quantity:   item.Quantity,
				PriceCents: product.PriceInCenters,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (s *svc) GetOrder(ctx context.Context, id int64) (models.Order, error) {
	var order models.Order
	// Preload Items to include OrderItems in response
	if err := s.db.WithContext(ctx).Preload("Items").First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Order{}, fmt.Errorf("order not found")
		}
		return models.Order{}, err
	}
	return order, nil
}
