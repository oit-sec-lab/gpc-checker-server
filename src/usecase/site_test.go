package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities"
	mockSite "github.com/oit-sec-lab/dnt-verify-server/src/mock"
	"testing"
)

const (
	TestURL = "https://www.example.com"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSiteRepository := mockSite.NewMockISiteRepository(ctrl)
	usecase := NewSiteInteractor(mockSiteRepository)

	mockSiteRepository.EXPECT().Store(gomock.Any()).Return(nil)

	site, _ := entities.NewSite(1, TestURL, true)
	err := usecase.Add(site)
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}
}

func TestFindByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSiteRepository := mockSite.NewMockISiteRepository(ctrl)
	usecase := NewSiteInteractor(mockSiteRepository)

	t.Run("success", func(t *testing.T) {
		site, _ := entities.NewSite(1, TestURL, true)
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
	usecase := NewSiteInteractor(mockSiteRepository)

	t.Run("success", func(t *testing.T) {

		site, _ := entities.NewSite(1, TestURL, true)
		mockSiteRepository.EXPECT().Exists(TestURL).Return(true, nil)

		mockSiteRepository.EXPECT().FindByURL(TestURL).Return(site, nil)
		_, e := usecase.VerifyGPC(site.URL())
		if e != nil {
			t.Fatalf("unexpected error: %v\n", e)
		}
	})

	t.Run("url not found", func(t *testing.T) {
		mockSiteRepository.EXPECT().Exists(TestURL).Return(false, nil)
		mockSiteRepository.EXPECT().CheckGPC(TestURL).Return(true, nil)

		_, e := usecase.VerifyGPC(TestURL)
		if e != nil {
			t.Fatalf("unexpected error: %v\n", e)
		}
	})
}
