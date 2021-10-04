package restaurantbiz

import (
	"context"
	"errors"
	"food_delivery_be/common"
	"food_delivery_be/modules/restaurant/restaurantmodel"
)

type DeleteRestaurant interface {
	FindDataByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)

	SoftDeleteData(ctx context.Context, id int) error
}

type deleteRestaurant struct {
	Store DeleteRestaurant
}

func NewDeleteRestaurantBiz(store DeleteRestaurant) *deleteRestaurant {
	return &deleteRestaurant{store}
}

func (biz *deleteRestaurant) DeleteRestaurantBiz(ctx context.Context, id int) error {
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

	if err := biz.Store.SoftDeleteData(ctx, id); err != nil {
		return err
	}
	return nil
}
