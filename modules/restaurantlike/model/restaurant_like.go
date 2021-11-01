package restaurantlikemodel

import (
	"food_delivery_be/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"restaurant_id"`
	UserId       int                `json:"user_id" gorm:"user_id"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"created_at"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (l Like) TableName() string {
	return "restaurant_likes"
}
