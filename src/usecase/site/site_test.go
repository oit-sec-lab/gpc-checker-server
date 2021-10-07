package site

import (
	"github.com/golang/mock/gomock"
	"github.com/oit-sec-lab/gpc-checker-server/src/domain/entities/gpc"
	gpcUsecase "github.com/oit-sec-lab/gpc-checker-server/src/usecase/site/gpc"
	"github.com/oit-sec-lab/gpc-checker-server/src/domain/entities/site"
	mockGpc "github.com/oit-sec-lab/gpc-checker-server/src/mock/gpc"
	mockSite "github.com/oit-sec-lab/gpc-checker-server/src/mock/site"
	"testing"
)

const (
	TestURL = "https://www.example.com"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSiteRepository := mockSite.NewMockISiteRepository(ctrl)
	mockGpcRepository := mockGpc.NewMockIGpcRepository(ctrl)
	gpcInteractor := gpcUsecase.NewGpcInteractor(mockGpcRepository)
	usecase := NewSiteInteractor(mockSiteRepository, gpcInteractor)

	mockSiteRepository.EXPECT().Store(gomock.Any()).Return(nil)

	site, _ := site.NewSite(1, TestURL, gpc.NewGpc(true))
	err := usecase.Add(site)
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}
}

func TestFindByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSiteRepository := mockSite.NewMockISiteRepository(ctrl)
	mockGpcRepository := mockGpc.NewMockIGpcRepository(ctrl)
	gpcInteractor := gpcUsecase.NewGpcInteractor(mockGpcRepository)
	usecase := NewSiteInteractor(mockSiteRepository, gpcInteractor)

	t.Run("success", func(t *testing.T) {
		site, _ := site.NewSite(1, TestURL, gpc.NewGpc(true))
		mockSiteRepository.EXPECT().Exists(TestURL).Return(true, nil)
		mockSiteRepository.EXPECT().FindByURL(TestURL).Return(site, nil)

		s, e := usecase.FindByURL(TestURL)
		if e != nil {
			t.Fatalf("unexpected error: %v\n", e)
		}
		if s != site {
			t.Fatalf("expected %v, but got %v", site, s)
		}
	})

	t.Run("url not found", func(t *testing.T) {
		mockSiteRepository.EXPECT().Exists(TestURL).Return(false, nil)

		_, e := usecase.FindByURL(TestURL)
		if e.Error() != URLNotFound {
			t.Fatalf("expected %v, but got %v", URLNotFound, e)
		}
	})

}

func TestVerifyGPC(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSiteRepository := mockSite.NewMockISiteRepository(ctrl)
	mockGpcRepository := mockGpc.NewMockIGpcRepository(ctrl)
	gpcInteractor := gpcUsecase.NewGpcInteractor(mockGpcRepository)
	usecase := NewSiteInteractor(mockSiteRepository, gpcInteractor)

	t.Run("success", func(t *testing.T) {

		site, _ := site.NewSite(1, TestURL, gpc.NewGpc(true))
		mockSiteRepository.EXPECT().Exists(TestURL).Return(true, nil)

		mockSiteRepository.EXPECT().FindByURL(TestURL).Return(site, nil)
		_, e := usecase.VerifyGPC(site.URL())
		if e != nil {
			t.Fatalf("unexpected error: %v\n", e)
		}
	})

	t.Run("url not found", func(t *testing.T) {
		mockSiteRepository.EXPECT().Exists(TestURL).Return(false, nil)
		mockGpcRepository.EXPECT().CheckGPC(TestURL).Return(gpc.NewGpc(true), nil)

		_, e := usecase.VerifyGPC(TestURL)
		if e != nil {
			t.Fatalf("unexpected error: %v\n", e)
		}
	})
}
