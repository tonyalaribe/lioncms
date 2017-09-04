package content

import (
	"fmt"
	"net/http"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Accomodation struct {
	item.Item

	Type        string `json:"type"`
	Hotel       string `json:"hotel"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	RoomPhoto   string `json:"room_photo"`
	
}

// MarshalEditor writes a buffer of html to edit a Accomodation within the CMS
// and implements editor.Editable
func (a *Accomodation) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Accomodation field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Type", a, map[string]string{
				"label":       "Type",
				"type":        "text",
				"placeholder": "Enter the Type here",
			}),
		},
		editor.Field{
			View: reference.Select("Hotel", a, map[string]string{
				"label": "Hotel",
			},
				"Hotel",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.Textarea("Description", a, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Input("Quantity", a, map[string]string{
				"label":       "Quantity",
				"type":        "text",
				"placeholder": "Enter the Quantity here",
			}),
		},
		editor.Field{
			View: editor.Input("Price", a, map[string]string{
				"label":       "Price",
				"type":        "text",
				"placeholder": "Enter the Price here",
			}),
		},
		editor.Field{
			View: editor.File("RoomPhoto", a, map[string]string{
				"label":       "Photo",
				"type":        "text",
				"placeholder": "Enter the Photo here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Accomodation editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Accomodation"] = func() interface{} { return new(Accomodation) }
}

// func SearchMapping() (*mapping.IndexMappingImpl, error) {

// }

func IndexContent() bool {
	return true
}
func (a *Accomodation) String() string {
	return a.Type
}

func (a *Accomodation) Update(res http.ResponseWriter, req *http.Request) error {
	return nil
}
