package repositories

import (
	"delivery-app/src/domain/entities"
	"delivery-app/src/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	*database.MongoDB
}
type IUserRepository interface {
	GetUsers() (interface{}, error)
	CreateUser(data *entities.User) (interface{}, error)
	GetUser(id string) (interface{}, error)
	GetUserByParam(params interface{}) (interface{}, error)
	UpdateUser(id string, data *entities.User) (interface{}, error)
	DeleteUser(id string) (bool, error)
}

func (repo *UserRepository) GetUsers() (interface{}, error) {
	users, err := repo.GetAll("users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) CreateUser(data *entities.User) (interface{}, error) {
	user, err := repo.Save("users", data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUser(id string) (interface{}, error) {
	parser, _ := primitive.ObjectIDFromHex(id)
	user, err := repo.GetByID("users", parser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserByParam(params interface{}) (interface{}, error) {
	user, err := repo.GetParam("users", params)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) UpdateUser(id string, data *entities.User) (interface{}, error) {
	parser, _ := primitive.ObjectIDFromHex(id)
	user, err := repo.Update("users", parser, data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) DeleteUser (id string) (bool, error) {
	parser, _ := primitive.ObjectIDFromHex(id)
	user, err := repo.Delete("users", parser)
	if err != nil {
		return false, err
	}
	return user, nil
}