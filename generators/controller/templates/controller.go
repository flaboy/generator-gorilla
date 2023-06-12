package controllers

import (
	"fmt"
	"<%= config.name %>/app/models"
	"github.com/revel/revel"
)

type <%- entity.name %> struct {
	Controller
}

func (c <%- entity.name %>) Index() revel.Result {
	var <%- entity.lname %>_list []models.<%- entity.name %>
	c.DB.Find(&<%- entity.lname %>_list)
	return c.Render(<%- entity.lname %>_list)
}

func (c <%- entity.name %>) Detail(id int) revel.Result {
	var <%- entity.lname %> models.<%- entity.name %>
	c.DB.First(&<%- entity.lname %>, id)
	if <%- entity.lname %>.ID == 0 {
		return c.NotFound(fmt.Sprintf("<%- entity.name %> %d does not exist", id))
	}
	return c.Render(<%- entity.lname %>)
}

func (c <%- entity.name %>) Edit(id int) revel.Result {
	var <%- entity.lname %> models.<%- entity.name %>
	c.DB.First(&<%- entity.lname %>, id)
	if <%- entity.lname %>.ID == 0 {
		return c.NotFound(fmt.Sprintf("<%- entity.name %> %d does not exist", id))
	}
	return c.Render(<%- entity.lname %>)
}

func (c <%- entity.name %>) Save(id int) revel.Result {
	var <%- entity.lname %> models.<%- entity.name %>
	c.DB.First(&<%- entity.lname %>, id)
	if <%- entity.lname %>.ID == 0 {
		return c.NotFound(fmt.Sprintf("<%- entity.lname %> %d does not exist", id))
	}
<% entity.body.forEach(function(field){ if(!field.skipInForm){-%>
<% if(field.typeChanger){-%>
	<%- entity.lname %>.<%- field.Uname %>  = <%- field.typeChanger %>(c.Params.Form.Get("<%- field.columnName %>"))
<% } else {-%>
	<%- entity.lname %>.<%- field.Uname %>  = c.Params.Form.Get("<%- field.columnName %>")
<% } -%>
<% }}); -%>
	c.DB.Save(&<%- entity.lname %>)
	return c.Redirect(<%- entity.name %>.Detail, <%- entity.lname %>.ID)
}

func (c <%- entity.name %>) Add() revel.Result {
	return c.Render()
}

func (c <%- entity.name %>) AddPost() revel.Result {
	<%- entity.lname %> := &models.<%- entity.name %>{}
<% entity.body.forEach(function(field){ if(!field.skipInForm){-%>
<% if(field.typeChanger){-%>
	<%- entity.lname %>.<%- field.Uname %>  = <%- field.typeChanger %>(c.Params.Form.Get("<%- field.columnName %>"))
<% } else {-%>
	<%- entity.lname %>.<%- field.Uname %>  = c.Params.Form.Get("<%- field.columnName %>")
<% } -%>
<% }}); -%>
	c.DB.Create(&<%- entity.lname %>)
	return c.Redirect(<%- entity.name %>.Detail, <%- entity.lname %>.ID)
}

func (c <%- entity.name %>) Delete(id int) revel.Result {
	var <%- entity.lname %> models.<%- entity.name %>
	c.DB.First(&<%- entity.lname %>, id)
	if <%- entity.lname %>.ID == 0 {
		return c.NotFound(fmt.Sprintf("<%- entity.lname %> %d does not exist", id))
	}
	c.DB.Delete(&<%- entity.lname %>)
	return c.RenderJSON(true)
}
