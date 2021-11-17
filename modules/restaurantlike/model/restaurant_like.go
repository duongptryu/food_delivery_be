package restaurantlikemodel

import (
	"fmt"
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

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func (l *Like) GetOwnerId() int {
	return l.UserId
}

func ErrUserCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprint("Cannot like this restaurant"),
		fmt.Sprint("ErrUserCannotLikeRestaurant"),
	)
}

func ErrUserCannotUnLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprint("Cannot unlike this restaurant"),
		fmt.Sprint("ErrUserCannotUnLikeRestaurant"),
	)
}
