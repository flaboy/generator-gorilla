package controllers

import (
	"github.com/revel/revel"
)

type HelloApi struct {
	ApiController
}

func (c HelloApi) SayHello() revel.Result {
	return c.Render("world")
}
