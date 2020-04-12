package services

import (
	"delivery-app/src/domain/entities"
	"delivery-app/src/domain/repositories"
	"delivery-app/src/infrastructure/utils"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	*entities.User
	repositories.UserRepository
}

type IUserService interface {
	Saved([]byte) (interface{}, error)
	Updated([]byte) (interface{}, error)
}

func (svc UserService) Saved(data []byte) (interface{}, error) {
	var user *entities.User
	json.Unmarshal(data, &user)
	password, _ := utils.GenerateHash(user.Password)
	slug := utils.RenderSLUG(user.Name)

	user.Password = password
	user.Slug = slug
	user.ID = primitive.NewObjectID()

	request, err := svc.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (svc UserService) Updated(id string, data []byte) (interface{}, error) {
	_, err := svc.GetUser(id)
	if err != nil {
		return nil, err
	}
	var user *entities.User
	json.Unmarshal(data, &user)

	slug := utils.RenderSLUG(user.Name)
	user.Slug = slug

	request, err := svc.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	return request, nil
}