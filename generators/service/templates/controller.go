package controllers

import (
	"<%= config.name %>/app/models"
	"github.com/revel/revel"
)

type <%- entity.name %>Api struct {
	AuthedController // changeto "ApiController" to open access
}

func (c <%- entity.name %>Api) Index() revel.Result {
	var <%- entity.lname %>_list []models.<%- entity.name %>
	c.DB.Find(&<%- entity.lname %>_list)
	return c.Render(<%- entity.lname %>_list)
}

func (c <%- entity.name %>Api) Detail(id int) revel.Result {
	var <%- entity.lname %> models.<%- entity.name %>
	c.DB.First(&<%- entity.lname %>, id)
	if <%- entity.lname %>.ID == 0 {
		return c.NotFound("<%- entity.name %> does not exist")
	}
	return c.Render(<%- entity.lname %>)
}

func (c <%- entity.name %>Api) Save(id int) revel.Result {
	var <%- entity.lname %> models.<%- entity.name %>
	c.DB.First(&<%- entity.lname %>, id)
	if <%- entity.lname %>.ID == 0 {
		return c.NotFound("<%- entity.lname %> does not exist")
	}
	c.Params.BindJSON(&<%- entity.lname %>)
	c.DB.Save(&<%- entity.lname %>)
	return c.Render(<%- entity.lname %>)
}

func (c <%- entity.name %>Api) Add() revel.Result {
	<%- entity.lname %> := &models.<%- entity.name %>{}
	c.Params.BindJSON(&<%- entity.lname %>)
	c.DB.Create(&<%- entity.lname %>)
	return c.Render(<%- entity.lname %>)
}

func (c <%- entity.name %>Api) Delete(id int) revel.Result {
	var <%- entity.lname %> models.<%- entity.name %>
	c.DB.First(&<%- entity.lname %>, id)
	if <%- entity.lname %>.ID == 0 {
		return c.NotFound("<%- entity.lname %> does not exist")
	}
	c.DB.Delete(&<%- entity.lname %>)
	return c.Render("deleted")
}
