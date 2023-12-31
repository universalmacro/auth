package repositories

import (
	"github.com/universalmacro/auth/dao/models"
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/singleton"
)

var verificationCodeRepository = singleton.NewSingleton[VerificationCodeRepository](newVerificationCodeRepository, singleton.Eager)

func GetVerificationCodeRepository() *VerificationCodeRepository {
	return verificationCodeRepository.Get()
}

type VerificationCodeRepository struct {
	*dao.Repository[models.VerificationCode]
}

func newVerificationCodeRepository() *VerificationCodeRepository {
	return &VerificationCodeRepository{
		dao.NewRepository[models.VerificationCode](),
	}
}
