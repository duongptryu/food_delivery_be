package ginuser

import (
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/component/hasher"
	"food_delivery_be/component/tokenprovider/jwt"
	"food_delivery_be/modules/user/userbiz"
	"food_delivery_be/modules/user/usermodel"
	"food_delivery_be/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(account))
	}
}
