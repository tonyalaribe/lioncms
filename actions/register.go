package actions

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
	client "github.com/ponzu-cms/go-client"
)

func RegisterPage(c buffalo.Context) error {
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("register.html"))
}

func Register2Handler(c buffalo.Context) error {
	// configure the http client
	name := c.Request().FormValue("name")
	club := c.Request().FormValue("club")
	region := c.Request().FormValue("region")
	district := c.Request().FormValue("district")
	modeOfPayment := c.Request().FormValue("mode_of_payment")

	// configure the http client
	cfg := client.Config{
		Host:         "http://localhost:8080",
		DisableCache: false, // defaults to false, here for documentation
	}

	// add custom header(s) if needed:
	// cfg.Header.Set("Authorization", "Bearer $ACCESS_TOKEN")
	// cfg.Header.Set("X-Client", "MyGoApp v0.9")
	if name != "" && club != "" && region != "" && district != "" && modeOfPayment != "" {
		ponzu := client.New(cfg)

		RegisterID := RandStringBytes(6)

		// create Content item of type with data
		data := make(url.Values)

		data.Set("name", name)
		data.Set("club", club)
		data.Set("region", region)
		data.Set("district", district)
		data.Set("payment_mode", modeOfPayment)
		data.Set("register_id", RegisterID)

		// nil indicates no data params are filepaths,
		// otherwise would be a []string of key names that are filepaths (docs coming)
		resp, err := ponzu.Create("RegisteredUser", data, nil)
		if err != nil {
			fmt.Println("Create:Blog error:", err)
			return errors.WithStack(err)
		}

		fmt.Println(resp.Data[0]["status"], resp.Data[0]["id"])
		// id := int(resp.Data[0]["id"].(float64))
		log.Printf("resp: %#v", resp.Data[0])
		if resp.Data[0]["status"] == "pending" {
			c.Set("HasRegisterError", false)
			c.Set("RegisterSuccess", "You have successfully been registered.")
			c.Set("RegisterID", RegisterID)
			return c.Render(200, publicR.HTML("register.html"))
		}
	}
	c.Set("HasRegisterError", true)
	c.Set("RegisterError", "Please fill out all fields")
	return c.Render(200, publicR.HTML("register.html"))
}

/// STILL NEED TO TEST THIS OUT...WANNA KNOW WHY ITS NOT WORKING...
func RegisterHandler(c buffalo.Context) error {
	c.Set("HasRegisterError", true)
	// configure the http client

	name := c.Request().FormValue("name")
	club := c.Request().FormValue("club")
	region := c.Request().FormValue("region")
	district := c.Request().FormValue("district")
	modeOfPayment := c.Request().FormValue("mode_of_payment")

	// get registrants :: http://localhost:8080/api/contents?type=RegisteredUser
	// add a registereduser = http://localhost:8080/api/content/create?type=RegisteredUser
	if name != "" && club != "" && region != "" && district != "" && modeOfPayment != "" {
		// cfg := client.Config{
		// 	Host:         "http://localhost:8080",
		// 	DisableCache: false, // defaults to false, here for documentation
		// }

		// // add custom header(s) if needed:
		// cfg.Header.Set("Authorization", "Bearer $ACCESS_TOKEN")
		// cfg.Header.Set("X-Client", "MyGoApp v0.9")
		// pon := client.New(cfg)
		postURL := "http://localhost:8080/api/content/create?type=RegisteredUser"
		// registeredUser := content.RegisteredUser{
		// 	Name:        name,
		// 	Club:        club,
		// 	Region:      region,
		// 	District:    district,
		// 	PaymentMode: modeOfPayment,
		// }
		// ru, err := json.Marshal(registeredUser)
		// if err != nil {
		// 	log.Println("Error json.Marshal: ", err)
		// 	c.Set("RegisterError", "Invalid Input values.")
		// 	return c.Render(200, publicR.HTML("register.html"))
		// }
		// log.Println("RU: ", string(ru))
		// data := make(url.Values)
		// data.Set("name", name)
		// data.Set("club", club)
		// data.Set("region", region)
		// data.Set("district", district)
		// data.Set("payment_mode", modeOfPayment)
		bufferBody := &bytes.Buffer{}
		writer := multipart.NewWriter(bufferBody)
		writer.WriteField("name", name)
		writer.WriteField("club", club)
		writer.WriteField("region", region)
		writer.WriteField("district", district)
		writer.WriteField("payment_mode", modeOfPayment)

		req, err := http.NewRequest("POST", postURL, bytes.NewReader(bufferBody.Bytes()))
		if err != nil {
			log.Println("request creation error: ", err)
		}
		_ = writer.Close()
		// req.Header.Set("Content-Type", "multipart/form-data; boundary=gc0p4Jq0M2Yt08jU534c0p")
		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := &http.Client{}
		resp, err := client.Do(req)

		// resp, err := pon.Create("RegisteredUser", data, nil)
		// log.Println(resp.Data[0])
		if err != nil {
			// log.Println("Create error : ", err)
			log.Println("Response(inside error): ", resp)
			fmt.Println("Create:Blog error:", err)
			c.Set("RegisterError", "Could not register you at the moment. Please try again.")
			return c.Render(200, publicR.HTML("register.html"))
		}
		// defer resp.Body.Close()

		if resp.StatusCode == 200 {
			c.Set("HasRegisterError", false)
			c.Set("RegisterSuccess", "You have successfully been registered.")
		}
		log.Println("response: ", resp, "\n\n request: ", req.Body)
		body, _ := ioutil.ReadAll(resp.Body)
		if len(body) == 0 {
			log.Println("response: ", resp)
			c.Set("RegisterError", "Could not register you at the moment. Please try again.")
		}
		log.Println("Response body: ", body)

		// fmt.Println(resp.Data[0]["status"], resp.Data[0]["id"])
		// id := int(resp.Data[0]["id"].(float64))
		// log.Println("ID: ", id)
	}
	// c.Set("RegisterError", "Please fill out all fields")
	// publicR.HTMLLayout = "index.html"
	// using the github.com/gobuffalo/plush for templating...
	return c.Render(200, publicR.HTML("register.html"))
}
