package models

import (
	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
	"goravel/database/factories"
)

type User struct {
	orm.Model
	Name   string
	Avatar string
	orm.SoftDeletes
}

func (u *User) Factory() factory.Factory {
	return &factories.UserFactory{}
}
