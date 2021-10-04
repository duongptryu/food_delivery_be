package restaurantstorage

import (
	"context"
	"food_delivery_be/common"
	"food_delivery_be/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant
	db := s.db

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, err
	}

	return &result, nil

}
