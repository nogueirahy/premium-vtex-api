package internal

import (
	"net/http"
)

func InitializeServer() *http.ServeMux {
	deps := Setup()
	app := AuthorizedRoutes(deps)

	return app
}
