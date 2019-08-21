package service

import (
	"context"
	"reflect"
)

type Admin struct {
	Name string
}

func (admin *Admin) SetName(name string) {

	admin.Name = name
}

func (admin Admin) GetName(ctx context.Context) string {

	admin.Name = "eiei"
	return admin.Name
}

func (admin *Admin) Clear() {

	p := reflect.ValueOf(admin).Elem()
	p.Set(reflect.Zero(p.Type()))
}
