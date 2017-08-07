// Package home displays the Home page.
package home

import (
	"net/http"

	"github.com/arapov/pile/lib/flight"
	"github.com/arapov/pile/model/ldap"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Get("/", Index)
}

// Index displays the home page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	result, err := ldap.Get(c.LDAP)
	if err != nil {
		c.FlashError(err)
		return
	}

	v := c.View.New("home/index")
	if c.Sess.Values["id"] != nil {
		v.Vars["first_name"] = c.Sess.Values["first_name"]
	}
	v.Vars["result"] = result

	v.Render(w, r)
}
