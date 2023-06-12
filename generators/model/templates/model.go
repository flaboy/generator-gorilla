package models

import "gorm.io/gorm"
<% entity.enums.forEach(function(item){ %>

type <%- item.name %> string

const (
<% item.values.forEach(function(value){ -%>
	<%- item.name %>_<%- value.key %> <%- item.name %> = "<%- value.key %>"
<% }); -%>)
<% }); %>
type <%- entity.name %> struct {
	gorm.Model
<% entity.body.forEach(function(field){ -%>
	<%- field.Uname %>		<%- field.type %> `json:"<%- field.name %>"`
<% }); -%>
}