package userauth

import (
	"testing"

	"github.com/masraga/meraki/models"
	"github.com/masraga/meraki/pkg"
	"github.com/masraga/meraki/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LoadConfig() *pkg.Autoload {
	config := pkg.NewConfig("../../.env") //get app.env path
	return &pkg.Autoload{Config: config}
}

func LoadUserRepo() *repositories.User {
	var model *models.User
	autoload := LoadConfig()
	db := autoload.Database()
	repo := autoload.MongoRepository("User")

	return &repositories.User{
		Db:    db,
		Repo:  repo,
		Model: model,
	}
}

func TestLenPass(t *testing.T) {
	req := &Request{
		Password: "123456",
	}

	register := Register{
		MinPassLen: 10,
		Request:    req,
	}

	if !register.IsLessPassword() {
		t.Errorf(`%s must be less password, minPass %d`, req.Password, register.MinPassLen)
	}
}

func TestMismatchPass(t *testing.T) {
	req := &Request{
		Password:   "12345678",
		RePassword: "123456890",
	}

	register := Register{
		MinPassLen: 5,
		Request:    req,
	}

	if register.IsSamePass() {
		t.Errorf(`%s is mismatch with %s`, req.Password, req.RePassword)
	}
}

func TestAddUser(t *testing.T) {
	req := &Request{
		Username:   "TestKoderpedia",
		Name:       "TestKoderpedia",
		Password:   "12345678",
		RePassword: "12345678",
	}

	userRepo := LoadUserRepo()
	register := Register{
		UserRepo: userRepo,
		Request:  req,
	}

	user, err := register.AddUser()
	if err != nil {
		t.Error("fail save to db", err)
	}

	userRepo.DeleteByID(user.InsertedID.(primitive.ObjectID).Hex())
}
