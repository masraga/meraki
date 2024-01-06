package userauth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/models"
	"github.com/masraga/meraki/pkg"
	"github.com/masraga/meraki/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	ID   string `json:"userId"`
	Name string `json:"name"`
}

type LoginResponse struct {
	AccessToken string    `json:"accessToken"`
	UserInfo    *UserInfo `json:"userInfo"`
}

type LoginRequset struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Ctx                 *gin.Context
	UserRepo            *repositories.User
	LoginRequest        *LoginRequset
	Autoload            *pkg.Autoload
	ExpiredTokenMinutes float64
}

func (l *Login) SetApiError(statusCode int, message error) {
	l.Ctx.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": fmt.Sprint(message),
	})
}

func (l *Login) GetUser() (*models.User, error) {
	userCursor, err := l.UserRepo.FindOne(bson.D{
		primitive.E{Key: "username", Value: l.LoginRequest.Username},
	})
	if err != nil {
		return nil, err
	}

	return userCursor, nil
}

func (l *Login) DecodePassword(reqPass string, dbPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(reqPass))
	if err != nil {
		return fmt.Errorf("password not match")
	}
	return nil
}

func (l *Login) GenerateToken(userModel *models.User) (string, error) {
	jwt := l.Autoload.JwtHelper()
	token, err := jwt.GenerateToken(time.Now().Add(time.Duration(l.ExpiredTokenMinutes)*time.Minute), userModel.ID.Hex())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (l *Login) UpdateUserToken(userModel *models.User, token string) (*models.User, error) {
	userModel.AccessToken = token
	_, err := l.UserRepo.UpdateByID(userModel.ID.Hex(), userModel)

	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func (l *Login) Run() error {
	fmt.Println("[log] get user data")
	user, err := l.GetUser()
	if err != nil {
		l.SetApiError(http.StatusPreconditionFailed, err)
		return err
	}

	fmt.Println("[log] decode user password")
	err = l.DecodePassword(l.LoginRequest.Password, user.Password)
	if err != nil {
		l.SetApiError(http.StatusPreconditionFailed, err)
		return err
	}

	fmt.Println("[log] generate user access token")
	token, err := l.GenerateToken(user)
	if err != nil {
		l.SetApiError(http.StatusPreconditionFailed, err)
	}

	fmt.Println("[log] save user access token")
	updateUserCursor, err := l.UpdateUserToken(user, token)
	if err != nil {
		l.SetApiError(http.StatusPreconditionFailed, err)
	}

	fmt.Println("[log] set response")
	response := LoginResponse{
		AccessToken: updateUserCursor.AccessToken,
		UserInfo: &UserInfo{
			ID:   updateUserCursor.ID.Hex(),
			Name: updateUserCursor.Name,
		},
	}

	l.Ctx.JSON(http.StatusOK, response)

	return nil
}

func NewLogin(ctx *gin.Context) *Login {
	autoload := pkg.NewAutoload()
	userRepo := repositories.NewUser()
	var loginRequest *LoginRequset
	ctx.ShouldBindJSON(&loginRequest)
	return &Login{
		Ctx:                 ctx,
		UserRepo:            userRepo,
		LoginRequest:        loginRequest,
		Autoload:            autoload,
		ExpiredTokenMinutes: 60,
	}
}
