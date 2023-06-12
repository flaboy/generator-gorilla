package controllers

import (
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type ApiController struct {
	gormc.Controller
}

type ApiResult struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func (c ApiController) Render(data interface{}) revel.Result {
	return c.RenderJSON(&ApiResult{
		Code:  0,
		Data:  data,
		Error: "",
	})
}

func (c ApiController) RenderError(data interface{}) revel.Result {
	return c.RenderJSON(&ApiResult{
		Code:  1,
		Data:  nil,
		Error: data.(string),
	})
}
