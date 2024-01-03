package userauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/models"
	"github.com/masraga/meraki/repositories"
	"golang.org/x/crypto/bcrypt"
)

type RequestRegister struct {
	Username   string `bson:"username" json:"username"`
	Password   string `bson:"password" json:"password"`
	RePassword string `bson:"re-password" json:"re-password"`
	Name       string `bson:"name" json:"name"`
}

type Register struct {
	Ctx             *gin.Context
	RequestRegister *RequestRegister
	MinPassLen      int
	UserRepo        *repositories.User
}

func (r *Register) CheckNilRequest() error {
	if r.RequestRegister == nil {
		return fmt.Errorf("not found request")
	}

	return nil
}

func (r *Register) CheckPassLen() error {
	if len(r.RequestRegister.Password) < r.MinPassLen {
		return fmt.Errorf(fmt.Sprintf("min password length is %d", r.MinPassLen))
	}
	return nil
}

func (r *Register) CheckMisMatchPass() error {
	if r.RequestRegister.Password != r.RequestRegister.RePassword {
		return fmt.Errorf("password & retype password is not match")
	}
	return nil
}

func (r *Register) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(r.RequestRegister.Password), 14)
	return string(bytes), err
}

func (r *Register) SetApiError(statusCode int, message error) {
	r.Ctx.JSON(http.StatusPreconditionFailed, gin.H{
		"status":  statusCode,
		"message": fmt.Sprint(message),
	})
}

func (r *Register) Save() error {
	err := r.CheckNilRequest()
	if err != nil {
		r.SetApiError(http.StatusPreconditionFailed, err)
		return err
	}

	err = r.CheckPassLen()
	if err != nil {
		r.SetApiError(http.StatusPreconditionFailed, err)
		return err
	}

	err = r.CheckMisMatchPass()
	if err != nil {
		r.SetApiError(http.StatusPreconditionFailed, err)
		return err
	}

	password, _ := r.HashPassword()
	r.RequestRegister.Password = password
	r.UserRepo.Create(models.User{
		Username: r.RequestRegister.Username,
		Name:     r.RequestRegister.Name,
		Password: r.RequestRegister.Password,
	})

	r.Ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "new user created successfully",
	})
	return nil

}

func NewRegister(ctx *gin.Context) *Register {
	var reqRegister *RequestRegister
	ctx.ShouldBindJSON(&reqRegister)
	userRepo := repositories.NewUser()
	return &Register{
		Ctx:             ctx,
		RequestRegister: reqRegister,
		MinPassLen:      8,
		UserRepo:        userRepo,
	}
}
