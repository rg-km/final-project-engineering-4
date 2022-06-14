package sqlite

import (
	"database/sql"
	"errors"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type siswaRepository struct {
	db *sql.DB
}

func NewSiswaRepository(db *sql.DB) domain.SiswaRepository {
	return &siswaRepository{db}
}

func (s *siswaRepository) GetByEmail(email string) (*domain.Siswa, error) {
	query := `SELECT * FROM siswa WHERE email = ?`
	row := s.db.QueryRow(query, email)

	var res domain.Siswa
	err := row.Scan(
		&res.ID,
		&res.Username,
		&res.Email,
		&res.Password,
		&res.Nama,
		&res.Gender,
		&res.NoHP,
		&res.TanggalLahir,
		&res.Alamat,
		&res.TempatLahir,
		&res.Filename,
	)
	if err != nil {
		return nil, errors.New("siswa not found")
	}

	return &res, nil
}
