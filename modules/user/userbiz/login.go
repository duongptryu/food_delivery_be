package userbiz

import (
	"context"
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/component/tokenprovider"
	"food_delivery_be/modules/user/usermodel"
	"go.opencensus.io/trace"
)

type LoginStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStore, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	_, span1 := trace.StartSpan(ctx, "user.biz.find_user")
	span1.AddAttributes(trace.StringAttribute("username", data.Email))
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	span1.End()
	if err != nil {
		return nil, usermodel.ErUsernameOrPasswordInvalid
	}
	_, span2 := trace.StartSpan(ctx, "user.biz.check_password")
	passHashed := biz.hasher.Hash(data.Password + user.Salt)
	span2.End()
	if user.Password != passHashed {
		return nil, usermodel.ErUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	_, span3 := trace.StartSpan(ctx, "user.biz.generate_token")
	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	span3.End()
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
