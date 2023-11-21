package repo

import (
	"project-4/database"
	"project-4/models"
	"project-4/pkg"
)

type transactionDomainRepo interface {
	CreateTransaction(
		user *models.User,
		product *models.Product,
		transaction *models.TransactionCreate,
		totalPrice int) (*models.TransactionHistory, pkg.Error)

	GetTransactionsByUserID(userID uint) ([]*models.TransactionHistory, pkg.Error)
	GetAllTransaction() ([]models.TransactionHistory, pkg.Error)
}

type transactionRepo struct{}

var TransactionRepo transactionDomainRepo = &transactionRepo{}

func (t *transactionRepo) CreateTransaction(
	user *models.User,
	product *models.Product,
	transaction *models.TransactionCreate,
	totalPrice int) (*models.TransactionHistory, pkg.Error) {
	db := database.GetDB()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	product.Stock -= transaction.Quantity
	user.Balance -= totalPrice

	if err := tx.Save(product).Error; err != nil {
		tx.Rollback()
		return nil, pkg.InternalServerError("Failed to update product stock")
	}

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return nil, pkg.InternalServerError("Failed to update user balance")
	}

	category := &models.Category{} // Inisialisasi objek kategori
	if err := tx.Where("id = ?", product.CategoryID).First(category).Error; err != nil {
		tx.Rollback()
		return nil, pkg.InternalServerError("Failed to find category")
	}

	category.SoldProductAmount += transaction.Quantity

	if err := tx.Save(category).Error; err != nil {
		tx.Rollback()
		return nil, pkg.InternalServerError("Failed to update category sold product amount")
	}

	transactionHistory := &models.TransactionHistory{
		Quantity:   transaction.Quantity,
		TotalPrice: totalPrice,
		UserID:     user.ID,
		ProductID:  product.ID,
		Product:    product,
	}

	if err := tx.Create(transactionHistory).Error; err != nil {
		tx.Rollback()
		return nil, pkg.InternalServerError("Failed to create new transaction: " + err.Error())
	}

	if err := tx.Model(user).Association("TransactionHistories").Append(transactionHistory); err != nil {
		tx.Rollback()
		return nil, pkg.InternalServerError("Failed to update user's transaction histories: ")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, pkg.InternalServerError("Failed to commit transaction")
	}

	return transactionHistory, nil
}

func (t *transactionRepo) GetTransactionsByUserID(userID uint) ([]*models.TransactionHistory, pkg.Error) {
	db := database.GetDB()

	var transactionUser []*models.TransactionHistory

	err := db.Preload("Product").Where("user_id = ?", userID).Find(&transactionUser).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return transactionUser, nil
}

func (t *transactionRepo) GetAllTransaction() ([]models.TransactionHistory, pkg.Error) {
	db := database.GetDB()

	var AllTransaction []models.TransactionHistory

	err := db.Preload("Product").Preload("User").Find(&AllTransaction).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}
	return AllTransaction, nil
}
