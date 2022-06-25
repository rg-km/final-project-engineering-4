package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/utils"
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
	_, err := s.siswaRepo.GetByEmail(siswa.Email)
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
	return s.siswaRepo.GetByEmail(siswa.Email)
}

func (s *siswaUseCase) Login(email, password string) (*domain.Siswa, string, error) {
	siswa, err := s.siswaRepo.GetByEmail(email)
	if err != nil {
		return nil, "", domain.ErrEmailNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(siswa.Password), ([]byte(password)))
	if err != nil {
		return nil, "", domain.ErrPasswordWrong
	}

	token, err := utils.GenerateToken(siswa.Email, "Siswa")
	if err != nil {
		return nil, "", err
	}

	return siswa, token, nil
}
