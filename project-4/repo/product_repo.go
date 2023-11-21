package repo

import (
	"project-4/database"
	"project-4/models"
	"project-4/pkg"
)

type productDomainRepo interface {
	CreateProduct(*models.Product) (*models.Product, pkg.Error)
	GetAllProducts() ([]*models.Product, pkg.Error)
	UpdateProduct(*models.ProductUpdate, int) (*models.Product, pkg.Error)
	DeleteProduct(int) pkg.Error
	GetProductById(int) (*models.Product, pkg.Error)
}

type productRepo struct{}

var ProductRepo productDomainRepo = &productRepo{}

func (p *productRepo) CreateProduct(product *models.Product) (*models.Product, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&product).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return product, nil
}

func (s *productRepo) GetAllProducts() ([]*models.Product, pkg.Error) {
	db := database.GetDB()
	var products []*models.Product

	err := db.Find(&products).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return products, nil
}

func (c *productRepo) UpdateProduct(update *models.ProductUpdate, productId int) (*models.Product, pkg.Error) {
	db := database.GetDB()

	var product models.Product

	err := db.First(&product, productId).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&product).Updates(update)

	return &product, nil
}

func (s *productRepo) DeleteProduct(productId int) pkg.Error {
	db := database.GetDB()
	var product models.Product

	err := db.Where("id = ?", productId).Delete(&product).Error

	if err != nil {
		return pkg.ParseError(err)
	}

	return nil
}

func (s *productRepo) GetProductById(productId int) (*models.Product, pkg.Error) {
	db := database.GetDB()
	var product models.Product

	err := db.First(&product, productId).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return &product, nil
}
