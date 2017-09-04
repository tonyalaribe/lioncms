package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/gobuffalo/buffalo"
	client "github.com/ponzu-cms/go-client"
	"github.com/tonyalaribe/lion2018/content"
	"github.com/tonyalaribe/lion2018/models"
)

var cfg = client.Config{
	Host:         "http://localhost:8080",
	DisableCache: false, // defaults to false, here for documentation
}

func AccomodationHandler(c buffalo.Context) error {
	// get list of accomodations based on the hotel slug
	slug := c.Param("slug")
	hotel, err := models.GetHotelBySlug(slug)
	if err != nil && err.Error() != "Not Found" {
		c.Set("HasAccomodationError", true)
		c.Set("AccomodationError", "No accomodations for this hotel")
		return c.Render(200, publicR.HTML("accomodations.html"))
	}
	log.Println("Hotel: ", hotel)
	// hotel.ID

	// configure the http client
	ponzu := client.New(cfg)
	// create list to hold accomodations for this hotel...
	var accomodations []content.Accomodation
	// loop through each accomodation of the hotel...
	for _, v := range hotel.Accomodation {
		s := strings.Split(v, "=")
		AccomodationID, err := strconv.Atoi(s[len(s)-1])
		if err != nil {
			continue
		}
		log.Printf("ID: %#v\n\n", AccomodationID)
		resp, err := ponzu.Content("Accomodation", AccomodationID)
		if err != nil {
			log.Printf("ponzu.Get(%s) Error: %#v", v, err.Error())
			continue
		}
		// log.Println("Response body: ", resp.Data)
		acc := struct {
			Data []content.Accomodation
		}{}
		err = json.Unmarshal(resp.JSON, &acc)
		if err != nil {
			log.Println("Error: ", err)
			continue
		}
		log.Printf("\nAccomodation: %#v\n", string(resp.JSON))
		accomodations = append(accomodations, acc.Data[0])
	}
	c.Set("accomodations", accomodations)
	for _, v := range accomodations {
		log.Printf("\n\nAccomodation: %#v\n\n", v)
	}
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("accomodations.html"))
}

func BookAccomodationPage(c buffalo.Context) error {
	accomodationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error: %#v", err)
	}
	c.Set("index", accomodationID)

	ponzu := client.New(cfg)
	resp, err := ponzu.Content("Accomodation", accomodationID)
	if err != nil {
		log.Printf("ponzu.Get(%#v) Error: %#v", accomodationID, err.Error())
		publicR.HTMLLayout = "main.html"
		return c.Render(200, publicR.HTML("book.html"))
	}

	acc := struct {
		Data []content.Accomodation
	}{}

	err = json.Unmarshal(resp.JSON, &acc)
	if err != nil {
		log.Println("Error: ", err)
		publicR.HTMLLayout = "main.html"
		return c.Render(200, publicR.HTML("book.html"))
	}
	log.Printf("\nAccomodation: %#v\n", string(resp.JSON))
	c.Set("accomodation", acc.Data[0])
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("book.html"))
}

func BookAccomodation(c buffalo.Context) error {
	accomodationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error: %#v", err)
	}
	c.Set("index", accomodationID)

	ponzu := client.New(cfg)
	resp, err := ponzu.Content("Accomodation", accomodationID)
	if err != nil {
		log.Printf("ponzu.Get(%#v) Error: %#v", accomodationID, err.Error())
		publicR.HTMLLayout = "main.html"
		return c.Render(200, publicR.HTML("book.html"))
	}

	acc := struct {
		Data []content.Accomodation
	}{}

	err = json.Unmarshal(resp.JSON, &acc)
	if err != nil {
		log.Println("Error: ", err)
		publicR.HTMLLayout = "main.html"
		return c.Render(200, publicR.HTML("book.html"))
	}
	log.Printf("\nAccomodation: %#v\n", string(resp.JSON))
	accomodation := acc.Data[0]

	name := c.Request().FormValue("name")
	email := c.Request().FormValue("email")
	registrationNumber := c.Request().FormValue("registration_number")
	ModeOfPayment := c.Request().FormValue("mode_of_payment")

	// check if registration number is registered or valid.
	regSearch, err := ponzu.Search("RegisteredUser", registrationNumber)
	if err != nil {
		log.Println("SEARCH REGISTERED USERS ERROR : ", err)
	}
	log.Println("regSearch: ", regSearch.Data)
	// HAVING PROBLEMS WITH SEARCH... //
	// check if email already exist in the bookings...//

	data := make(url.Values)

	data.Set("name", name)
	data.Set("email", email)
	data.Set("registration_id", registrationNumber)
	data.Set("mode_of_payment", ModeOfPayment)
	data.Set("accomodation", "?type=Accomodation&id="+strconv.Itoa(accomodationID))

	bookResponse, err := ponzu.Create("Bookings", data, nil)
	if err != nil {
		fmt.Println("Create:Bookings error: ", err)
		c.Set("HasBookError", true)
		c.Set("BookError", "Could not book this hotel at this time, please try again.")
		publicR.HTMLLayout = "main.html"
		return c.Render(200, publicR.HTML("book.html"))
	}

	log.Printf("resp: %#v", bookResponse.Data[0])
	if bookResponse.Data[0]["status"] != "pending" && bookResponse.Data[0]["status"] != "" {
		c.Set("BookSuccess", "You have successfully booked this accomodation.")
		// reduce the quantity of accomdation...
		accomodation.Quantity = accomodation.Quantity - 1
		UpdateData := make(url.Values)
		UpdateData.Set("quantity", strconv.Itoa(accomodation.Quantity))
		updateResp, err := ponzu.Update("Accomodation", accomodation.ID, UpdateData, nil)
		if err != nil {
			log.Println("Error: ", err)
		}
		log.Println("updateResp: ", updateResp)

	}
	c.Set("accomodation", accomodation)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("book.html"))
}
