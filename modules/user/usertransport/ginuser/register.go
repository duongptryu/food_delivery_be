package ginuser

import (
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/component/hasher"
	"food_delivery_be/modules/user/userbiz"
	"food_delivery_be/modules/user/usermodel"
	"food_delivery_be/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterStorage(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(200, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
