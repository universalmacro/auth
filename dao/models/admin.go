package models

import (
	"github.com/universalmacro/common/snowflake"
	"github.com/universalmacro/common/utils"
	"gorm.io/gorm"
)

var adminIdGenerator = snowflake.NewIdGenertor(0)

type Password struct {
	Password string `json:"password" gorm:"type:CHAR(128)"`
	Salt     []byte `json:"salt"`
}

func (p *Password) SetPassword(password string) (string, []byte) {
	hashed, salt := utils.HashWithSalt(password)
	p.Password = hashed
	p.Salt = salt
	return hashed, salt
}
func (p *Password) PasswordMatching(password string) bool {
	return utils.PasswordsMatch(p.Password, password, p.Salt)
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = adminIdGenerator.Uint()
	return err
}

type Admin struct {
	gorm.Model
	Account string `json:"account" gorm:"index:uniqueIndex"`
	Password
}
