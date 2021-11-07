package ginrestaurant

import (
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/modules/restaurant/restaurantbiz"
	"food_delivery_be/modules/restaurant/restaurantmodel"
	"food_delivery_be/modules/restaurant/restaurantrepo"
	"food_delivery_be/modules/restaurant/restaurantstorage"
	restaurantlikestorage "food_delivery_be/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		likeStore := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
