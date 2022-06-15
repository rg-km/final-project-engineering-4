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
	query := `SELECT id_siswa, username, email, password, nama, jenis_kelamin, no_hp, alamat FROM siswa WHERE email = ? AND show_item = 1`
	row := s.db.QueryRow(query, email)

	var res domain.Siswa
	err := row.Scan(
		&res.ID,
		&res.Username,
		&res.Email,
		&res.Password,
		&res.Nama,
		&res.JenisKelamin,
		&res.NoHP,
		&res.Alamat,
	)
	if err != nil {
		return nil, errors.New("siswa not found")
	}

	return &res, nil
}

func (o *siswaRepository) GetByUsername(username string) (*domain.Siswa, error) {
	query := `SELECT SELECT id_siswa, username, email, password, nama, jenis_kelamin, no_hp, alamat FROM siswa WHERE username = ? AND show_item = 1`
	row := o.db.QueryRow(query, username)

	var res domain.Siswa
	err := row.Scan(
		&res.ID,
		&res.Username,
		&res.Email,
		&res.Password,
		&res.Nama,
		&res.JenisKelamin,
		&res.NoHP,
		&res.Alamat,
	)
	if err != nil {
		return nil, errors.New("siswa not found")
	}
	return &res, nil
}

func (s *siswaRepository) Create(siswa domain.Siswa) error {
	query := `INSERT INTO siswa (username, email, password, nama, jenis_kelamin, no_hp, alamat, tempat_lahir, tanggal_lahir, agama, file_name) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?,?)`
	_, err := s.db.Exec(query,
		siswa.Username,
		siswa.Email,
		siswa.Password,
		siswa.Nama,
		siswa.JenisKelamin,
		siswa.NoHP,
		siswa.Alamat,
		siswa.TempatLahir,
		siswa.TanggalLahir,
		siswa.Agama,
		siswa.Filename,
	)
	return err
}
