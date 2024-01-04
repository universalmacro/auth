package services

import (
	"github.com/universalmacro/auth/dao/repositories"
	"github.com/universalmacro/common/singleton"
)

var adminService = singleton.NewSingleton[AdminService](newAdminService, singleton.Eager)

func newAdminService() *AdminService {
	return &AdminService{
		adminRepository: repositories.GetAdminRepository(),
	}
}

func GetAdminService() *AdminService {
	return adminService.Get()
}

type AdminService struct {
	adminRepository *repositories.AdminRepository
}
