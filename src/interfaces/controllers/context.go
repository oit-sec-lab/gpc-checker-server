package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(interface{}, interface{})
}
