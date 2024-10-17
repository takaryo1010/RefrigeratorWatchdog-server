package usecase

import (
	"RefrigeratorWatchdog-server/model"

	"RefrigeratorWatchdog-server/repository"
	"RefrigeratorWatchdog-server/validator"

	"golang.org/x/crypto/bcrypt"
	"errors"
)

type IUserUsecase interface {
	GetUserByEmail(email string) (model.UserResponse, error)
	CreateUser(user model.User) (model.UserResponse, error)
	UpdateUser(user model.User, email string) (model.UserResponse, error)
	DeleteUser(user model.User) error
	LoginUser(user model.User,decodedEmail string) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) GetUserByEmail(email string) (model.UserResponse, error) {
	user := model.User{}
	if err := uu.ur.GetUserByEmail(&user, email); err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (uu *userUsecase) CreateUser(user model.User) (model.UserResponse, error) {
	if err := uu.uv.ValidateUser(user); err != nil {
		return model.UserResponse{}, err
	}
	user.Password = hashPassword(user.Password)
	if err := uu.ur.CreateUser(&user); err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (uu *userUsecase) UpdateUser(user model.User, email string) (model.UserResponse, error) {
	// ユーザー情報の検証
	if err := uu.uv.ValidateUser(user); err != nil {
		return model.UserResponse{}, err
	}

	// パスワードが更新されている場合はハッシュ化
	if user.Password != "" {
		user.Password = hashPassword(user.Password)
	}
	// ユーザー情報の更新
	if err := uu.ur.UpdateUser(&user, email); err != nil {
		return model.UserResponse{}, err
	}

	// 更新されたユーザー情報を返す
	return model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (uu *userUsecase) DeleteUser(user model.User) error {
	getuser := model.User{}
	if err := uu.ur.GetUserByEmail(&getuser, user.Email); err != nil {
		return err
	}
	if !comparePassword(getuser.Password, user.Password) {
		return model.ErrInvalidPassword
	}

	if err := uu.ur.DeleteUser(&user); err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) LoginUser(user model.User,decodedEmail string) (model.UserResponse, error) {
	getuser := model.User{}
	if err := uu.ur.GetUserByEmail(&getuser, decodedEmail); err != nil {
		return model.UserResponse{}, err
	}
	if !comparePassword(getuser.Password, user.Password) {
		return model.UserResponse{}, model.ErrInvalidPassword
	}
	if user.Username != getuser.Username {
		return model.UserResponse{}, errors.New("invalid username")
	}
	return model.UserResponse{
		ID:        getuser.ID,
		Username:  getuser.Username,
		Email:     getuser.Email,
		CreatedAt: getuser.CreatedAt,
	}, nil
}
