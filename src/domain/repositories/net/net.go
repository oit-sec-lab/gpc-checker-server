package net

type INetRepository interface {
	CheckGPC(string) (bool, error)
}
