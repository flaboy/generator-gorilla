package controllers

import (
	"<%= config.name %>/app/models"
	"github.com/revel/revel"
)

type Migrate struct {
	Controller
}

func (c Migrate) Index() revel.Result {
	return c.Render()
}

func (c Migrate) Update() revel.Result {
	c.CreateTables()
	c.CreateAdmin("admin", "admin")
	return c.Redirect(Auth.Login)
}

func (c Migrate) CreateTables() revel.Result {
	c.DB.CreateTable(&models.User{})
	c.Flash.Success("Tables created")
	return c.Redirect(Migrate.Index)
}

func (c Migrate) CreateAdmin(username, password string) revel.Result {
	user := &models.User{}
	c.DB.First(user, "name = ?", username)
	if user.ID != 0 {
		c.Flash.Error("Admin user already exists")
		return c.Redirect(Migrate.Index)
	}

	user.Name = username
	user.SetPassword(password)
	user.IsAdmin = true
	c.DB.Save(user)

	c.Flash.Success("Admin user created")
	return c.Redirect(Migrate.Index)
}
