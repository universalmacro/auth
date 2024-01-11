package models

import (
	"time"

	abstract "github.com/universalmacro/common/abstract"
	"github.com/universalmacro/common/auth"
	"github.com/universalmacro/common/snowflake"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	auth.Password
	Account  *string `json:"account" gorm:"index:uniqueIndex"`
	Email    string  `json:"email" gorm:"index:email_index,unique"`
	Role     string  `json:"role" gorm:"type:VARCHAR(128)"`
	Gender   string  `json:"gender"`
	Birthday time.Time
}

func (a Account) ID() uint {
	return a.Model.ID
}

func (a Account) Own(asset abstract.Asset) bool {
	return abstract.Own(a, asset)
}

var accountIdGenerator = snowflake.NewIdGenertor(1)

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = accountIdGenerator.Uint()
	return err
}

func (a Account) Owner() abstract.Owner {
	return nil
}
