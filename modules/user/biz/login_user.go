package biz

import (
	"context"
	"food-client/common"
	"food-client/component"
	"food-client/component/tokenprovider"
	"food-client/modules/user/model"
	"go.opencensus.io/trace"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

func (business *loginBusiness) Login(ctx context.Context, data *model.UserLogin) (*tokenprovider.Token, error) {
	ctx1, span1 := trace.StartSpan(ctx, "user.biz.login")

	user, err := business.storeUser.FindUser(ctx1, map[string]interface{}{"email": data.Email})

	span1.End()

	if err != nil {
		return nil, model.ErrUsernameOrPasswordInvalid
	}

	_, span2 := trace.StartSpan(ctx, "user.biz.login.gen-jwt")
	passHashed := business.hasher.Hash(data.Password)

	if user.Password != passHashed {
		span2.End()
		return nil, model.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	span2.End()

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
