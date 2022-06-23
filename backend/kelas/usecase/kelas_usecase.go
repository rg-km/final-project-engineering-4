package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/utils"
)

type kelasUseCase struct {
	kelasRepo domain.KelasRepository
	guruRepo  domain.GuruRepository
}

func NewKelasUseCase(kelasRepo domain.KelasRepository, guruRepo domain.GuruRepository) domain.KelasUseCase {
	return &kelasUseCase{
		kelasRepo: kelasRepo,
		guruRepo:  guruRepo,
	}
}

func (k *kelasUseCase) CreateKelas(kelas domain.Kelas) (*domain.Kelas, error) {
	kelas.KodeKelas = utils.GenerateClassCode()
	guru, err := k.guruRepo.GetByEmail(kelas.Guru.Email)
	if err != nil {
		return nil, domain.ErrGuruNotFound
	}
	kelas.Guru = guru

	err = k.kelasRepo.Create(kelas)
	if err != nil {
		return nil, err
	}

	return &kelas, nil
}
