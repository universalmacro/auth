package repositories

import (
	"github.com/universalmacro/auth/dao/models"
	"github.com/universalmacro/common/dao"
)

func init() {
	db := dao.GetDBInstance()
	db.AutoMigrate(&models.Account{}, &models.VerificationCode{})
}
