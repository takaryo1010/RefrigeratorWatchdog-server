package usecase

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/repository"
	"RefrigeratorWatchdog-server/repository/mocks"
	"RefrigeratorWatchdog-server/validator"
	"reflect"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func Test_foodUsecase_GetFoodsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFoodRepository(ctrl)
	Validator := validator.NewFoodValidator()

	type fields struct {
		fr repository.IFoodRepository
		fv validator.IFoodValidator
	}
	type args struct {
		userID uint
		foods  []model.Food
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.FoodResponse
		wantErr bool
	}{
		{
			name: "正常系：ユーザーIDに紐づく食材を取得できる",
			fields: fields{
				fr: mockRepo,
				fv: Validator,
			},
			args: args{
				userID: 1,
				foods: []model.Food{
					{
						ID:             1,
						Name:           "food1",
						UserID:         1,
						OriginalCode:   123,
						Quantity:       1,
						CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						ImageURL:       "https://example.com",
						Memo:           "memo",
					},
					{
						ID:             2,
						Name:           "food2",
						UserID:         1,
						OriginalCode:   123,
						Quantity:       1,
						CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						ImageURL:       "https://example.com",
						Memo:           "memo",
					},
				},
			},
			want: []model.FoodResponse{
				{
					ID:             1,
					Name:           "food1",
					UserID:         1,
					OriginalCode:   123,
					Quantity:       1,
					CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ImageURL:       "https://example.com",
					Memo:           "memo",
				},
				{
					ID:             2,
					Name:           "food2",
					UserID:         1,
					OriginalCode:   123,
					Quantity:       1,
					CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ImageURL:       "https://example.com",
					Memo:           "memo",
				},
			},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーIDに紐づく食材が取得できない",
			fields: fields{
				fr: mockRepo,
				fv: Validator,
			},
			args: args{
				userID: 0,
				foods:  []model.Food{},
			},
			want:    []model.FoodResponse{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := &foodUsecase{
				fr: tt.fields.fr,
				fv: tt.fields.fv,
			}
			mockRepo.EXPECT().GetFoodsByUserID(gomock.Any(), tt.args.userID).Do(func(foods *[]model.Food, userID uint) {
				*foods = tt.args.foods
			}).Return(nil).Times(1)

			got, err := fu.GetFoodsByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodUsecase.GetFoodsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foodUsecase.GetFoodsByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foodUsecase_CreateFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFoodRepository(ctrl)
	Validator := validator.NewFoodValidator()

	type fields struct {
		fr repository.IFoodRepository
		fv validator.IFoodValidator
	}
	type args struct {
		food model.Food
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.FoodResponse
		wantErr bool
	}{
		{
			name: "正常系：食材を作成できる",
			fields: fields{
				fr: mockRepo,
				fv: Validator,
			},
			args: args{
				food: model.Food{
					ID:             1,
					Name:           "food1",
					UserID:         1,
					OriginalCode:   123,
					Quantity:       1,
					CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ImageURL:       "https://example.com",
					Memo:           "memo",
				},
			},
			want: model.FoodResponse{
				ID:             1,
				Name:           "food1",
				UserID:         1,
				OriginalCode:   123,
				Quantity:       1,
				CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				ImageURL:       "https://example.com",
				Memo:           "memo",
			},
			wantErr: false,
		},
		{
			name: "異常系：バリデーションエラー",
			fields: fields{
				fr: mockRepo,
				fv: Validator,
			},
			args: args{
				food: model.Food{
					ID:             1,
					Name:           "",
					User:           model.User{},
					OriginalCode:   123,
					Quantity:       1,
					CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ImageURL:       "https://example.com",
					Memo:           "memo",
				},
			},
			want:    model.FoodResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := &foodUsecase{
				fr: tt.fields.fr,
				fv: tt.fields.fv,
			}

			if tt.wantErr {
				got, err := fu.CreateFood(tt.args.food)
				if (err != nil) != tt.wantErr {
					t.Errorf("foodUsecase.CreateFood() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("foodUsecase.CreateFood() = %v, want %v", got, tt.want)
				}
				return 
			}

		
			mockRepo.EXPECT().CreateFood(gomock.Any()).Do(func(food *model.Food) {
				*food = tt.args.food 
			}).Return(nil).Times(1)

			got, err := fu.CreateFood(tt.args.food)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodUsecase.CreateFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foodUsecase.CreateFood() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_foodUsecase_UpdateFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFoodRepository(ctrl)
	Validator := validator.NewFoodValidator()

	type fields struct {
		fr repository.IFoodRepository
		fv validator.IFoodValidator
	}
	type args struct {
		food model.Food
		id   uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.FoodResponse
		wantErr bool
	}{
		{
			name: "正常系：食材を更新できる",
			fields: fields{
				fr: mockRepo,
				fv: Validator,
			},
			args: args{
				food: model.Food{
					ID:             1,
					Name:           "food1",
					UserID:         1,
					OriginalCode:   123,
					Quantity:       1,
					CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ImageURL:       "https://example.com",
					Memo:           "memo",
				},
				id: 1,
			},
			want: model.FoodResponse{
				ID:             1,
				Name:           "food1",
				UserID:         1,
				OriginalCode:   123,
				Quantity:       1,
				CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				ImageURL:       "https://example.com",
				Memo:           "memo",
			},
			wantErr: false,
		},
		{
			name: "異常系：バリデーションエラー",
			fields: fields{
				fr: mockRepo,
				fv: Validator,
			},
			args: args{
				food: model.Food{
					ID: 		   1,
					Name:           "",
					UserID:         1,
					OriginalCode:   123,
					Quantity:       1,
					CreatedAt:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					ExpirationDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),

					ImageURL:       "https://example.com",
					Memo:           "memo",
				},
				id: 1,
			},
			want:    model.FoodResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := &foodUsecase{
				fr: tt.fields.fr,
				fv: tt.fields.fv,
			}

			if tt.wantErr {
				got, err := fu.UpdateFood(tt.args.food, tt.args.id)
				if (err != nil) != tt.wantErr {
					t.Errorf("foodUsecase.UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("foodUsecase.UpdateFood() = %v, want %v", got, tt.want)
				}
				return 
			}

			mockRepo.EXPECT().UpdateFood(gomock.Any(), tt.args.id).Do(func(food *model.Food, id uint) {
				*food = tt.args.food
			}).Return(nil).Times(1)

			got, err := fu.UpdateFood(tt.args.food, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodUsecase.UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foodUsecase.UpdateFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foodUsecase_DeleteFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFoodRepository(ctrl)

	type fields struct {
		fr repository.IFoodRepository
		fv validator.IFoodValidator
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常系：食材を削除できる",
			fields: fields{
				fr: mockRepo,
				fv: validator.NewFoodValidator(),
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "異常系：食材を削除できない",
			fields: fields{
				fr: mockRepo,
				fv: validator.NewFoodValidator(),
			},
			args: args{
				id: 0,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := &foodUsecase{
				fr: tt.fields.fr,
				fv: tt.fields.fv,
			}

			if tt.wantErr {
				mockRepo.EXPECT().DeleteFood(tt.args.id).Return(nil).Times(1)
				if err := fu.DeleteFood(tt.args.id); err == nil {
					t.Errorf("foodUsecase.DeleteFood() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			mockRepo.EXPECT().DeleteFood(tt.args.id).Return(nil).Times(1)
			if err := fu.DeleteFood(tt.args.id); err != nil {
				t.Errorf("foodUsecase.DeleteFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

