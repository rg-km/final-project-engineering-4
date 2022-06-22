package domain

import "errors"

var (
	ErrUsernameExists     = errors.New("Username Telah Digunakan")
	ErrEmailExists        = errors.New("Email Telah Digunakan")
	ErrEmailSiswaNotFound = errors.New("Email Siswa Tidak Ditemukan")
	ErrPasswordWrong      = errors.New("Password Salah")
	ErrUsernameWrong      = errors.New("Username Salah")
)
