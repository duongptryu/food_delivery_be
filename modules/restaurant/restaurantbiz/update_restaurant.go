package restaurantbiz

import (
	"context"
	"errors"
	"food_delivery_be/common"
	"food_delivery_be/modules/restaurant/restaurantmodel"
)

type UpdateRestaurant interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)

	UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurant struct {
	Store UpdateRestaurant
}

func NewUpdateRestaurantBiz(store UpdateRestaurant) *updateRestaurant {
	return &updateRestaurant{store}
}

func (biz *updateRestaurant) UpdateRestaurantBiz(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := biz.Store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return common.RecordNotFound
		}
		return err
	}
	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	if err := biz.Store.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
