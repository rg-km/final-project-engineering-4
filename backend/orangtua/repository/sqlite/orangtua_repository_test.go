package sqlite_test

import (
	"database/sql"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
	repository "github.com/rg-km/final-project-engineering-4/backend/orangtua/repository/sqlite"
)

var _ = Describe("OrangtuaRepository", func() {
	var (
		db                 *sql.DB
		mock               sqlmock.Sqlmock
		err                error
		orangTuaRepository domain.OrangTuaRepository
	)

	BeforeEach(func() {
		db, mock, err = sqlmock.New()
		Expect(err).ToNot(HaveOccurred())
		orangTuaRepository = repository.NewOrangTuaRepository(db)
	})

	Describe("get orangtua by its username", func() {
		var (
			query        string
			username     string
			mockOrangTua domain.OrangTua
		)

		BeforeEach(func() {
			query = "SELECT \\* FROM orangtua WHERE username = \\?"
			username = "Username1"
			mockOrangTua = domain.OrangTua{
				ID:           1,
				Username:     "Username1",
				Email:        "Email1",
				Password:     "Password1",
				Nama:         "Nama1",
				JenisKelamin: "JenisKelamin1",
				NoHP:         "NoHP1",
				Alamat:       "Alamat1",
				Siswa:        domain.Siswa{ID: 1},
			}
		})

		When("orangtua is not exist", func() {
			It("should return error not found", func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "nama", "jenis_kelamin", "no_hp", "alamat", "id_siswa"})
				mock.ExpectQuery(query).WithArgs(username).WillReturnRows(rows)
				orangtua, err := orangTuaRepository.GetByUsername(username)
				Expect(err).To(HaveOccurred())
				Expect(orangtua).To(BeNil())
			})
		})

		When("orangtua is not exist", func() {
			It("should return error not found", func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "nama", "jenis_kelamin", "no_hp", "alamat", "id_siswa"}).
					AddRow(mockOrangTua.ID, mockOrangTua.Username, mockOrangTua.Email, mockOrangTua.Password, mockOrangTua.Nama, mockOrangTua.JenisKelamin, mockOrangTua.NoHP, mockOrangTua.Alamat, mockOrangTua.Siswa.ID)
				mock.ExpectQuery(query).WithArgs(username).WillReturnRows(rows)
				orangtua, err := orangTuaRepository.GetByUsername(username)
				Expect(err).ToNot(HaveOccurred())
				Expect(orangtua).ToNot(BeNil())
			})
		})
	})

	Describe("get orangtua by its email", func() {
		var (
			query        string
			email        string
			mockOrangTua domain.OrangTua
		)

		BeforeEach(func() {
			query = "SELECT \\* FROM orangtua WHERE email = \\?"
			email = "Email1"
			mockOrangTua = domain.OrangTua{
				ID:           1,
				Username:     "Username1",
				Email:        "Email1",
				Password:     "Password1",
				Nama:         "Nama1",
				JenisKelamin: "JenisKelamin1",
				NoHP:         "NoHP1",
				Alamat:       "Alamat1",
				Siswa:        domain.Siswa{ID: 1},
			}
		})

		When("orangtua is not exist", func() {
			It("should return error not found", func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "nama", "jenis_kelamin", "no_hp", "alamat", "id_siswa"})
				mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)
				orangtua, err := orangTuaRepository.GetByEmail(email)
				Expect(err).To(HaveOccurred())
				Expect(orangtua).To(BeNil())
			})
		})

		When("orangtua is not exist", func() {
			It("should return error not found", func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "nama", "jenis_kelamin", "no_hp", "alamat", "id_siswa"}).
					AddRow(mockOrangTua.ID, mockOrangTua.Username, mockOrangTua.Email, mockOrangTua.Password, mockOrangTua.Nama, mockOrangTua.JenisKelamin, mockOrangTua.NoHP, mockOrangTua.Alamat, mockOrangTua.Siswa.ID)
				mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)
				orangtua, err := orangTuaRepository.GetByEmail(email)
				Expect(err).ToNot(HaveOccurred())
				Expect(orangtua).ToNot(BeNil())
			})
		})
	})

	Describe("create new orangtua", func() {
		var (
			query        string
			mockOrangTua domain.OrangTua
		)

		BeforeEach(func() {
			query = "INSERT INTO orangtua"
			mockOrangTua = domain.OrangTua{
				ID:           1,
				Username:     "Username1",
				Email:        "Email1",
				Password:     "Password1",
				Nama:         "Nama1",
				JenisKelamin: "JenisKelamin1",
				NoHP:         "NoHP1",
				Alamat:       "Alamat1",
				Siswa: domain.Siswa{
					ID:           1,
					Username:     "Username1",
					Email:        "Email1",
					Password:     "Password1",
					Nama:         "Nama1",
					Gender:       "Gender1",
					NoHP:         "NoHP1",
					TanggalLahir: "TanggalLahir1",
					Alamat:       "Alamat1",
					TempatLahir:  "TempatLahir1",
					Filename:     "Filename1",
				},
			}
		})

		When("parameters given are invalid", func() {
			It("should return error", func() {
				mock.ExpectExec(query).WithArgs(mockOrangTua.Username, mockOrangTua.Email, mockOrangTua.Password, mockOrangTua.Nama, mockOrangTua.JenisKelamin, mockOrangTua.NoHP, mockOrangTua.Alamat, mockOrangTua.Siswa.ID).WillReturnError(errors.New("invalid parameters"))
				err := orangTuaRepository.Create(mockOrangTua)
				Expect(err).To(HaveOccurred())
			})
		})

		When("parameters given are valid", func() {
			It("should return no error", func() {
				mock.ExpectExec(query).WithArgs(mockOrangTua.Username, mockOrangTua.Email, mockOrangTua.Password, mockOrangTua.Nama, mockOrangTua.JenisKelamin, mockOrangTua.NoHP, mockOrangTua.Alamat, mockOrangTua.Siswa.ID).WillReturnError(nil)
				err := orangTuaRepository.Create(mockOrangTua)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
