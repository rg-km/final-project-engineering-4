package sqlite_test

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
	repository "github.com/rg-km/final-project-engineering-4/backend/siswa/repository/sqlite"
)

var _ = Describe("SiswaRepository", func() {
	var (
		db              *sql.DB
		mock            sqlmock.Sqlmock
		err             error
		siswaRepository domain.SiswaRepository
	)

	BeforeEach(func() {
		db, mock, err = sqlmock.New()
		Expect(err).ToNot(HaveOccurred())
		siswaRepository = repository.NewSiswaRepository(db)
	})

	Describe("get siswa by its email", func() {
		var (
			query     string
			email     string
			mockSiswa domain.Siswa
		)

		BeforeEach(func() {
			query = "SELECT \\* FROM siswa WHERE email = \\?"
			email = "Email1"
			mockSiswa = domain.Siswa{
				ID:           1,
				Username:     "Username1",
				Email:        "Email1",
				Password:     "Password1",
				Nama:         "Nama1",
				JenisKelamin: "JenisKelamin1",
				NoHP:         "NoHP1",
				TanggalLahir: "TanggalLahir1",
				Alamat:       "Alamat1",
				TempatLahir:  "TempatLahir1",
				Filename:     "Filename1",
			}
		})

		When("siswa is not exist", func() {
			It("should return error not found", func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "nama", "gender", "no_hp", "tanggal_lahir", "alamat", "tempat", "filename"})
				mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)
				siswa, err := siswaRepository.GetByEmail(email)
				Expect(err).To(HaveOccurred())
				Expect(siswa).To(BeNil())
			})
		})

		When("siswa is exist", func() {
			It("should return siswa with the given email", func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "nama", "gender", "no_hp", "tanggal_lahir", "alamat", "tempat", "filename"}).
					AddRow(mockSiswa.ID, mockSiswa.Username, mockSiswa.Email, mockSiswa.Password, mockSiswa.Nama, mockSiswa.JenisKelamin, mockSiswa.NoHP, mockSiswa.TanggalLahir, mockSiswa.Alamat, mockSiswa.TempatLahir, mockSiswa.Filename)
				mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)
				siswa, err := siswaRepository.GetByEmail(email)
				Expect(err).ToNot(HaveOccurred())
				Expect(siswa).ToNot(BeNil())
				Expect(*siswa).To(Equal(mockSiswa))
			})
		})
	})
})
