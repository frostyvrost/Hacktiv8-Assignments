package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FullName  string    `gorm:"not null" json:"full_name" valid:"required"`
	Email     string    `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password  string    `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
	Role      string    `gorm:"not null" json:"role" valid:"required,in(admin|member)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Task
}

type UserCreate struct {
	FullName string `gorm:"not null" json:"full_name" valid:"required"`
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
}

type UserUpdate struct {
	FullName string `gorm:"not null" json:"full_name"`
	Email    string `gorm:"not null;unique" json:"email" valid:"email"`
}

type LoginCredential struct {
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
}

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey" `
	Title       string    `json:"title" gorm:"not null" valid:"required"`
	Description string    `json:"description" gorm:"not null" valid:"required"`
	Status      bool      `json:"status" gorm:"not null"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        *User
}

type TaskCreate struct {
	Title       string `json:"title" gorm:"not null" valid:"required"`
	Description string `json:"description" gorm:"not null" valid:"required"`
	CategoryID  uint   `json:"category_id"`
}

type TaskUpdate struct {
	Title       string `json:"title" gorm:"not null" valid:"required"`
	Description string `json:"description" gorm:"not null" valid:"required"`
}

type TaskStatusUpdate struct {
	Status bool `json:"status" gorm:"not null"`
}

type TaskCategoryUpdate struct {
	CategoryID uint `json:"category_id"`
}

type TaskCreateResponse struct {
	ID          uint      `json:"id" gorm:"primaryKey" `
	Title       string    `json:"title" gorm:"not null" valid:"required"`
	Status      bool      `json:"status" gorm:"not null"`
	Description string    `json:"description" gorm:"not null" valid:"required"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskUpdateResponse struct {
	ID          uint      `json:"id" gorm:"primaryKey" `
	Title       string    `json:"title" gorm:"not null" valid:"required"`
	Description string    `json:"description" gorm:"not null" valid:"required"`
	Status      bool      `json:"status" gorm:"not null"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey" `
	Type      string    `json:"type" gorm:"not null" valid:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tasks     []Task
}

type CategoryCreate struct {
	Type string `json:"type" gorm:"not null" valid:"required"`
}

type CategoryUpdate struct {
	Type string `json:"type" gorm:"not null" valid:"required"`
}
