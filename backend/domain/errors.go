package domain

import "errors"

var (
	ErrUsernameExists     = errors.New("Username Telah Digunakan")
	ErrEmailExists        = errors.New("Email Telah Digunakan")
	ErrEmailSiswaNotFound = errors.New("Email Siswa Tidak Ditemukan")
	ErrPasswordWrong      = errors.New("Password Salah")
	ErrEmailNotFound      = errors.New("Email Tidak Ditemukan")
	ErrGuruNotFound       = errors.New("Guru Tidak Ditemukan")
	ErrInvalidRequest     = errors.New("Request Invalid")
	ErrInvalidToken       = errors.New("Token Invalid")
)
