package repository

import (
	"RefrigeratorWatchdog-server/mocks"
	"RefrigeratorWatchdog-server/model"
	"errors"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func Test_foodRepository_GetFoodsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリの作成
	mockRepo := mocks.NewMockIFoodRepository(ctrl)

	UserID := uint(1)

	type args struct {
		foods  *[]model.Food
		userID uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常系：ユーザーIDに紐づく食材を取得できる",
			args: args{
				foods: &[]model.Food{
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
				userID: UserID,
			},
			wantErr: false,
		},
		{name: "異常系：ユーザーIDに紐づく食材が取得できない",
			args: args{
				foods:  &[]model.Food{},
				userID: 0,
			},
			wantErr: true,
		},
		{name: "異常系：引数のfoodsがnilの場合",
			args: args{
				foods:  nil,
				userID: UserID,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantErr {
				mockRepo.EXPECT().GetFoodsByUserID(tt.args.foods, tt.args.userID).Return(errors.New("error"))

			} else {
				mockRepo.EXPECT().GetFoodsByUserID(tt.args.foods, tt.args.userID).Return(nil)
			}

			if err := mockRepo.GetFoodsByUserID(tt.args.foods, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("foodRepository.GetFoodsByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_foodRepository_CreateFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリの作成
	mockRepo := mocks.NewMockIFoodRepository(ctrl)

	type args struct {
		food *model.Food
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常系：食材を作成できる",
			args: args{
				food: &model.Food{
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
			wantErr: false,
		},
		{name: "異常系：引数のfoodがnilの場合",
			args: args{
				food: nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantErr {
				mockRepo.EXPECT().CreateFood(tt.args.food).Return(errors.New("error"))

			} else {
				mockRepo.EXPECT().CreateFood(tt.args.food).Return(nil)
			}

			if err := mockRepo.CreateFood(tt.args.food); (err != nil) != tt.wantErr {
				t.Errorf("foodRepository.CreateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_foodRepository_UpdateFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリの作成
	mockRepo := mocks.NewMockIFoodRepository(ctrl)

	type args struct {
		food *model.Food
		id   uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常系：食材を更新できる",
			args: args{
				food: &model.Food{
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
			wantErr: false,
		},
		{name: "異常系：引数のfoodがnilの場合",
			args: args{
				food: nil,
				id:   1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantErr {
				mockRepo.EXPECT().UpdateFood(tt.args.food, tt.args.id).Return(errors.New("error"))

			} else {
				mockRepo.EXPECT().UpdateFood(tt.args.food, tt.args.id).Return(nil)
			}

			if err := mockRepo.UpdateFood(tt.args.food, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("foodRepository.UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_foodRepository_DeleteFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリの作成
	mockRepo := mocks.NewMockIFoodRepository(ctrl)

	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常系：食材を削除できる",
			args: args{
				id: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantErr {
				mockRepo.EXPECT().DeleteFood(tt.args.id).Return(errors.New("error"))

			} else {
				mockRepo.EXPECT().DeleteFood(tt.args.id).Return(nil)
			}

			if err := mockRepo.DeleteFood(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("foodRepository.DeleteFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
