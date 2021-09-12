package gpc

import "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"

type IGpcRepository interface {
	CheckGPC(string) (gpc.Gpc, error)
}

