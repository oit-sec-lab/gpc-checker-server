package gpc

type Gpc struct {
	Enable bool	`json:"gpc"`
}
type Gpcs []Gpc

func NewGpc(e bool) Gpc {
	return Gpc{Enable: e}
}

