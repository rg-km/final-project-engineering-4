package sqlite

import (
	"database/sql"
	"errors"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type orangTuaRepository struct {
	db *sql.DB
}

func NewOrangTuaRepository(db *sql.DB) domain.OrangTuaRepository {
	return &orangTuaRepository{db}
}

func (o *orangTuaRepository) GetByUsername(username string) (*domain.OrangTua, error) {
	query := `SELECT * FROM orangtua WHERE username = ? AND show_item = 1`
	row := o.db.QueryRow(query, username)

	var res domain.OrangTua
	var idSiswa int64
	err := row.Scan(
		&res.ID,
		&res.Username,
		&res.Email,
		&res.Password,
		&res.Nama,
		&res.JenisKelamin,
		&res.NoHP,
		&res.Alamat,
		&idSiswa,
	)
	if err != nil {
		return nil, errors.New("orangtua not found")
	}

	res.Siswa = domain.Siswa{ID: idSiswa}
	return &res, nil
}

func (o *orangTuaRepository) GetByEmail(email string) (*domain.OrangTua, error) {
	query := `SELECT * FROM wali_orangtua WHERE email = ? AND show_item = 1`
	row := o.db.QueryRow(query, email)

	var res domain.OrangTua
	var idSiswa int64
	err := row.Scan(
		&res.ID,
		&res.Username,
		&res.Email,
		&res.Password,
		&res.Nama,
		&res.JenisKelamin,
		&res.NoHP,
		&res.Alamat,
		&idSiswa,
	)
	if err != nil {
		return nil, errors.New("orangtua not found")
	}

	res.Siswa = domain.Siswa{ID: idSiswa}
	return &res, nil
}

func (o *orangTuaRepository) Create(orangTua domain.OrangTua) error {
	query := `INSERT INTO orangtua (username, email, password, nama, gender, no_hp, alamat, siswa) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := o.db.Exec(query, orangTua.Username, orangTua.Email, orangTua.Password, orangTua.Nama, orangTua.JenisKelamin, orangTua.NoHP, orangTua.Alamat, orangTua.Siswa.ID)
	return err
}
