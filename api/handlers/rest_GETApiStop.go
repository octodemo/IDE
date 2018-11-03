package handlers

import (
	"net/http"
	"os"

	"github.com/cheikhshift/gos/core"
	"github.com/gorilla/sessions"
	templates "github.com/thestrukture/IDE/api/templates"

	methods "github.com/thestrukture/IDE/api/methods"

	types "github.com/thestrukture/IDE/types"
)

func GETApiStop(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	gp := os.ExpandEnv("$GOPATH")

	os.Chdir(gp + "/src/" + r.FormValue("pkg"))
	apps := methods.GetApps()
	sapp := methods.GetApp(apps, r.FormValue("pkg"))

	if sapp.Pid == "" {
		response = templates.Alert(types.Alertbs{Type: "danger", Text: "No build running."})
	} else {
		core.RunCmdB("kill -3 " + sapp.Pid)
		response = templates.Alert(types.Alertbs{Type: "success", Text: "Build stopped."})
	}

	//DebugLogs.Insert(dObj)
	apps = methods.UpdateApp(apps, r.FormValue("pkg"), sapp)
	methods.SaveApps(apps)

	//Users.Update(bson.M{"uid":me.UID}, me)

	callmet = true
	return
}