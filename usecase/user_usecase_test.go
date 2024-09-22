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

func Test_userUsecase_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)
	Validator := validator.NewUserValidator()

	type fields struct {
		ur repository.IUserRepository
		uv validator.IUserValidator
	}

	type args struct {
		email string
		user  model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.UserResponse
		wantErr bool
	}{
		{
			name: "正常系：ユーザーを取得できる",
			fields: fields{
				ur: mockRepo,
				uv: Validator,
			},
			args: args{
				email: "sample@test.com",
				user: model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Password:  "password",
				},
			},
			want: model.UserResponse{
				ID:        1,
				Username:  "test",
				Email:     "sample@test.com",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーが取得できない",
			fields: fields{
				ur: mockRepo,
				uv: Validator,
			},
			args: args{
				email: "sample@test.com",
				user:  model.User{},
			},
			want:    model.UserResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				uv: tt.fields.uv,
			}

			mockRepo.EXPECT().GetUserByEmail(gomock.Any(), tt.args.email).Do(func(user *model.User, email string) {
				*user = tt.args.user
			}).Return(nil).Times(1)

			got, err := uu.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)
	Validator := validator.NewUserValidator()

	type fields struct {
		ur repository.IUserRepository
		uv validator.IUserValidator
	}

	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.UserResponse
		wantErr bool
	}{
		{
			name: "正常系：ユーザーを作成できる",
			fields: fields{
				ur: mockRepo,
				uv: Validator,
			},
			args: args{
				user: model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Password:  "password",
				},
			},
			want: model.UserResponse{
				ID:        1,
				Username:  "test",
				Email:     "sample@test.com",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーを作成できない",
			fields: fields{
				ur: mockRepo,
				uv: Validator,
			},
			args: args{
				user: model.User{},
			},
			want:    model.UserResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				uv: tt.fields.uv,
			}

			if tt.wantErr == false {
				// ここでモックの期待値を設定する
				mockRepo.EXPECT().CreateUser(gomock.Any()).Do(func(user *model.User) {
					*user = tt.args.user
				}).Return(nil).Times(1)

				got, err := uu.CreateUser(tt.args.user)
				if (err != nil) != tt.wantErr {
					t.Errorf("userUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("userUsecase.CreateUser() = %v, want %v", got, tt.want)
				}
			} else {
				got, err := uu.CreateUser(tt.args.user)
				if (err != nil) != tt.wantErr {
					t.Errorf("userUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("userUsecase.CreateUser() = %v, want %v", got, tt.want)
				}
			}
		})
	}

}

func Test_userUsecase_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)
	Validator := validator.NewUserValidator()

	type fields struct {
		ur repository.IUserRepository
		uv validator.IUserValidator
	}

	type args struct {
		user  model.User
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.UserResponse
		wantErr bool
	}{
		{
			name: "正常系：ユーザーを更新できる",
			fields: fields{
				ur: mockRepo,
				uv: Validator,
			},
			args: args{
				user: model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Password:  "password",
				},
				email: "sample@test.com",
			},
			want: model.UserResponse{
				ID:        1,
				Username:  "test",
				Email:     "sample@test.com",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーを更新できない",
			fields: fields{
				ur: mockRepo,
				uv: Validator,
			},
			args: args{
				user:  model.User{},
				email: "sample@test.com",
			},
			want:    model.UserResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				uv: tt.fields.uv,
			}

			if tt.wantErr == false {
				// ここでモックの期待値を設定する
				mockRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Do(func(user *model.User, email string) {
					*user = tt.args.user
				}).Return(nil).Times(1)

				got, err := uu.UpdateUser(tt.args.user, tt.args.email)
				if (err != nil) != tt.wantErr {
					t.Errorf("userUsecase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("userUsecase.UpdateUser() = %v, want %v", got, tt.want)
				}
			} else {
				got, err := uu.UpdateUser(tt.args.user, tt.args.email)
				if (err != nil) != tt.wantErr {
					t.Errorf("userUsecase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("userUsecase.UpdateUser() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
func Test_userUsecase_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)
	Validator := validator.NewUserValidator()

	type fields struct {
		ur repository.IUserRepository
		uv validator.IUserValidator
	}

	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常系：ユーザーを削除できる",
			fields: fields{
				ur: mockRepo,
				uv: Validator,
			},
			args: args{
				user: model.User{
					ID:        1,
					Username:  "test",
					Email:     "sample@test.com",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Password:  "password",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				uv: tt.fields.uv,
			}

			if !tt.wantErr {
				// ユーザーを削除する前に期待されるモックの設定
				mockRepo.EXPECT().GetUserByEmail(gomock.Any(), tt.args.user.Email).Do(func(user *model.User, email string) {
					tt.args.user.Password = hashPassword(tt.args.user.Password)
					*user = tt.args.user
				}).Return(nil).Times(1)
				mockRepo.EXPECT().DeleteUser(gomock.Any()).Return(nil).Times(1)

				err := uu.DeleteUser(tt.args.user)
				if (err != nil) != tt.wantErr {
					t.Errorf("userUsecase.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				err := uu.DeleteUser(tt.args.user)
				if (err != nil) != tt.wantErr {
					t.Errorf("userUsecase.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
