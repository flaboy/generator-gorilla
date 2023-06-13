package models

import "gorm.io/gorm"
<% for(var key in entity.imports){ %>
import "<%- key %>"
<% } %>
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
	<% if(field.comment){ %>//<% } %><%- field.Uname %>		<%- field.type %> <% if(field.columnName){ %>`json:"<%- field.columnName %>"`<% } %>
<% }); -%>
}

func (m *<%- entity.name %>) ToString() string {
	return "<%- entity.name %>"
}