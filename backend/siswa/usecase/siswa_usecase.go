package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"golang.org/x/crypto/bcrypt"
)

type siswaUseCase struct {
	siswaRepo domain.SiswaRepository
}

func NewSiswaUseCase(siswaRepo domain.SiswaRepository) domain.SiswaUseCase {
	return &siswaUseCase{
		siswaRepo: siswaRepo,
	}
}

func (s *siswaUseCase) Register(siswa domain.Siswa) error {
	_, err := s.siswaRepo.GetByUsername(siswa.Username)
	if err == nil {
		return domain.ErrUsernameExists
	}

	_, err = s.siswaRepo.GetByEmail(siswa.Email)
	if err == nil {
		return domain.ErrEmailExists
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(siswa.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	siswa.Password = string(hashedPass)

	return s.siswaRepo.Create(siswa)
}
