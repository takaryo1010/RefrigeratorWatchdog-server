package controller

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/usecase"
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IUserController interface {
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	LoginUser(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// GetUser godoc
// @Summary Get user by email
// @Description Get user by email
// @ID get-user
// @Accept  json
// @Produce  json
// @Param email path string true "Email"
// @Success 200 {object} model.UserResponse
// @Router /users/{email} [get]
// @Tags users
func (uc *userController) GetUser(c echo.Context) error {
	// URLパラメータからメールアドレスを取得
	email := c.Param("email")

	// メールアドレスをデコード
	decodedEmail, err := url.QueryUnescape(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid email format"})
	}

	// デコードしたメールアドレスを使ってユーザーを取得
	user, err := uc.uu.GetUserByEmail(decodedEmail)
	if err != nil {
		// ユーザーが見つからない場合、404エラーを返すことも考慮
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create user
// @Description Create user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param user body model.UserRequest true "User"
// @Success 200 {object} model.UserResponse
// @Router /users [post]
// @Tags users
func (uc *userController) CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdUser, err := uc.uu.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, createdUser)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param email path string true "Email"
// @Param user body model.UserRequest true "User"
// @Success 200 {object} model.UserResponse
// @Router /users/{email} [put]
// @Tags users
func (uc *userController) UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	email := c.Param("email")

	// メールアドレスをデコード
	decodedEmail, err := url.QueryUnescape(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid email format"})
	}

	// デコードしたメールアドレスを使用してユーザーを更新
	updatedUser, err := uc.uu.UpdateUser(user, decodedEmail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param user body model.UserRequest true "User"
// @Success 200 {string} string "deleted"
// @Router /users [delete]
// @Tags users
func (uc *userController) DeleteUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := uc.uu.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "User deleted successfully")
}



// LoginUser godoc
// @Summary Login user
// @Description Login user
// @ID login-user
// @Accept  json
// @Produce  json
// @Param user body model.UserRequest true "User"
// @Success 200 {object} model.UserResponse
// @Router /users/login [post]
// @Tags users
func (uc *userController) LoginUser(c echo.Context)error{
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// メールアドレスをデコード
	decodedEmail, err := url.QueryUnescape(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid email format"})
	}

	response ,err := uc.uu.LoginUser(user, decodedEmail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}