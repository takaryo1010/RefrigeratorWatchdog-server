package repository

import (
	"RefrigeratorWatchdog-server/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IFoodRepository is an interface for managing food data.
type IFoodRepository interface {
	GetFoodByUserID(food *model.Food, userID uint) error
	CreateFood(food *model.Food) error
	UpdateFood(food *model.Food, id uint) error
	DeleteFood(id uint) error
}

type foodRepository struct {
	db *gorm.DB
}

// NewFoodRepository creates a new instance of the foodRepository struct.
func NewFoodRepository(db *gorm.DB) IFoodRepository {
	return &foodRepository{db}
}

func (fr *foodRepository) GetFoodByUserID(food *model.Food, userID uint) error {
	if err := fr.db.First(food, "user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

func (fr *foodRepository) CreateFood(food *model.Food) error {
	if err :=fr.db.Create(food).Error; err != nil {
		return err
	}
	return nil
}

func (fr *foodRepository) UpdateFood(food *model.Food, id uint) error {
	if err := fr.db.Model(food).Clauses(clause.Returning{}).Where("id = ?", id).Updates(food).Error; err != nil {
		return err
	}
	return nil
}

func (fr *foodRepository) DeleteFood(id uint) error {
	result := fr.db.Where("id = ?", id).Delete(&model.Food{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("record not found")
	}
	return nil
}
