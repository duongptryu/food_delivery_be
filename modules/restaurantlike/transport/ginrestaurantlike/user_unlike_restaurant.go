package ginrestaurantlike

import (
	"food_delivery_be/common"
	"food_delivery_be/component"
	rstlikebiz "food_delivery_be/modules/restaurantlike/biz"
	restaurantlikestorage "food_delivery_be/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

//PUT /v1/restaurant/:id/unlike

func UserUnLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		//decStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserUnLikeRestaurantBiz(store, appCtx.GetPubsub())

		err = biz.UnLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
