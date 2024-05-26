package userauth

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/masraga/meraki/models"
	"github.com/masraga/meraki/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Request struct {
	Username     string `bson:"username" json:"username"`
	Password     string `bson:"password" json:"password"`
	RePassword   string `bson:"re-password" json:"re-password"`
	Name         string `bson:"name" json:"name"`
	HashPassword string
}

type Register struct {
	Ctx        *fiber.Ctx
	Request    *Request
	MinPassLen int16
	UserRepo   *repositories.User
}

func (r *Register) SetApiErr(statusCode int, message string) {
	r.Ctx.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode,
		"message": message,
	})
}

func (r *Register) IsReq() bool {
	return r.Request != nil
}

func (r *Register) IsLessPassword() bool {
	return len(r.Request.Password) < int(r.MinPassLen)
}

func (r *Register) IsSamePass() bool {
	return r.Request.Password == r.Request.RePassword
}

func (r *Register) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(r.Request.Password), 14)
	return string(bytes), err
}

func (r *Register) AddUser() (*mongo.InsertOneResult, error) {
	return r.UserRepo.Create(models.User{
		Username: r.Request.Username,
		Password: r.Request.HashPassword,
		Name:     r.Request.Name,
	})
}

func (r *Register) Save() {
	if !r.IsReq() {
		r.SetApiErr(http.StatusPreconditionFailed, "request not found")
		return
	}
	if r.IsLessPassword() {
		r.SetApiErr(http.StatusPreconditionFailed, fmt.Sprintf("minimun password length is %d", r.MinPassLen))
		return
	}
	if !r.IsSamePass() {
		r.SetApiErr(http.StatusPreconditionFailed, "password is mismatch")
		return
	}

	hashPassword, err := r.HashPassword()
	if err != nil {
		r.SetApiErr(http.StatusInternalServerError, "error hashing password")
		return
	}
	r.Request.HashPassword = hashPassword

	_, err = r.AddUser()
	if err != nil {
		r.SetApiErr(http.StatusInternalServerError, "error save data to db")
	}

	r.Ctx.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "new user created successfully",
	})
}

func NewRegister(ctx *fiber.Ctx) *Register {
	var (
		request *Request
	)

	ctx.BodyParser(&request)
	userRepo := repositories.NewUser()

	return &Register{
		Ctx:        ctx,
		Request:    request,
		MinPassLen: 8,
		UserRepo:   userRepo,
	}
}
