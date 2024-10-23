package validator

import (
	"RefrigeratorWatchdog-server/model"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IFoodValidator interface {
	ValidateFood(food model.Food) error
}

type foodValidator struct{}

func NewFoodValidator() IFoodValidator {
	return &foodValidator{}
}

func (fv *foodValidator) ValidateFood(food model.Food) error {
	if food.Tag == "" {
		food.Tag = "その他"
	}
	return validation.ValidateStruct(&food,
		validation.Field(&food.Name,  validation.Length(1, 255)),
		validation.Field(&food.UserID, validation.Required),
		validation.Field(&food.OriginalCode,  validation.Min(0), validation.Max(10000000000000)),
		validation.Field(&food.Quantity, validation.Min(0.0), validation.Max(10000000000000.0)),
		validation.Field(&food.ExpirationDate, validation.By(allowNilTime)),
		validation.Field(&food.ImageURL,  validation.Length(0, 10000)),
		validation.Field(&food.Memo, validation.Length(0, 1000)),
		validation.Field(&food.Tag, validation.In("野菜", "肉", "魚", "乳製品", "調味料", "卵", "飲料", "果物", "加工食品", "その他","")),
	)
}

func allowNilTime(value interface{}) error {
    if value == "" {
        return nil
    }
    _, ok := value.(*time.Time)
    if !ok {
        return validation.NewError("validation_invalid", "無効な日付フォーマットです")
    }
    return nil
}
