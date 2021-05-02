package admin

import (
	"net/http"

	"github.com/serveye/backend-tech/pkg/render"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "./templates/dashboard")
}
