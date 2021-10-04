package main

import (
	"fmt"
	"food_delivery_be/component"
	"food_delivery_be/modules/restaurant/restauranttranspot/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DatabaseConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect database - ", err)
	}
	fmt.Println("Connected to database")
	runServices(db)
}

func runServices(db *gorm.DB) {
	r := gin.Default()

	appCtx := component.NewAppContent(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		//restaurants.GET("/:id", func(c *gin.Context) {
		//	id, err := strconv.Atoi(c.Param("id"))
		//	if err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	var data Restaurant
		//
		//	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	c.JSON(http.StatusOK, data)
		//})
		//
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		//
		//restaurants.PATCH("/:id", func(c *gin.Context) {
		//	id, err := strconv.Atoi(c.Param("id"))
		//	if err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	var data RestaurantUpdate
		//
		//	if err := c.ShouldBind(&data); err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	c.JSON(http.StatusOK, gin.H{"ok": 1})
		//})
		//
		//restaurants.DELETE("/:id", func(c *gin.Context) {
		//	id, err := strconv.Atoi(c.Param("id"))
		//	if err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		//		c.JSON(401, gin.H{
		//			"error": err.Error(),
		//		})
		//		return
		//	}
		//
		//	c.JSON(http.StatusOK, gin.H{"ok": 1})
		//})
	}

	r.Run()
}
