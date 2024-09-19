package repository

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/repository/mocks"
	"errors"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func Test_userRepository_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)

	email := "sample@test.com"

	type args struct {
		user  *model.User
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常系：メールアドレスに紐づくユーザーを取得できる",
			args: args{
				user: &model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					Password:  "password",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				email: email,
			},
			wantErr: false,
		},
		{name: "異常系：メールアドレスに紐づくユーザーが存在しない",
			args: args{
				user: &model.User{
					ID:        1,
					Username:  "test",
					Email:     "notsample@test.com",
					Password:  "password",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				email: email,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				mockRepo.EXPECT().GetUserByEmail(tt.args.user, tt.args.email).Return(errors.New("object does not exist"))
			} else {
				mockRepo.EXPECT().GetUserByEmail(tt.args.user, tt.args.email).Return(nil)
			}
			if err := mockRepo.GetUserByEmail(tt.args.user, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)

	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常系：ユーザーを作成できる",
			args: args{
				user: &model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					Password:  "password",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				mockRepo.EXPECT().CreateUser(tt.args.user).Return(errors.New("error"))
			} else {
				mockRepo.EXPECT().CreateUser(tt.args.user).Return(nil)
			}
			if err := mockRepo.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)

	email := "sample@test.com"

	type args struct {
		user  *model.User
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常系：ユーザーを更新できる",

			args: args{
				user: &model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					Password:  "password",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				email: email,
			},
			wantErr: false,
		},
		{name: "異常系：ユーザーが存在しない",

			args: args{
				user: &model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					Password:  "password",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				email: email,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				mockRepo.EXPECT().UpdateUser(tt.args.user, tt.args.email).Return(errors.New("object does not exist"))
			} else {
				mockRepo.EXPECT().UpdateUser(tt.args.user, tt.args.email).Return(nil)
			}
			if err := mockRepo.UpdateUser(tt.args.user, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)

	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常系：ユーザーを削除できる",
			args: args{
				user: &model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					Password:  "password",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{name: "異常系：ユーザーが存在しない",
			args: args{
				user: &model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					Password:  "password",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				mockRepo.EXPECT().DeleteUser(tt.args.user).Return(errors.New("object does not exist"))
			} else {
				mockRepo.EXPECT().DeleteUser(tt.args.user).Return(nil)
			}
			if err := mockRepo.DeleteUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
