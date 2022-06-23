package sqlite

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type kelasRepository struct {
	db *sql.DB
}

func NewKelasRepository(db *sql.DB) domain.KelasRepository {
	return &kelasRepository{db}
}

func (k *kelasRepository) GetByCode(code string) (*domain.Kelas, error) {
	query := `SELECT id_class, id_guru, nama_kelas, keterangan, kode_kelas, file_name FROM class WHERE kode_kelas = ? AND show_item = 1`
	row := k.db.QueryRow(query, code)

	var res domain.Kelas
	var guruID int64
	err := row.Scan(
		&res.ID,
		&guruID,
		&res.NamaKelas,
		&res.Keterangan,
		&res.KodeKelas,
		&res.Filename,
	)
	if err != nil {
		return nil, err
	}

	res.Guru = &domain.Guru{
		ID: guruID,
	}

	return &res, nil
}

func (k *kelasRepository) Create(kelas domain.Kelas) error {
	query := `INSERT INTO class (id_guru, nama_kelas, keterangan, kode_kelas, file_name, status_kelas) VALUES (?, ?, ?, ?, ?, '1')`
	_, err := k.db.Exec(query, kelas.Guru.ID, kelas.NamaKelas, kelas.Keterangan, kelas.KodeKelas, kelas.Filename)
	return err
}
