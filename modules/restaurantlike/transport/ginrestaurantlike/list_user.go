package ginrestaurantlike

import (
	"food_delivery_be/common"
	"food_delivery_be/component"
	rstlikebiz "food_delivery_be/modules/restaurantlike/biz"
	restaurantlikemodel "food_delivery_be/modules/restaurantlike/model"
	restaurantlikestorage "food_delivery_be/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GET /v1/restaurant/:id/liked-user

func ListUserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var filter restaurantlikemodel.Filter
		//

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewListUserLikeRestaurant(store)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
