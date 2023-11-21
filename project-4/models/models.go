package models

import "time"

type User struct {
	ID                   int       `gorm:"primaryKey" json:"id,omitempty"`
	FullName             string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email                string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password             string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role                 string    `gorm:"not null" json:"role" valid:"matches(admin|customer)"`
	Balance              int       `gorm:"not null" json:"balance" valid:"range(0|100000000)"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	TransactionHistories []TransactionHistory
}

type LoginCredential struct {
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type BalanceUpdate struct {
	Balance int `gorm:"not null" json:"balance" valid:"range(0|100000000)"`
}

type UserRegister struct {
	FullName string `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type UserRegisterResponse struct {
	ID        int       `gorm:"primaryKey" json:"id,omitempty"`
	FullName  string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Balance   int       `gorm:"not null" json:"balance" valid:"range(0|100000000)"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID         int       `gorm:"primaryKey" json:"id,omitempty"`
	Title      string    `gorm:"not null" json:"title" valid:"required~Product title is required"`
	Price      int       `gorm:"not null" json:"price" valid:"required~Product price is required,range(0|50000000)"`
	Stock      int       `gorm:"not null" json:"stock" valid:"required~Product stock is required"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductCreate struct {
	Title      string `gorm:"not null" json:"title" valid:"required~Product title is required"`
	Price      int    `gorm:"not null" json:"price" valid:"required~Product price is required,range(0|50000000)"`
	Stock      int    `gorm:"not null" json:"stock" valid:"required~Product stock is required"`
	CategoryID int    `json:"category_id"`
}

type ProductUpdate struct {
	Title      string `gorm:"not null" json:"title" valid:"required~Product title is required"`
	Price      int    `gorm:"not null" json:"price" valid:"required~Product price is required,range(0|50000000)"`
	Stock      int    `gorm:"not null" json:"stock" valid:"required~Product stock is required"`
	CategoryID int    `json:"category_id"`
}

type Category struct {
	ID                int       `gorm:"primaryKey" json:"id,omitempty"`
	Type              string    `gorm:"not null" json:"type" valid:"required~Product type is required"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Products          []Product
}

type CategoryCreate struct {
	Type string `gorm:"not null" json:"type" valid:"required~Product type is required"`
}

type CategoryUpdate struct {
	Type string `gorm:"not null" json:"type" valid:"required~Product type is required"`
}

type TransactionHistory struct {
	ID         int       `gorm:"primaryKey" json:"id,omitempty"`
	Quantity   int       `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
	TotalPrice int       `gorm:"not null" json:"total_price" valid:"required~Total price is required"`
	ProductID  int       `json:"product_id"`
	UserID     int       `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Product    *Product
	User       *User
}

type TransactionCreate struct {
	ProductID int `json:"product_id"`
	Quantity  int `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
}

type TransactionCreateResponse struct {
	Message string `json:"message"`

	TransactionBill struct {
		TotalPrice   int    `gorm:"not null" json:"total_price" valid:"required~Total price is required"`
		Quantity     int    `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
		ProductTitle string `json:"product_title"`
	} `json:"transaction_bill"`
}

type GetTransactionUserResponse struct {
	ID         int       `gorm:"primaryKey" json:"id,omitempty"`
	Quantity   int       `gorm:"not null" json:"quantity" valid:"required~Product quantity is required"`
	TotalPrice int       `gorm:"not null" json:"total_price" valid:"required~Total price is required"`
	ProductID  int       `json:"product_id"`
	UserID     int       `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Product    *Product
}
