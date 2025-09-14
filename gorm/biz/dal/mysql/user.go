package mysql

import (
	"github.com/ggstudy11/learn_go_from_cloudwego/gorm/biz/model"
)

// crud for user
func CreateUser(user []*model.User) error {
	return DB.Create(user).Error
}

func DeleteUser(userId int64) error {
	return DB.Where("id = ?", userId).Delete(&model.User{}).Error
}

func UpdateUser(user *model.User) error {
	return DB.Updates(user).Error
}

func QueryUser(keyword *string, page, pageSize int64) ([]*model.User, int64, error) {
	db := DB.Model(&model.User{})

	if keyword != nil && len(*keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+*keyword+"%").Or("introduce LIKE ?", "%"+*keyword+"%")
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var res []*model.User
	if err := db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, total, nil
}
