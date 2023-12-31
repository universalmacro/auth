package models

import (
	"gorm.io/gorm"
)

type VerificationCode struct {
	gorm.Model
	Identity string `json:"identity" gorm:"index:verification_identity_index"`
	Code     string `json:"code" gorm:"type:VARCHAR(12)"`
	Tries    int64
}
