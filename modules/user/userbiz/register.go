package userbiz

import (
	"context"
	"food_delivery_be/common"
	"food_delivery_be/modules/user/usermodel"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}
type registerBiz struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterStorage(registerStorage RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}
func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		if err != common.RecordNotFound {
			return common.ErrDB(err)
		}
	}

	if user.Id != 0 {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code
	data.Status = 1

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
