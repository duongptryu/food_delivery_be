package restaurantbiz

import (
	"context"
	"food_delivery_be/common"
	"food_delivery_be/modules/restaurant/restaurantmodel"
)

type GetRestaurant interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurant struct {
	Store GetRestaurant
}

func NewGetRestaurantBiz(store GetRestaurant) *getRestaurant {
	return &getRestaurant{
		Store: store,
	}
}

func (biz *getRestaurant) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.Store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}
	return data, err
}
