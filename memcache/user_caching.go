package memcache

import (
	"context"
	"fmt"
	"food_delivery_be/modules/user/usermodel"
)

type RealStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type userCaching struct {
	store Caching
	realStore RealStore
}

func NewUserCaching(store Caching, realStore RealStore) *userCaching {
	return &userCaching{
		store: store,
		realStore: realStore,
	}
}

func (uc *userCaching) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	userId := condition["id"].(int)
	key := fmt.Sprintf("user-%d", userId)
	userInCache := uc.store.Read(key)
	if userInCache != nil {
		return userInCache.(*usermodel.User), nil
	}

	user, err := uc.realStore.FindUser(ctx, condition, moreInfo...)
	if err != nil {
		return nil, err
	}

	uc.store.Write(key, user)
	return user, nil
}
