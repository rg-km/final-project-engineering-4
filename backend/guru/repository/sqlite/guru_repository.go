package sqlite

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type guruRepository struct {
	db *sql.DB
}

func NewGuruRepository(db *sql.DB) domain.GuruRepository {
	return &guruRepository{db}
}

func (g *guruRepository) GetByID(id int64) (*domain.Guru, error) {
	query := `SELECT id_guru, email, password, nama, jenis_kelamin, no_hp, alamat FROM guru WHERE id_guru = ? AND show_item = 1`
	row := g.db.QueryRow(query, id)

	var res domain.Guru
	err := row.Scan(
		&res.ID,
		&res.Email,
		&res.Password,
		&res.Nama,
		&res.JenisKelamin,
		&res.NoHP,
		&res.Alamat,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (g *guruRepository) GetByEmail(email string) (*domain.Guru, error) {
	query := `SELECT id_guru, email, password, nama, jenis_kelamin, no_hp, alamat FROM guru WHERE email = ? AND show_item = 1`
	row := g.db.QueryRow(query, email)

	var res domain.Guru
	err := row.Scan(
		&res.ID,
		&res.Email,
		&res.Password,
		&res.Nama,
		&res.JenisKelamin,
		&res.NoHP,
		&res.Alamat,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (g *guruRepository) Create(guru domain.Guru) error {
	query := `INSERT INTO guru (email, password, nama, jenis_kelamin, no_hp, alamat, agama, pendidikan, file_name) VALUES (?,?,?,?,?,?,?,?,?)`
	_, err := g.db.Exec(query,
		guru.Email,
		guru.Password,
		guru.Nama,
		guru.JenisKelamin,
		guru.NoHP,
		guru.Alamat,
		guru.Agama,
		guru.Pendidikan,
		guru.Filename,
	)
	return err
}
