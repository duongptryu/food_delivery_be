package subscriber

import (
	"context"
	"food_delivery_be/component"
	"food_delivery_be/modules/restaurant/restaurantstorage"
	"food_delivery_be/pubsub"
	"food_delivery_be/skio"
)

func RunDecreaseLikeCountAfterUserUnlikeRestaurant(appCtx component.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			//_ = rtEngine.EmitToUser(likeData.GetOwnerId(), string(message.Channel()), likeData)

			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
