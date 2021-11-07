package rstlikebiz

import (
	"context"
	"food_delivery_be/common"
	restaurantlikemodel "food_delivery_be/modules/restaurantlike/model"
	"food_delivery_be/pubsub"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

//
//type DecreaseLikeCountStore interface {
//	DecreaseLikeCount(ctx context.Context, id int) error
//}

type userUnLikeRestaurantBiz struct {
	store UserUnLikeRestaurantStore
	//decStore DecreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore, pubsub pubsub.Pubsub) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{
		store:  store,
		pubsub: pubsub,
	}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrUserCannotUnLikeRestaurant(err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{
		restaurantId,
		userId,
		nil,
		nil,
	}))

	//go func() {
	//	defer common.AppRecovery()
	//
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	//	})
	//
	//	job.SetRetryDurations([]time.Duration{time.Second})
	//
	//	_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	return nil
}
