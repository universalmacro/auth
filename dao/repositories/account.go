package repositories

import (
	"github.com/universalmacro/auth/dao/models"
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/singleton"
)

var accountRepository = singleton.NewSingleton[AccountRepository](func() *AccountRepository {
	return &AccountRepository{
		dao.NewRepository[models.Account](),
	}
}, singleton.Eager)

func GetAccounrRepository() *AccountRepository {
	return accountRepository.Get()
}

type AccountRepository struct {
	*dao.Repository[models.Account]
}
