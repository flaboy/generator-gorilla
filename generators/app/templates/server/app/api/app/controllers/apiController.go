package controllers

import (
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type ApiController struct {
	gormc.Controller
}

func (c ApiController) Render(data interface{}) revel.Result {
	return c.RenderJSON(data)
}

func (c ApiController) RenderError(data interface{}, httpCode int) revel.Result {
	c.Response.Status = httpCode
	return c.RenderJSON(data)
}

func (c ApiController) NotFound(message string) revel.Result {
	return c.RenderError(message, 404)
}
