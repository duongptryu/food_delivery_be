package rstlikebiz

import (
	"context"
	"food_delivery_be/common"
	restaurantlikemodel "food_delivery_be/modules/restaurantlike/model"
	"food_delivery_be/pubsub"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

//type IncreaseLikeCountStore interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	//incStore IncreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	//incStore IncreaseLikeCountStore,
	pubsub pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		//incStore: incStore,
		pubsub: pubsub,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrUserCannotLikeRestaurant(err)
	}

	//side effect
	//new solution - use pubsub
	biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	//old solution
	//go func() {
	//	defer common.AppRecovery()
	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	return nil
}
