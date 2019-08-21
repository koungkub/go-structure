package service

import "context"

type User interface {
	GetName(context.Context) string
	SetName(string)
	Clear()
}
