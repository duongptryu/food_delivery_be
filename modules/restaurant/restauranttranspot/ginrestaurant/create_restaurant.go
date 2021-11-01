package ginrestaurant

import (
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/modules/restaurant/restaurantbiz"
	"food_delivery_be/modules/restaurant/restaurantmodel"
	"food_delivery_be/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrCannotCreateEntity(restaurantmodel.EntityName, err))
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(restaurantmodel.EntityName, err))
			return
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.NewSuccessResponse(data.FakeId.String(), nil, nil))
	}
}
