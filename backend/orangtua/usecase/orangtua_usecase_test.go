package usecase_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/domain/mocks"
	"github.com/rg-km/final-project-engineering-4/backend/orangtua/usecase"
)

var _ = Describe("OrangtuaUsecase", func() {
	var (
		ctrl                   *gomock.Controller
		mockSiswaRepository    *mocks.MockSiswaRepository
		mockOrangTuaRepository *mocks.MockOrangTuaRepository
		orangTuaUseCase        domain.OrangTuaUseCase
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockSiswaRepository = mocks.NewMockSiswaRepository(ctrl)
		mockOrangTuaRepository = mocks.NewMockOrangTuaRepository(ctrl)
		orangTuaUseCase = usecase.NewOrangTuaUseCase(mockOrangTuaRepository, mockSiswaRepository)
	})

	AfterEach(func() {
		defer ctrl.Finish()
	})

	Describe("register orangtua account", func() {
		var (
			mockOrangTua domain.OrangTua
			mockSiswa    domain.Siswa
		)

		BeforeEach(func() {
			mockSiswa = domain.Siswa{
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
			}
			mockOrangTua = domain.OrangTua{
				Username: "Username1",
				Email:    "Email1",
				Password: "Password1",
				Nama:     "Nama1",
				NoHP:     "NoHP1",
				Alamat:   "Alamat1",
				Siswa:    domain.Siswa{ID: 1},
			}
		})

		When("username given is used", func() {
			It("should return error", func() {
				mockOrangTuaRepository.EXPECT().GetByUsername(gomock.Any()).Return(&mockOrangTua, nil)
				err := orangTuaUseCase.Register(mockOrangTua)
				Expect(err).To(HaveOccurred())
			})
		})

		When("email given is used", func() {
			It("should return error", func() {
				mockOrangTuaRepository.EXPECT().GetByUsername(gomock.Any()).Return(nil, errors.New("error"))
				mockOrangTuaRepository.EXPECT().GetByEmail(gomock.Any()).Return(&mockOrangTua, nil)
				err := orangTuaUseCase.Register(mockOrangTua)
				Expect(err).To(HaveOccurred())
			})
		})

		When("email siswa given is not exist", func() {
			It("should return error", func() {
				mockOrangTuaRepository.EXPECT().GetByUsername(gomock.Any()).Return(nil, errors.New("error"))
				mockOrangTuaRepository.EXPECT().GetByEmail(gomock.Any()).Return(nil, errors.New("error"))
				mockSiswaRepository.EXPECT().GetByEmail(gomock.Any()).Return(nil, errors.New("error"))
				err := orangTuaUseCase.Register(mockOrangTua)
				Expect(err).To(HaveOccurred())
			})
		})

		When("register success", func() {
			It("should return no error", func() {
				mockOrangTuaRepository.EXPECT().GetByUsername(gomock.Any()).Return(nil, errors.New("not found"))
				mockOrangTuaRepository.EXPECT().GetByEmail(gomock.Any()).Return(nil, errors.New("not found"))
				mockSiswaRepository.EXPECT().GetByEmail(gomock.Any()).Return(&mockSiswa, nil)
				mockOrangTuaRepository.EXPECT().Create(gomock.Any()).Return(nil)
				err := orangTuaUseCase.Register(mockOrangTua)
				Expect(err).To(BeNil())
			})
		})
	})
})
