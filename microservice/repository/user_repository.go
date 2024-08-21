package repository

import (
	"cc.tech/main/db"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *db.User) error {
	return r.db.Create(user).Error
}

// Add other CRUD operations...
func foo() {
	bar := CreateUser()
}