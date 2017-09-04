package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/tonyalaribe/lion2018/models"
	"strconv"
)

type SponsorResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *SponsorResource) List(c buffalo.Context) error {
	sl, err := models.GetSponsorList()
	if err != nil {
		return c.Error(500, err)
	}
	ssl := models.SortedSponsorList(sl)
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("sponsors", ssl)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("sponsors.html"))

}

// Show default implementation.
func (v *SponsorResource) Show(c buffalo.Context) error {
	idString :=c.Param("speaker_id")
	id , err := strconv.Atoi(idString)

	if err != nil {
		return c.Error(400, err)
	}
	p, err := models.GetSponsor(id)
	if err != nil {
		return c.Error(404, err)
	}

	c.Set("sponsor", p)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("sponsor.html"))
}
