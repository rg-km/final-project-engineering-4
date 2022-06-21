package domain

import "errors"

var (
	ErrUsernameExists     = errors.New("Username Telah Digunakan")
	ErrEmailExists        = errors.New("Email Telah Digunakan")
	ErrEmailSiswaNotFound = errors.New("Email Siswa Tidak Ditemukan")
	ErrUsernameWrong      = errors.New("Username salah")
	ErrPasswordWrong      = errors.New("Password salah")
)
