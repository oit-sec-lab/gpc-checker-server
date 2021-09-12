package gpc

type Gpc struct {
	enable bool
}
type Gpcs []Gpc

func NewGpc(e bool) Gpc {
	return Gpc{enable: e}
}

func (g Gpc) Enable() bool {
	return g.enable
}
