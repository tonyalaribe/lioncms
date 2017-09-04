package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/tonyalaribe/lion2018/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	hp, err := models.GetHomePage()
	if err != nil && err.Error() != "Not Found" {
		log.Println(err)
		return c.Error(500, err)
	}
	// log.Println("after get homepage")
	// log.Printf("%#v", hp)
	// log.Println("---------")
	// log.Printf("%#v", hp.Sponsors)
	// log.Println("_____")
	// log.Printf("%#v", hp.Sponsors.Gold)
	c.Set("homepage", hp)
	c.Set("sponsors", hp.Sponsors)
	c.Set("baseurl", models.BaseURL)
	// publicR.HTMLLayout = "mainhome.html"
	// publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("mainhome.html"))
	// return c.Render(200, publicR.HTML("index.html"))
}

// AboutHandler is a default handler to serve up
// content pages
func AboutHandler(c buffalo.Context) error {
	slug := c.Param("slug")
	p, err := models.GetPageBySlug(slug)
	if err != nil {
		return c.Error(404, err)
	}

	c.Set("page", p)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("page.html"))
}

// EventHandler is a default handler to serve up
// event pages
func EventHandler(c buffalo.Context) error {
	slug := c.Param("slug")
	p, err := models.GetEventBySlug(slug)
	if err != nil {
		return c.Error(404, err)
	}

	c.Set("event", p)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("event.html"))
}
