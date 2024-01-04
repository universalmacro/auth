package repositories

import (
	"github.com/universalmacro/auth/dao/models"
	"github.com/universalmacro/common/dao"
)

// type TestingObject struct {
// 	gorm.Model
// 	Name        *string
// 	RegionCode  string `gorm:"index:idx_phone_number,unique"`
// 	PhoneNumber string `gorm:"index:idx_phone_number,unique"`
// }

func init() {
	db := dao.GetDBInstance()
	db.AutoMigrate(&models.Account{}, &models.VerificationCode{}, &models.Admin{})
}
