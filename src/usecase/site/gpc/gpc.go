package gpc

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
	gpcRepository "github.com/oit-sec-lab/dnt-verify-server/src/domain/repositories/gpc"
)

type GpcInteractor struct {
	gpcRepository gpcRepository.IGpcRepository
}

func NewGpcInteractor(gr gpcRepository.IGpcRepository) GpcInteractor {
	return GpcInteractor{gr}
}

func (interactor *GpcInteractor) CheckGPC(u string) (gpc gpc.Gpc, err error) {
		return interactor.gpcRepository.CheckGPC(u)
}