package services

import (
	"time"

	"github.com/gin-gonic/gin"
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
		Email: email,
		Password: models.Password{
			Password: hashed,
			Salt:     salt,
		},
		Role: string(USER),
	}
	a.accountRepository.Create(account)
	return Account{entity: *account}, nil
}

func (a AuthService) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := Authorize(c)
		if auth.Status == Authorized {
			account := a.GetAccount(auth.AccountId)
			c.Set("account", *account)
		}
		c.Next()
	}
}

func (a AuthService) GetAccount(id uint) *Account {
	account, _ := a.accountRepository.FindOne(id)
	if account == nil {
		return nil
	}
	return &Account{*account}
}

func (a AuthService) CreateSession(email, password string) (string, error) {
	account, _ := a.accountRepository.FindOne("email = ?", email)
	if account == nil {
		return "", fault.ErrUnauthorized
	}
	if !account.PasswordMatching(password) {
		return "", fault.ErrUnauthorized
	}
	expiredAt := time.Now().AddDate(1, 0, 0).Unix()
	token, err := utils.SignJwt(
		utils.UintToString(account.ID()),
		account.Email,
		string(account.Role),
		expiredAt,
	)
	if err != nil {
		return "", fault.ErrUndefined
	}
	return token, nil
}
