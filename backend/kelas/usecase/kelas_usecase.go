package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/utils"
)

type kelasUseCase struct {
	kelasRepo       domain.KelasRepository
	guruRepo        domain.GuruRepository
	siswaRepo       domain.SiswaRepository
	orangTuaRepo    domain.OrangTuaRepository
	detailKelasRepo domain.DetailKelasSiswaRepository
}

func NewKelasUseCase(kelasRepo domain.KelasRepository, guruRepo domain.GuruRepository, siswaRepo domain.SiswaRepository, orangTuaRepo domain.OrangTuaRepository, detailKelasRepo domain.DetailKelasSiswaRepository) domain.KelasUseCase {
	return &kelasUseCase{
		kelasRepo:       kelasRepo,
		guruRepo:        guruRepo,
		siswaRepo:       siswaRepo,
		orangTuaRepo:    orangTuaRepo,
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

func (k *kelasUseCase) FetchKelas(role, email string) ([]domain.Kelas, error) {
	var res []domain.Kelas

	if role == "Siswa" || role == "Orang Tua" {
		var id int64

		if role == "Siswa" {
			siswa, err := k.siswaRepo.GetByEmail(email)
			if err != nil {
				return nil, domain.ErrEmailNotFound
			}
			id = siswa.ID
		} else {
			orangtua, err := k.orangTuaRepo.GetByEmail(email)
			if err != nil {
				return nil, domain.ErrEmailNotFound
			}
			id = orangtua.Siswa.ID
		}

		detailKelas, err := k.detailKelasRepo.GetBySiswaID(id)
		if err != nil {
			return nil, domain.ErrDetailKelasNotFound
		}

		for _, dt := range detailKelas {
			kelas, _ := k.kelasRepo.GetByID(dt.Kelas.ID)
			kelas.Guru = nil
			res = append(res, *kelas)
		}
	} else if role == "Guru" {
		guru, err := k.guruRepo.GetByEmail(email)
		if err != nil {
			return nil, domain.ErrEmailNotFound
		}

		res, err = k.kelasRepo.GetByGuruID(guru.ID)
		if err != nil {
			return nil, err
		}

		for _, r := range res {
			r.Guru = nil
		}
	}

	return res, nil
}

func (k *kelasUseCase) FetchKelasByID(role, email string, kelasID int64) (*domain.Kelas, error) {
	var res *domain.Kelas

	if role == "Siswa" || role == "Orang Tua" {
		var id int64

		if role == "Siswa" {
			siswa, err := k.siswaRepo.GetByEmail(email)
			if err != nil {
				return nil, domain.ErrEmailNotFound
			}
			id = siswa.ID
		} else {
			orangtua, err := k.orangTuaRepo.GetByEmail(email)
			if err != nil {
				return nil, domain.ErrEmailNotFound
			}
			id = orangtua.Siswa.ID
		}

		detailKelas, err := k.detailKelasRepo.GetBySiswaID(id)
		if err != nil {
			return nil, domain.ErrDetailKelasNotFound
		}

		found := false
		for _, dt := range detailKelas {
			if dt.Kelas.ID == kelasID {
				found = true
				res, _ = k.kelasRepo.GetByID(dt.Kelas.ID)
				guru, _ := k.guruRepo.GetByID(res.Guru.ID)
				res.Guru = guru
				break
			}
		}

		if !found {
			return nil, domain.ErrKelasNotFound
		}
	} else if role == "Guru" {
		guru, err := k.guruRepo.GetByEmail(email)
		if err != nil {
			return nil, domain.ErrEmailNotFound
		}

		kelas, err := k.kelasRepo.GetByID(kelasID)
		if err != nil {
			return nil, err
		}

		if kelas.Guru.ID != guru.ID {
			return nil, domain.ErrKelasNotFound
		}

		res = kelas
		res.Guru = guru
	}

	return res, nil
}
