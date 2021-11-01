package main

import (
	"fmt"
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/component/uploadprovider"
	"food_delivery_be/middleware"
	"food_delivery_be/modules/restaurant/restauranttranspot/ginrestaurant"
	"food_delivery_be/modules/restaurantlike/transport/ginrestaurantlike"
	"food_delivery_be/modules/upload/uploadtransport/ginupload"
	"food_delivery_be/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DatabaseConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect database - ", err)
	}
	db = db.Debug()
	fmt.Println("Connected to database")
	runServices(db, s3Provider, secretKey)
}

func runServices(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) {
	appCtx := component.NewAppContent(db, upProvider, secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", middleware.RequireAuth(appCtx), ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", middleware.RequireAuth(appCtx), ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", middleware.RequireAuth(appCtx), ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", middleware.RequireAuth(appCtx), ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", middleware.RequireAuth(appCtx), ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", middleware.RequireAuth(appCtx), ginrestaurantlike.ListUserLikeRestaurant(appCtx))
	}

	v1.GET("encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(200, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
		})
	})

	r.Run()
}
