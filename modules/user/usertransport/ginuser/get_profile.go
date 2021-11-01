package ginuser

import (
	"food_delivery_be/common"
	"food_delivery_be/component"
	"github.com/gin-gonic/gin"
)

func GetProfile(app component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(200, common.SimpleSuccessResponse(data))
	}
}
