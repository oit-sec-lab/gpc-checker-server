package gpc

type Gpc struct {
	enable bool
}

func NewGpc(e bool) Gpc {
	return Gpc{enable: e}
}

func (g Gpc) Enable() bool {
	return g.enable
}
