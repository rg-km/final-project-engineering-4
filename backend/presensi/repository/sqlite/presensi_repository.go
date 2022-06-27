package sqlite

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type presensiRepository struct {
	db *sql.DB
}

func NewPresensiRepository(db *sql.DB) domain.PresensiRepository {
	return &presensiRepository{db}
}

func (p *presensiRepository) GetByID(id int64) (*domain.Presensi, error) {
	query := `SELECT id_presensi, id_dt_kelassiswa, tanggal, status_hadir FROM presensi WHERE id_presensi = ? AND show_item = 1`
	row := p.db.QueryRow(query, id)

	var res domain.Presensi
	var idDetailKelas int64
	err := row.Scan(
		&res.ID,
		&idDetailKelas,
		&res.Tanggal,
		&res.StatusHadir,
	)
	if err != nil {
		return nil, err
	}

	res.DetailKelasSiswa = &domain.DetailKelasSiswa{ID: idDetailKelas}
	return &res, nil
}

func (p *presensiRepository) GetByDetailKelasID(id int64) ([]domain.Presensi, error) {
	query := `SELECT id_presensi, id_dt_kelassiswa, tanggal, status_hadir FROM presensi WHERE id_dt_kelassiswa = ? AND show_item = 1`
	rows, err := p.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []domain.Presensi

	for rows.Next() {
		var presensi domain.Presensi
		var idDetailKelas int64

		rows.Scan(
			&presensi.ID,
			&idDetailKelas,
			&presensi.Tanggal,
			&presensi.StatusHadir,
		)
		if err != nil {
			return nil, err
		}

		presensi.DetailKelasSiswa = &domain.DetailKelasSiswa{ID: idDetailKelas}
		res = append(res, presensi)

	}

	return res, nil
}

func (p *presensiRepository) Create(presensi domain.Presensi) error {
	query := `INSERT INTO presensi (id_dt_kelassiswa, tanggal, jam, status_hadir) VALUES (?, ?, ?,?)`
	_, err := p.db.Exec(query, presensi.DetailKelasSiswa.ID, presensi.Tanggal, presensi.Jam, presensi.StatusHadir)
	return err
}
