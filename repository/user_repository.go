package repository

import (
	"RefrigeratorWatchdog-server/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IUserRepository is an interface for user repository.
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	UpdateUser(user *model.User,email string) error
	DeleteUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of the userRepository struct.
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err

	}
	return nil
}

func (ur *userRepository) UpdateUser(user *model.User,email string) error {
	result := ur.db.Model(user).Clauses(clause.Returning{}).Where("email = ?", email).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

func (ur *userRepository) DeleteUser(user *model.User) error {
	result := ur.db.Where("email = ?", user.Email).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
