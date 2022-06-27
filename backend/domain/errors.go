package domain

import "errors"

var (
	ErrUsernameExists      = errors.New("Username Telah Digunakan")
	ErrEmailExists         = errors.New("Email Telah Digunakan")
	ErrClassJoined         = errors.New("Kelas sudah di-enroll")
	ErrEmailSiswaNotFound  = errors.New("Email Siswa Tidak Ditemukan")
	ErrPasswordWrong       = errors.New("Password Salah")
	ErrEmailNotFound       = errors.New("Email Tidak Ditemukan")
	ErrKelasNotFound       = errors.New("Kelas Tidak Ditemukan")
	ErrGuruNotFound        = errors.New("Guru Tidak Ditemukan")
	ErrDetailKelasNotFound = errors.New("Detail Tidak Ditemukan")
	ErrInvalidRequest      = errors.New("Request Invalid")
	ErrInvalidToken        = errors.New("Token Invalid")
)
