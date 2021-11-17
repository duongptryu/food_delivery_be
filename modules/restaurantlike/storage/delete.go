package restaurantlikestorage

import (
	"context"
	"food_delivery_be/common"
	restaurantlikemodel "food_delivery_be/modules/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, userId, restaurantId int) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).Where("user_id = ? and restaurant_id = ?", userId, restaurantId).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}