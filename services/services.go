package services

import (
	"github.com/universalmacro/auth/dao/models"
	"github.com/universalmacro/auth/dao/repositories"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/common/utils"
)

var authService = singleton.NewSingleton[AuthService](newAuthService, singleton.Eager)

func GetAuthService() *AuthService {
	return authService.Get()
}

type AuthService struct {
	accountRepository *repositories.AccountRepository
}

func newAuthService() *AuthService {
	return &AuthService{
		accountRepository: repositories.GetAccounrRepository(),
	}
}

func (a AuthService) CreateAccount(email, password string) (Account, error) {
	if account, _ := a.accountRepository.FindOne("email = ?", email); account != nil {
		return Account{}, fault.ErrEmailExists
	}
	hashed, salt := utils.HashWithSalt(password)
	account := &models.Account{
		Email:    email,
		Password: hashed,
		Salt:     salt,
		Role:     string(USER),
	}
	a.accountRepository.Create(account)
	return Account{entity: *account}, nil
}
