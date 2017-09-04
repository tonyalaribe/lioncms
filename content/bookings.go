package content

import (
	"fmt"
	"net/http"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Bookings struct {
	item.Item

	Name           string `json:"name"`
	Email          string `json:"email"`
	RegistrationId string `json:"registration_id"`
	ModeOfPayment  string `json:"mode_of_payment"`
	Accomodation   string `json:"accomodation"`
}

// MarshalEditor writes a buffer of html to edit a Bookings within the CMS
// and implements editor.Editable
func (b *Bookings) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(b,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Bookings field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", b, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Email", b, map[string]string{
				"label":       "Email",
				"type":        "text",
				"placeholder": "Enter the Email here",
			}),
		},
		editor.Field{
			View: editor.Input("RegistrationId", b, map[string]string{
				"label":       "RegistrationId",
				"type":        "text",
				"placeholder": "Enter the RegistrationId here",
			}),
		},
		editor.Field{
			View: editor.Input("ModeOfPayment", b, map[string]string{
				"label":       "ModeOfPayment",
				"type":        "text",
				"placeholder": "Enter the ModeOfPayment here",
			}),
		},
		editor.Field{
			View: editor.Input("Accomodation", b, map[string]string{
				"label":       "Accomodation",
				"type":        "text",
				"placeholder": "Enter the Accomodation here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Bookings editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Bookings"] = func() interface{} { return new(Bookings) }
}

func (b *Bookings) AutoApprove(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (b *Bookings) Create(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (b *Bookings) String() string {
	return b.Email
}
