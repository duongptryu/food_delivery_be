package restaurantlikestorage

import (
	"context"
	"food_delivery_be/common"
	restaurantlikemodel "food_delivery_be/modules/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
