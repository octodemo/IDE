package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	methods "github.com/thestrukture/IDE/api/methods"
)

func GETApiPkgBugs(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	bugs := methods.GetLogs(r.FormValue("pkg"))
	sapp := methods.GetApp(methods.GetApps(), r.FormValue("pkg"))
	if len(bugs) == 0 || sapp.Passed {
		response = "{}"
	} else {
		response = mResponse(bugs[0])
	}

	callmet = true
	return
}