package database

import (
	"context"
	"go-microservices/internal/models"

	"github.com/google/uuid"
)

func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).Where(models.Product{VendorID: vendorId}).Find(&products)
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()
	result := c.DB.WithContext(ctx).Save(product)
	return product, result.Error
}
