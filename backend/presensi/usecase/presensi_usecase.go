package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type presensiUseCase struct {
	orangTuaRepo    domain.OrangTuaRepository
	detailKelasRepo domain.DetailKelasSiswaRepository
	presensiRepo    domain.PresensiRepository
	siswaRepo       domain.SiswaRepository
	guruRepo        domain.GuruRepository
}

func NewPresensiUseCase(orangTuaRepo domain.OrangTuaRepository, detailKelasRepo domain.DetailKelasSiswaRepository, presensiRepo domain.PresensiRepository, siswaRepo domain.SiswaRepository, guruRepo domain.GuruRepository) domain.PresensiUseCase {
	return &presensiUseCase{
		orangTuaRepo:    orangTuaRepo,
		detailKelasRepo: detailKelasRepo,
		presensiRepo:    presensiRepo,
		siswaRepo:       siswaRepo,
		guruRepo:        guruRepo,
	}
}

func (p *presensiUseCase) FetchPresensiSiswa(idKelas int64, emailOrtu string) ([]domain.Presensi, error) {
	var presensi []domain.Presensi

	ortu, err := p.orangTuaRepo.GetByEmail(emailOrtu)
	if err != nil {
		return nil, domain.ErrEmailNotFound
	}
	detailKelas, err := p.detailKelasRepo.GetByKelasID(idKelas)
	if err != nil {
		return nil, domain.ErrDetailKelasNotFound
	}
	found := false
	for _, dt := range detailKelas {
		if dt.Siswa.ID == ortu.Siswa.ID {
			presensi, _ = p.presensiRepo.GetByDetailKelasID(dt.ID)
			found = true
			break
		}
	}
	if !found {
		return nil, domain.ErrKelasNotFound
	}

	siswa, _ := p.siswaRepo.GetByID(ortu.Siswa.ID)
	for i := range presensi {
		presensi[i].DetailKelasSiswa = &domain.DetailKelasSiswa{
			Siswa: siswa,
		}
	}

	return presensi, nil
}

func (p *presensiUseCase) CreatePresensiSiswa(idSiswa, idKelas int64, presensi domain.Presensi) (*domain.Presensi, error) {
	detailKelas, err := p.detailKelasRepo.GetByKelasID(idKelas)
	if err != nil {
		return nil, domain.ErrDetailKelasNotFound
	}

	for _, dt := range detailKelas {
		if dt.Siswa.ID == idSiswa {
			presensi.DetailKelasSiswa = &domain.DetailKelasSiswa{
				ID: dt.ID,
			}
			break
		}
	}
	// _, err = p.siswaRepo.GetByEmail(presensi.DetailKelasSiswa.Siswa.Email)
	// if err != nil {
	// 	return nil, domain.ErrEmailSiswaNotFound
	// }

	err = p.presensiRepo.Create(presensi)
	if err != nil {
		return nil, err
	}

	return &presensi, nil
}
