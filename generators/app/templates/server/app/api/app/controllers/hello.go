package controllers

import (
	"github.com/revel/revel"
)

type Hello struct {
	ApiController
}

func (c Hello) SayHello() revel.Result {
	return c.Render("world")
}
