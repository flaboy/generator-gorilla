package controllers

import (
	"github.com/revel/config"
	"github.com/revel/revel"
	"os"
	"path"
	"strconv"
	"time"
)

var menus []Menu

func init() {
	revel.InterceptFunc(checkUser, revel.BEFORE, &Controller{})
	revel.InterceptFunc(adminUI, revel.AFTER, &Controller{})
	revel.InterceptFunc(doNothing, revel.BEFORE, &Setup{})
	menus = []Menu{}
	revel.OnAppStart(LoadSitemaps)
}

func doNothing(c *revel.Controller) revel.Result { return nil }

func checkUser(c *revel.Controller) revel.Result {
	if c.Name == "Setup" || c.Name == "Auth" {
		return nil
	}
	if c.Session["user"] == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(Auth.Login)
	} else {
		c.ViewArgs["user"] = c.Session["user"]
	}
	return nil
}

type Menu struct {
	Section string
	Name    string
	Route   string
	Icon    string
	Active  bool
}

func LoadSitemaps() {
	for _, dirname := range revel.ConfPaths {
		entries, err := os.ReadDir(path.Join(dirname, "sitemap"))
		if err == nil {
			for _, entry := range entries {
				c, err := config.ReadDefault(path.Join(dirname, "sitemap", entry.Name()))
				if err == nil {
					for _, sec := range c.Sections() {
						menu_name, err1 := c.String(sec, "name")
						menu_route, err2 := c.String(sec, "route")
						menu_icon, _ := c.String(sec, "name")
						if err1 == nil && err2 == nil {
							menu := Menu{
								Section: sec,
								Name:    menu_name,
								Route:   menu_route,
								Icon:    menu_icon,
							}
							menus = append(menus, menu)
						}
					}
				}
			}
		}
	}
}

func adminUI(c *revel.Controller) revel.Result {
	if c.Name == "Setup" || c.Name == "Auth" {
		return nil
	}

	c.ViewArgs["admin_ui"] = true
	c.ViewArgs["admin_in_section"] = c.Name
	c.ViewArgs["admin_menus"] = menus
	return nil
}

func toInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func toInt64(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	return 0
}

func toFloat32(s string) float32 {
	if f, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(f)
	}
	return 0.0
}

func toFloat64(s string) float64 {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	return 0.0
}

func toBool(s string) bool {
	if b, err := strconv.ParseBool(s); err == nil {
		return b
	}
	return false
}

func toTime(s string) time.Time {
	if t, err := time.Parse("2006-01-02 15:04:05", s); err == nil {
		return t
	}
	return time.Time{}
}
