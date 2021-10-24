package main

import (
	"fmt"
	"food_delivery_be/component"
	"food_delivery_be/component/uploadprovider"
	"food_delivery_be/middleware"
	"food_delivery_be/modules/restaurant/restauranttranspot/ginrestaurant"
	"food_delivery_be/modules/upload/uploadtransport/ginupload"
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

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect database - ", err)
	}
	fmt.Println("Connected to database")
	runServices(db, s3Provider)
}

func runServices(db *gorm.DB, upProvider uploadprovider.UploadProvider) {
	appCtx := component.NewAppContent(db, upProvider)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.POST("/upload", ginupload.Upload(appCtx))

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	r.Run()
}
