package content

import (
	"fmt"
	"net/http"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type RegisteredUser struct {
	item.Item

	Name        string `json:"name"`
	Club        string `json:"club"`
	Region      string `json:"region"`
	District    string `json:"district"`
	PaymentMode string `json:"payment_mode"`
	RegisterID  string `json:"register_id"`
}

// MarshalEditor writes a buffer of html to edit a RegisteredUser within the CMS
// and implements editor.Editable
func (r *RegisteredUser) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(r,
		// Take note that the first argument to these Input-like functions
		// is the string version of each RegisteredUser field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", r, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Club", r, map[string]string{
				"label":       "Club",
				"type":        "text",
				"placeholder": "Enter the Club here",
			}),
		},
		editor.Field{
			View: editor.Input("Region", r, map[string]string{
				"label":       "Region",
				"type":        "text",
				"placeholder": "Enter the Region here",
			}),
		},
		editor.Field{
			View: editor.Input("District", r, map[string]string{
				"label":       "District",
				"type":        "text",
				"placeholder": "Enter the District here",
			}),
		},
		editor.Field{
			View: editor.Input("PaymentMode", r, map[string]string{
				"label":       "PaymentMode",
				"type":        "text",
				"placeholder": "Enter the PaymentMode here",
			}),
		},
		editor.Field{
			View: editor.Input("RegisterID", r, map[string]string{
				"label":       "RegisterID",
				"type":        "text",
				"placeholder": "Enter the RegisterID here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render RegisteredUser editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["RegisteredUser"] = func() interface{} { return new(RegisteredUser) }
}

func (r *RegisteredUser) String() string {
	return r.Name
}

func (r *RegisteredUser) Create(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (r *RegisteredUser) Approve(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (r *RegisteredUser) IndexContent() bool {
	return true
}
