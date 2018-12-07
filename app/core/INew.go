package core

type INew interface {
	NewFrom(option interface{}) interface{}
	New() interface{}
}
