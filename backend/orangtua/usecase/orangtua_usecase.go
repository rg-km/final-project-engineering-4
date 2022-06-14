package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type orangTuaUseCase struct {
	orangTuaRepo domain.OrangTuaRepository
	siswaRepo    domain.SiswaRepository
}

func NewOrangTuaUseCase(orangTuaRepo domain.OrangTuaRepository, siswaRepo domain.SiswaRepository) domain.OrangTuaUseCase {
	return &orangTuaUseCase{
		orangTuaRepo: orangTuaRepo,
		siswaRepo:    siswaRepo,
	}
}

func (o *orangTuaUseCase) Register(orangTua domain.OrangTua) error {
	_, err := o.orangTuaRepo.GetByUsername(orangTua.Username)
	if err == nil {
		return domain.ErrUsernameExists
	}

	_, err = o.orangTuaRepo.GetByEmail(orangTua.Email)
	if err == nil {
		return domain.ErrEmailExists
	}

	_, err = o.siswaRepo.GetByEmail(orangTua.Siswa.Email)
	if err != nil {
		return domain.ErrEmailSiswaNotFound
	}

	return o.orangTuaRepo.Create(orangTua)
}
