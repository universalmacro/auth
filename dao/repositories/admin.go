package repositories

import (
	"github.com/universalmacro/auth/dao/models"
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/singleton"
)

var adminRepository = singleton.NewSingleton[AdminRepository](func() *AdminRepository {
	return &AdminRepository{
		dao.NewRepository[models.Admin](),
	}
}, singleton.Eager)

func GetAdminRepository() *AdminRepository {
	return adminRepository.Get()
}

type AdminRepository struct {
	*dao.Repository[models.Admin]
}
