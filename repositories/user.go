package repositories

import(
	"time"

	"github.com/masraga/meraki/models"
	"github.com/masraga/meraki/pkg"
	driver "github.com/masraga/meraki/pkg/driver"
	app "github.com/masraga/meraki/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Db *driver.MongodbDriver
	Repo *app.MongoRepository
	Model *models.User
}

func (r *User) Create(request models.User) (*mongo.InsertOneResult, error) {
	request.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	request.IsDeleted = false 

	query, err := r.Repo.InsertOne(request)
	if err != nil {
		return nil, err
	}

	return query, err
}

func (r *User) FindById(id string) (*models.User, error) {
	err := r.Repo.FindById(id).Decode(&r.Model)
	if err != nil {
		return nil, err
	}

	return r.Model, nil
}

func (r *User) UpdateByID(id string, fieldSet *models.User) (*mongo.UpdateResult, error) {
	fieldSet.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	updateResult, err := r.Repo.UpdateByID(id, fieldSet)
	return updateResult, err
}

func (r *User) DeleteByID(id string) (*mongo.DeleteResult, error) {
	result, err := r.Repo.DeleteByID(id)
	return result, err
}

func (r *User) SoftDeleteByID(id string) (*mongo.UpdateResult, error) {
	r.Model.IsDeleted = true
	r.Model.DeletedAt = primitive.NewDateTimeFromTime(time.Now())
	deleteResult, err := r.Repo.UpdateByID(id, r.Model)
	return deleteResult, err
}

func (r *User) FindOne(filter bson.D) (*models.User, error) {
	result, err := r.Repo.FindOne(filter)
	if err != nil {
		return nil, err
	}

	result.Decode(&r.Model)
	return r.Model, err
}

func (r *User) Aggregate(filter bson.D) (*mongo.Cursor, error) {
	cursor, err := r.Repo.Aggregate(filter)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func NewUser() *User {
	var model *models.User
	autoload := pkg.NewAutoload()
	db := autoload.Database()
	repo := autoload.MongoRepository("User")
	return &User{
		Db: db,
		Repo: repo,
		Model: model,
	}
}

