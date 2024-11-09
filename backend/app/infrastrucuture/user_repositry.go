package infrastructure

import (
	"go_next_todo/domain/model"
	"go_next_todo/domain/repository"

	"gorm.io/gorm"
)

// UserRepository struct
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository function
func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{db}
}

// GetUserByEmail function
// Get user by email
// @param user *model.User
// @param email string
// @return error
func (ur *userRepository) GetUserByEmail(user *model.User, username string) error {
	if err := ur.db.Where("username = ?", username).First(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser function
// Create user
// @param user *model.User
// @return error
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
