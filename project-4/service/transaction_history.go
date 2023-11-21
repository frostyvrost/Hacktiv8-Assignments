package service

import (
	"project-4/models"
	"project-4/pkg"
	"project-4/repo"

	"github.com/asaskevich/govalidator"
)

type transactionServiceRepo interface {
	CreateTransaction(transaction *models.TransactionCreate, userId int) (*models.TransactionHistory, pkg.Error)
	GetTransactionsByUserID(userID uint) ([]*models.TransactionHistory, pkg.Error)
	GetAllTransaction() ([]models.TransactionHistory, pkg.Error)
}

type transactionService struct{}

var TransactionService transactionServiceRepo = &transactionService{}

func (t *transactionService) CreateTransaction(transaction *models.TransactionCreate, userId int) (*models.TransactionHistory, pkg.Error) {

	if _, err := govalidator.ValidateStruct(transaction); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	user, err := repo.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	product, err := repo.ProductRepo.GetProductById(transaction.ProductID)
	if err != nil {
		return nil, err
	}

	if product.Stock < transaction.Quantity {
		return nil, pkg.BadRequest("Insufficient product stock")
	}

	totalPrice := transaction.Quantity * product.Price

	if user.Balance < totalPrice {
		return nil, pkg.BadRequest("Insufficient user balance")
	}

	createdTransaction, err := repo.TransactionRepo.CreateTransaction(user, product, transaction, totalPrice)
	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}

func (t *transactionService) GetTransactionsByUserID(userID uint) ([]*models.TransactionHistory, pkg.Error) {
	transactionResponse, err := repo.TransactionRepo.GetTransactionsByUserID(userID)

	if err != nil {
		return nil, err
	}

	return transactionResponse, nil
}

func (t *transactionService) GetAllTransaction() ([]models.TransactionHistory, pkg.Error) {
	transactionResponse, err := repo.TransactionRepo.GetAllTransaction()

	if err != nil {
		return nil, err
	}

	return transactionResponse, nil
}
