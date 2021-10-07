package gpc

import "github.com/oit-sec-lab/gpc-checker-server/src/domain/entities/gpc"

type IGpcRepository interface {
	CheckGPC(string) (gpc.Gpc, error)
}

