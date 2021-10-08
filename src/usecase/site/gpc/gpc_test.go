package gpc

import (
	"github.com/golang/mock/gomock"
	"github.com/oit-sec-lab/gpc-checker-server/src/domain/entities/gpc"
	"github.com/oit-sec-lab/gpc-checker-server/src/domain/entities/site"
	mockGpc "github.com/oit-sec-lab/gpc-checker-server/src/mock/gpc"
	"testing"
)

const (
	TestURL = "https://www.example.com"
)

func TestCheckGpc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGpcRepository := mockGpc.NewMockIGpcRepository(ctrl)
	interactor := NewGpcInteractor(mockGpcRepository)

	t.Run("gpc true", func(t *testing.T) {
		gpc := gpc.NewGpc(true)
		site, _ := site.NewSite(1, TestURL, gpc)
		mockGpcRepository.EXPECT().CheckGPC(TestURL).Return(gpc, nil)
		_, e := interactor.CheckGPC(site.URL())
		if e != nil {
			t.Fatalf("unexpected error: %v\n", e)
		}
	})

	t.Run("gpc false", func(t *testing.T) {
		gpc := gpc.NewGpc(false)
		mockGpcRepository.EXPECT().CheckGPC(TestURL).Return(gpc, nil)
		_, e := interactor.CheckGPC(TestURL)
		if e != nil {
			t.Fatalf("unexpected error: %v\n", e)
		}
	})
}
