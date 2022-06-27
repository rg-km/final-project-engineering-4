package repository

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type detailKelasSiswaRepository struct {
	db *sql.DB
}

func NewDetailKelasSiswaRepository(db *sql.DB) domain.DetailKelasSiswaRepository {
	return &detailKelasSiswaRepository{db}
}

func (d *detailKelasSiswaRepository) GetByKelasID(kelasID int64) ([]domain.DetailKelasSiswa, error) {
	query := `SELECT id_dt_kelassiswa, id_siswa, id_class FROM dt_kelassiswa WHERE id_class = ? AND show_item = 1`
	rows, err := d.db.Query(query, kelasID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []domain.DetailKelasSiswa
	for rows.Next() {
		var detailKelas domain.DetailKelasSiswa
		var idSiswa, idKelas int64

		rows.Scan(
			&detailKelas.ID,
			&idSiswa,
			&idKelas,
		)

		detailKelas.Siswa = &domain.Siswa{ID: idSiswa}
		detailKelas.Kelas = &domain.Kelas{ID: idKelas}
		res = append(res, detailKelas)
	}

	if len(res) == 0 {
		return nil, domain.ErrDetailKelasNotFound
	}

	return res, nil
}

func (d *detailKelasSiswaRepository) GetBySiswaID(siswaID int64) ([]domain.DetailKelasSiswa, error) {
	query := `SELECT id_dt_kelassiswa, id_siswa, id_class FROM dt_kelassiswa WHERE id_siswa = ? AND show_item = 1`
	rows, err := d.db.Query(query, siswaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []domain.DetailKelasSiswa
	for rows.Next() {
		var detailKelas domain.DetailKelasSiswa
		var idSiswa, idKelas int64

		rows.Scan(
			&detailKelas.ID,
			&idSiswa,
			&idKelas,
		)

		detailKelas.Siswa = &domain.Siswa{ID: idSiswa}
		detailKelas.Kelas = &domain.Kelas{ID: idKelas}
		res = append(res, detailKelas)
	}

	if len(res) == 0 {
		return nil, domain.ErrDetailKelasNotFound
	}

	return res, nil
}

func (d *detailKelasSiswaRepository) Create(detailKelas domain.DetailKelasSiswa) error {
	query := `INSERT INTO dt_kelassiswa (id_siswa, id_class) VALUES (?, ?)`
	_, err := d.db.Exec(query, detailKelas.Siswa.ID, detailKelas.Kelas.ID)
	return err
}
