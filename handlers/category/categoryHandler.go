package category

import (
	"fmt"
	"github.com/serveye/backend-tech/render"

	"net/http"
)

type category struct {
	name string
	id string
	is_deleted bool
}
 var categories = [] category {
	 {
		 name:       "education",
		 id:         "xyz1",

		 is_deleted: false,
	 },
	 {
		 name:       "eletrician",
		 id:         "xyz2",
		 is_deleted: false,
	 },
	 {
		 name:       "home services",
		 id:         "xyz3",
		 is_deleted: false,
	 },
 }
func GetCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("came to handler")
	render.RenderTemplate(w, "categoriesDatatable.page.tmpl")
}

