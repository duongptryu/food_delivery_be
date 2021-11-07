package subscriber

import (
	"context"
	"food_delivery_be/component"
	"food_delivery_be/modules/restaurant/restaurantstorage"
	"food_delivery_be/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

// run with setup without lib
//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//
//	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
//
//	go func() {
//		defer common.AppRecovery()
//		msg := <-c
//		likeData := msg.Data().(HasRestaurantId)
//
//		_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//	}()
//}

func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user liked restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
