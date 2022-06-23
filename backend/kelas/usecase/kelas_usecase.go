package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/utils"
)

type kelasUseCase struct {
	kelasRepo       domain.KelasRepository
	guruRepo        domain.GuruRepository
	siswaRepo       domain.SiswaRepository
	detailKelasRepo domain.DetailKelasSiswaRepository
}

func NewKelasUseCase(kelasRepo domain.KelasRepository, guruRepo domain.GuruRepository, siswaRepo domain.SiswaRepository, detailKelasRepo domain.DetailKelasSiswaRepository) domain.KelasUseCase {
	return &kelasUseCase{
		kelasRepo:       kelasRepo,
		guruRepo:        guruRepo,
		siswaRepo:       siswaRepo,
		detailKelasRepo: detailKelasRepo,
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

func (k *kelasUseCase) JoinKelas(emailSiswa, code string) (*domain.Kelas, error) {
	siswa, err := k.siswaRepo.GetByEmail(emailSiswa)
	if err != nil {
		return nil, domain.ErrEmailNotFound
	}

	kelas, err := k.kelasRepo.GetByCode(code)
	if err != nil {
		return nil, domain.ErrKelasNotFound
	}

	detailKelas, err := k.detailKelasRepo.GetBySiswaID(siswa.ID)
	if err == nil {
		for _, dt := range detailKelas {
			if dt.Kelas.ID == kelas.ID {
				return nil, domain.ErrClassJoined
			}
		}
	}

	newDetailKelas := domain.DetailKelasSiswa{
		Siswa: siswa,
		Kelas: kelas,
	}
	err = k.detailKelasRepo.Create(newDetailKelas)
	if err != nil {
		return nil, err
	}

	return kelas, nil
}
