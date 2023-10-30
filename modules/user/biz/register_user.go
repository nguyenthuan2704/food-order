package biz

import (
	"context"
	"food-client/common"
	"food-client/modules/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreation) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *model.UserCreation) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return model.ErrEmailExisted
	}

	data.Password = business.hasher.Hash(data.Password)
	data.Role = "user" // hard code

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
