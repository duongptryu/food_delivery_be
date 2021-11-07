package restaurantrepo

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

type LikeStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		likeStore: likeStore,
	}
}

func (repo *listRestaurantRepo) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListDataByCondition(ctx, nil, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	//ids := make([]int, len(result))

	//for i := range result {
	//	ids[i] = result[i].Id
	//}

	//mapResLike, err := repo.likeStore.GetRestaurantLike(ctx, ids)
	//if err != nil {
	//	log.Println("Cannot get restaurant likes: ", err)
	//}
	//
	//if v := mapResLike; v != nil {
	//	for i, item := range result {
	//		result[i].LikeCount = mapResLike[item.Id]
	//	}
	//}

	return result, nil
}
