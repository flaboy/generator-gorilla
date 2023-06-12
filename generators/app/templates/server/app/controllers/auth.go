package controllers

import (
	"<%= config.name %>/app/models"
	"github.com/revel/revel"
)

type Auth struct {
	Controller
}

func (c Auth) Login() revel.Result {
	return c.Render()
}

func (c Auth) LoginPost(username, password string) revel.Result {
	user := &models.User{}
	c.DB.First(user, "name = ?", username)
	if user.ID != 0 && user.CheckPassword(password) {
		c.Session["user"] = username
		return c.Redirect(App.Index)
	}

	c.Flash.Error("Login failed")
	return c.Redirect(Auth.Login)
}

func (c Auth) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(Auth.Login)
}
