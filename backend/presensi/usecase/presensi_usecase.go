package usecase

import (
	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type presensiUseCase struct {
	orangTuaRepo    domain.OrangTuaRepository
	detailKelasRepo domain.DetailKelasSiswaRepository
	presensiRepo    domain.PresensiRepository
	siswaRepo       domain.SiswaRepository
}

func NewPresensiUseCase(orangTuaRepo domain.OrangTuaRepository, detailKelasRepo domain.DetailKelasSiswaRepository, presensiRepo domain.PresensiRepository, siswaRepo domain.SiswaRepository) domain.PresensiUseCase {
	return &presensiUseCase{
		orangTuaRepo:    orangTuaRepo,
		detailKelasRepo: detailKelasRepo,
		presensiRepo:    presensiRepo,
		siswaRepo:       siswaRepo,
	}
}

func (p *presensiUseCase) FetchPresensi(idKelas int64, emailOrtu string) ([]domain.Presensi, error) {
	var presensi []domain.Presensi

	ortu, err := p.orangTuaRepo.GetByEmail(emailOrtu)
	if err != nil {
		return nil, domain.ErrEmailNotFound
	}
	detailKelas, err := p.detailKelasRepo.GetBySiswaID(idKelas)
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
