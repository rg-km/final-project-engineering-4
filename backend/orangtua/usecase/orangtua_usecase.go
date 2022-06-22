package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/util"
	"golang.org/x/crypto/bcrypt"
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

func (o *orangTuaUseCase) Register(orangTua domain.OrangTua) (*domain.OrangTua, error) {
	_, err := o.orangTuaRepo.GetByEmail(orangTua.Email)
	if err == nil {
		return nil, domain.ErrEmailExists
	}

	siswa, err := o.siswaRepo.GetByEmail(orangTua.Siswa.Email)
	if err != nil {
		return nil, domain.ErrEmailSiswaNotFound
	}
	orangTua.Siswa = siswa

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(orangTua.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	orangTua.Password = string(hashedPass)

	err = o.orangTuaRepo.Create(orangTua)
	if err != nil {
		return nil, err
	}

	res, _ := o.orangTuaRepo.GetByEmail(orangTua.Email)
	res.Siswa = siswa

	return res, nil
}

func (o *orangTuaUseCase) Login(email, password string) (*domain.OrangTua, string, error) {
	orangTua, err := o.orangTuaRepo.GetByEmail(email)
	if err != nil {
		return nil, "", domain.ErrEmailNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(orangTua.Password), []byte(password))
	if err != nil {
		return nil, "", domain.ErrPasswordWrong
	}

	token, err := util.GenerateToken(orangTua.ID, "Orang Tua")
	if err != nil {
		return nil, "", err
	}

	orangTua.Siswa = nil
	return orangTua, token, nil
}
