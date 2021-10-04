package restaurantbiz

import (
	"context"
	"food_delivery_be/common"
	"food_delivery_be/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		condition map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)

	return result, err
}
