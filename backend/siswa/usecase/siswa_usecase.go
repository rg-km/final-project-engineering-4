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

func (s *siswaUseCase) Register(siswa domain.Siswa) (*domain.Siswa, error) {
	_, err := s.siswaRepo.GetByUsername(siswa.Username)
	if err == nil {
		return nil, domain.ErrUsernameExists
	}

	_, err = s.siswaRepo.GetByEmail(siswa.Email)
	if err == nil {
		return nil, domain.ErrEmailExists
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(siswa.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	siswa.Password = string(hashedPass)

	err = s.siswaRepo.Create(siswa)
	if err != nil {
		return nil, err
	}
	return s.siswaRepo.GetByUsername(siswa.Username)
}
