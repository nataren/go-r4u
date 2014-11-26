package controllers 

import "github.com/revel/revel"

type User struct {
	*revel.Controller
}

type UserBE struct {
	FirstName string
	LastName string
}

func (u User) GetUsers() revel.Result {
	return u.RenderJson([]UserBE { UserBE { "césar", "lópez-natarén" } })
}