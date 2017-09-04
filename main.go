package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markbates/going/defaults"
	_ "github.com/ponzu-cms/ponzu/content"
	"github.com/tonyalaribe/lion2018/actions"
	"github.com/tonyalaribe/lion2018/models"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func main() {
	port := defaults.String(os.Getenv("PORT"), "3002")
	// baseURL := defaults.String(os.Getenv("CMS_URL"), "https://cms.gophercon.com")
	baseURL := defaults.String(os.Getenv("CMS_URL"), "http://localhost:8080")
	models.BaseURL = baseURL
	log.Printf("Starting gcon on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}
