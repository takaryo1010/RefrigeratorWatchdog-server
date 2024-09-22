package controller

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/usecase/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func Test_userController_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックユースケースの作成
	mockUsecase := mocks.NewMockIUserUsecase(ctrl)

	Email := "sample@test.com"

	type args struct {
		email string
	}
	tests := []struct {
		name        string
		args        args
		mockReturns model.UserResponse
		wantErr     bool
	}{
		{
			name: "正常系：メールアドレスに紐づくユーザーを取得できる",
			args: args{
				email: Email,
			},
			mockReturns: model.UserResponse{
				ID:       1,
				Username: "test",
				Email:    "sample@test.com",
			},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーが存在しない",
			args: args{
				email: Email,
			},
			mockReturns: model.UserResponse{},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().GetUserByEmail(tt.args.email).Return(tt.mockReturns, nil)

			uc := NewUserController(mockUsecase)
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/user/"+tt.args.email, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/user/:email")
			c.SetParamNames("email")
			c.SetParamValues(tt.args.email)

			if err := uc.GetUser(c); (err != nil) != tt.wantErr {
				t.Errorf("userController.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userController_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックユースケースの作成
	mockUsecase := mocks.NewMockIUserUsecase(ctrl)

	type args struct {
		user model.User
	}
	tests := []struct {
		name        string
		args        args
		mockReturns model.UserResponse
		wantErr     bool
	}{
		{
			name: "正常系：ユーザーを作成できる",
			args: args{
				user: model.User{
					Username: "test",
					Email:    "sample@test.com",
					Password: "password",
				},
			},
			mockReturns: model.UserResponse{
				ID:        1,
				Username:  "test",
				Email:     "sample@test.com",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーが作成できない",
			args: args{
				user: model.User{
					Username: "test",
					Email:    "sample@test.com",
					Password: "password",
				},
			},
			mockReturns: model.UserResponse{
				ID:        1,
				Username:  "test",
				Email:     "sample@test.com",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックの戻り値として model.UserResponse を返す
			mockUsecase.EXPECT().CreateUser(tt.args.user).Return(tt.mockReturns, nil).AnyTimes()

			uc := NewUserController(mockUsecase)
			e := echo.New()

			bodyBytes, err := json.Marshal(tt.args.user)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
			}

			body := string(bodyBytes)

			req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/user")

			if err := uc.CreateUser(c); (err != nil) != tt.wantErr {
				t.Errorf("userController.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userController_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックユースケースの作成
	mockUsecase := mocks.NewMockIUserUsecase(ctrl)

	Email := "sample@test.com"

	type args struct {
		user  model.User
		email string
	}
	tests := []struct {
		name        string
		args        args
		mockReturns model.UserResponse
		wantErr     bool
	}{
		{
			name: "正常系：ユーザーを更新できる",
			args: args{
				user: model.User{
					Username: "test",
					Email:    "sample@test.com",
					Password: "password",
				},
				email: Email,
			},
			mockReturns: model.UserResponse{
				ID:        1,
				Username:  "test",
				Email:     "sample@test.com",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーが更新できない",
			args: args{
				user: model.User{
					Username: "test",
					Email:    "sample@test.com",
					Password: "password",
				},
				email: Email,
			},
			mockReturns: model.UserResponse{
				ID:        1,
				Username:  "test",
				Email:     "sample@test.com",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().UpdateUser(tt.args.user, tt.args.email).Return(tt.mockReturns, nil).AnyTimes()

			uc := NewUserController(mockUsecase)
			e := echo.New()

			bodyBytes, err := json.Marshal(tt.args.user)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
			}

			body := string(bodyBytes)

			req := httptest.NewRequest(http.MethodPut, "/user/"+tt.args.email, strings.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/user/:email")
			c.SetParamNames("email")
			c.SetParamValues(tt.args.email)

			if err := uc.UpdateUser(c); (err != nil) != tt.wantErr {
				t.Errorf("userController.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userController_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックユースケースの作成
	mockUsecase := mocks.NewMockIUserUsecase(ctrl)

	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系：ユーザーを削除できる",
			args: args{
				user: model.User{
					Username: "test",
					Email:    "sample@test.com",
					Password: "password",
				},
			},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーが削除できない",
			args: args{
				user: model.User{
					Username: "test",
					Email:    "sample@test.com",
					Password: "password",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().DeleteUser(tt.args.user).Return(nil).AnyTimes()

			uc := NewUserController(mockUsecase)
			e := echo.New()

			bodyBytes, err := json.Marshal(tt.args.user)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
			}

			body := string(bodyBytes)

			req := httptest.NewRequest(http.MethodDelete, "/user", strings.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/user")

			if err := uc.DeleteUser(c); (err != nil) != tt.wantErr {
				t.Errorf("userController.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
