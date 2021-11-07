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

//POST /v1/restaurant/:id/like

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		//incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubsub())

		err = biz.LikeRestaurant(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
