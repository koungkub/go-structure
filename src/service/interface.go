package service

type User interface {
	GetName() string
	SetName(string)
	Clear()
}
