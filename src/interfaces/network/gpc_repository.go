package network

import (
	"encoding/json"
	"github.com/oit-sec-lab/gpc-checker-server/src/domain/entities/gpc"
)

type GpcRepository struct {
	HttpHandler
}

func (gr *GpcRepository) CheckGPC(url string) (gpc.Gpc, error) {
	r, e := gr.GET(url + "/.well-known/gpc.json")
	if e != nil {
		return gpc.Gpc{}, e
	}

	if r.StatusCode() == 200 {
		g := new(gpc.Gpc)
		err := json.Unmarshal([]byte(r.Body()), g)
		if err != nil {
			return gpc.Gpc{}, err
		}
		return *g, nil
	} else {
		return gpc.NewGpc(false), nil
	}
}
