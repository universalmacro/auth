package services

import (
	"github.com/universalmacro/auth/dao/models"
	abstract "github.com/universalmacro/common/abstract"
)

type AccountRole string

const (
	ROOT  AccountRole = "ROOT"
	ADMIN AccountRole = "ADMIN"
	USER  AccountRole = "USER"
)

type Account struct {
	entity models.Account
}

func (a Account) ID() uint {
	return a.entity.ID()
}

func (a Account) Own(asset abstract.Asset) bool {
	return abstract.Own(a, asset)
}

func (a Account) Owner() abstract.Owner {
	return nil
}

func (a Account) Role() AccountRole {
	return AccountRole(a.entity.Role)
}

func (a Account) Email() string {
	return a.entity.Email
}
