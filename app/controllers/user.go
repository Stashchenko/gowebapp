package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type User struct {
	BaseController
}

type Stuff struct {
	Foo string `json:"foo" xml:"foo"`
	Bar int    `json:"bar" xml:"bar"`
}

func (c User) Index() revel.Result {
	data := make(map[string]interface{})
	data["error"] = nil
	stuff := Stuff{Foo: "xyz", Bar: 999}
	data["stuff"] = stuff
	format := c.GetParam("format", "json")
	fmt.Printf("=============================format = %v\n", format)
	if format == "xml" {
		return c.RenderXML(data)
	}
	return c.RenderJSON(data)
}
