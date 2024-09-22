package controller

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/usecase/mocks" 
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func Test_foodController_GetFoodsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックユースケースの作成
	mockUsecase := mocks.NewMockIFoodUsecase(ctrl)

	UserID := uint(1)

	type args struct {
		userID uint
	}
	tests := []struct {
		name        string
		args        args
		mockReturns []model.FoodResponse
		wantErr     bool
	}{
		{
			name: "正常系：ユーザーIDに紐づく食材を取得できる",
			args: args{
				userID: UserID,
			},
			mockReturns: []model.FoodResponse{
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
			args: args{
				userID: 0,
			},
			mockReturns: []model.FoodResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モック設定：GetFoodsByUserID の引数 userID に対して、mockReturns を返す
			mockUsecase.EXPECT().GetFoodsByUserID(tt.args.userID).Return(tt.mockReturns, nil)

			fc := NewFoodController(mockUsecase)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/foods/"+strconv.Itoa(int(tt.args.userID)), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/foods/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(int(tt.args.userID)))

			// メソッド呼び出し
			if err := fc.GetFoodsByUserID(c); (err != nil) != tt.wantErr {
				t.Errorf("foodController.GetFoodsByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_foodController_CreateFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックユースケースの作成
	mockUsecase := mocks.NewMockIFoodUsecase(ctrl)

	type args struct {
		food model.Food
	}
	tests := []struct {
		name        string
		args        args
		mockReturns model.FoodResponse
		wantErr     bool
	}{
		{
			name: "正常系：食材を作成できる",
			args: args{
				food: model.Food{
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
			mockReturns: model.FoodResponse{
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
			name: "異常系：食材を作成できない",
			args: args{
				food: model.Food{
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
			mockReturns: model.FoodResponse{},
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モック設定：CreateFood の引数 food に対して、mockReturns を返す
			mockUsecase.EXPECT().CreateFood(tt.args.food).Return(tt.mockReturns, nil)

			fc := NewFoodController(mockUsecase)
			e := echo.New()

			// 正しいリクエストボディを含める
			bodyBytes, err := json.Marshal(tt.args.food)
			if err != nil {
				t.Fatalf("failed to marshal food: %v", err)
			}
			body := string(bodyBytes)

			req := httptest.NewRequest(http.MethodPost, "/foods", strings.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/foods")

			if err := fc.CreateFood(c); (err != nil) != tt.wantErr {
				t.Errorf("foodController.CreateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


func Test_foodController_UpdateFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックユースケースの作成
	mockUsecase := mocks.NewMockIFoodUsecase(ctrl)

	type args struct {
		food model.Food
		id   uint
	}
	tests := []struct {
		name        string
		args        args
		mockReturns model.FoodResponse
		wantErr     bool
	}{
		{
			name: "正常系：食材を更新できる",
			args: args{
				food: model.Food{
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
			mockReturns: model.FoodResponse{
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
			name: "異常系：食材を更新できない",
			args: args{
				food: model.Food{
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
			mockReturns: model.FoodResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モック設定：UpdateFood の引数 food に対して、mockReturns を返す
			mockUsecase.EXPECT().UpdateFood(tt.args.food, tt.args.id).Return(tt.mockReturns, nil)

			fc := NewFoodController(mockUsecase)
			e := echo.New()

			// 正しいリクエストボディを含める
			bodyBytes, err := json.Marshal(tt.args.food)
			if err != nil {
				t.Fatalf("failed to marshal food: %v", err)
			}
			body := string(bodyBytes)

			req := httptest.NewRequest(http.MethodPut, "/foods/"+strconv.Itoa(int(tt.args.id)), strings.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/foods/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(int(tt.args.id)))

			if err := fc.UpdateFood(c); (err != nil) != tt.wantErr {
				t.Errorf("foodController.UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_foodController_DeleteFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックユースケースの作成
	mockUsecase := mocks.NewMockIFoodUsecase(ctrl)

	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系：食材を削除できる",
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "異常系：食材を削除できない",
			args: args{
				id: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モック設定：DeleteFood の引数 id に対して、nil を返す
			mockUsecase.EXPECT().DeleteFood(tt.args.id).Return(nil)

			fc := NewFoodController(mockUsecase)
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/foods/"+strconv.Itoa(int(tt.args.id)), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/foods/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(int(tt.args.id)))

			if err := fc.DeleteFood(c); (err != nil) != tt.wantErr {
				t.Errorf("foodController.DeleteFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
