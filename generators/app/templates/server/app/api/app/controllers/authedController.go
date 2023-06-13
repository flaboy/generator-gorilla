package controllers

import (
	"<%- config.name %>/app/models"
	"reflect"
	"strings"

	"github.com/revel/revel"
)

type authModel interface {
	RetrieveByAccessToken(token string) error
}

type AuthedController struct {
	ApiController
	auth authModel
}

func init() {
	revel.InterceptFunc(checkAuth, revel.BEFORE, &AuthedController{})
}

func checkAuth(c *revel.Controller) revel.Result {
	controller := reflect.ValueOf(c.AppController).Elem().FieldByName("AuthedController").Interface().(AuthedController)

	controller.auth = &models.Login{}
	if c.Request.Header.Get("Authorization") == "" {
		return controller.RenderError("Authorization required", 401)
	}
	bearer_token := c.Request.Header.Get("Authorization")
	token := strings.TrimSpace(bearer_token[7:])

	err := controller.auth.RetrieveByAccessToken(token)
	if err != nil {
		return controller.RenderError("Invalid AccessToken", 401)
	}
	return nil
}
