package ginrestaurant

import (
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/modules/restaurant/restaurantbiz"
	"food_delivery_be/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		//id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurantBiz(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
