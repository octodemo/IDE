// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/api/assets"
	"github.com/thestrukture/IDE/types"
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"log"

	"github.com/fatih/color"
)

//
// Renders HTML of template
// Modal with struct types.SModal
func Modal(d types.SModal) string {
	return netbModal(d)
}

// recovery function used to log a
// panic.
func templateFNModal(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/modal) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDModal = "tmpl/ui/modal.tmpl"

// Render template with JSON string as
// data.
func netModal(args ...interface{}) string {

	localid := templateIDModal
	var d *types.SModal
	defer templateFNModal(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.SModal{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Modal")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}

// alias of template render function.
func bModal(d types.SModal) string {
	return netbModal(d)
}

//

// template render function
func netbModal(d types.SModal) string {
	localid := templateIDModal
	defer templateFNModal(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Modal")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = types.SModal{}
	output.Reset()
	output = nil
	return outpescaped
}

// Unmarshal a json string to the template's struct
// type
func netcModal(args ...interface{}) (d types.SModal) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.SModal{}
	}
	return
}

// Create a struct variable of template.
func cModal(args ...interface{}) (d types.SModal) {
	if len(args) > 0 {
		d = netcModal(args[0])
	} else {
		d = netcModal()
	}
	return
}
