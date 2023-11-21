package service

import (
	"project-4/models"
	"project-4/pkg"
	"project-4/repo"

	"github.com/asaskevich/govalidator"
)

type productServiceRepo interface {
	CreateProduct(*models.Product) (*models.Product, pkg.Error)
	GetAllProducts() ([]*models.Product, pkg.Error)
	UpdateProduct(*models.ProductUpdate, int) (*models.Product, pkg.Error)
	DeleteProduct(int) pkg.Error
}

type productService struct{}

var ProductService productServiceRepo = &productService{}

func (p *productService) CreateProduct(product *models.Product) (*models.Product, pkg.Error) {
	if _, err := govalidator.ValidateStruct(product); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	if product.Stock < 5 {
		return nil, pkg.BadRequest("The product stock must not be less than 5")
	}

	if _, err := repo.CategoryRepo.GetCategoryById(product.CategoryID); err != nil {
		return nil, err
	}

	productResponse, err := repo.ProductRepo.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return productResponse, nil
}

func (p *productService) GetAllProducts() ([]*models.Product, pkg.Error) {
	productResponse, err := repo.ProductRepo.GetAllProducts()

	if err != nil {
		return nil, err
	}

	return productResponse, nil
}

func (p *productService) UpdateProduct(product *models.ProductUpdate, productId int) (*models.Product, pkg.Error) {
	if _, err := govalidator.ValidateStruct(product); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	if product.Stock < 5 {
		return nil, pkg.BadRequest("The product stock must not be less than 5")
	}

	if _, err := repo.CategoryRepo.GetCategoryById(product.CategoryID); err != nil {
		return nil, err
	}

	productResponse, err := repo.ProductRepo.UpdateProduct(product, productId)

	if err != nil {
		return nil, err
	}

	return productResponse, nil
}

func (p *productService) DeleteProduct(productId int) pkg.Error {
	err := repo.ProductRepo.DeleteProduct(productId)

	if err != nil {
		return err
	}

	return nil
}
