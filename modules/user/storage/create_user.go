package storage

import (
	"context"
	"food-client/common"
	"food-client/modules/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *model.UserCreation) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
