package userstorage

import (
	"context"
	"food_delivery_be/common"
	"food_delivery_be/modules/user/usermodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	err := db.Where(condition).Find(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
