package subscriber

import (
	"context"
	"food_delivery_be/component"
	"food_delivery_be/modules/restaurant/restaurantstorage"
	"food_delivery_be/pubsub"
)

func RunDecreaseLikeCountAfterUserUnlikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
