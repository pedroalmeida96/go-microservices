package database

import (
	"context"
	"errors"
	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).Find(&vendors)
	return vendors, result.Error
}

func (c Client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	vendor.VendorID = uuid.NewString()
	result := c.DB.WithContext(ctx).Save(vendor)
	return vendor, result.Error
}

func (c Client) GetVendorById(ctx context.Context, ID string) (*models.Vendor, error) {
	//pointer/address in memory to some empty vendor in memory
	vendor := &models.Vendor{}

	//query db for some ID and store in that pointer/address what was retrieved
	result := c.DB.WithContext(ctx).
		Where(models.Vendor{VendorID: ID}).
		First(&vendor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "vendor", ID: ID}
		}
		return nil, result.Error
	}
	return vendor, nil
}

func (c Client) UpdateVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	result := c.DB.WithContext(ctx).
		Model(&vendor).
		Where(models.Vendor{VendorID: vendor.VendorID}).
		Clauses(clause.Returning{}).
		Updates(models.Vendor{
			Name:    vendor.Name,
			Contact: vendor.Contact,
			Phone:   vendor.Phone,
			Address: vendor.Address,
			Email:   vendor.Email,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "vendor", ID: vendor.VendorID}
	}

	return vendor, nil
}

func (c Client) DeleteVendor(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(models.Vendor{VendorID: ID}).Error
}
