// This file was generated by Gopher Sauce.
// UPDATE IT TO BETTER TEST YOUR CODE.
package handlers

import (
	"testing"
	"github.com/cheikhshift/gos/onyx"
)

func TestPOSTApiDockerfile(t *testing.T) {


}

func TestPOSTApiComposer(t *testing.T) {


}

func TestApiSocket(t *testing.T) {

	
}

func TestGETApiPkgBugs(t *testing.T) {


}

func TestGETApiKanban(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequest("GET", "/api/kanban")

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
	
}

func TestPOSTApiGit(t *testing.T) {


}

func TestPOSTApiKanban(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/kanban", `{}`)

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)

}

func TestGETApiEmpty(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequest("GET", "/api/empty")

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestPOSTApiTester(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/tester/", `pkg=test&mode=foo&c=x`)

	onyx.Header(req, "Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)

}

func TestPOSTApiCreate(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/create", ``)

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestPOSTApiDelete(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/delete", ``)

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestPOSTApiRename(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/rename", ``)

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestPOSTApiNew(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/new", ``)

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestPOSTApiAct(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/act", ``)

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestPOSTApiPut(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/put", ``)

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestGETApiBuild(t *testing.T) {

	
}

func TestGETApiStart(t *testing.T) {


}

func TestGETApiStop(t *testing.T) {


}

func TestGETApiBin(t *testing.T) {


}

func TestGETApiExport(t *testing.T) {


}

func TestApiComplete(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/complete", ``)

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestPOSTApiConsole(t *testing.T) {

	h := MakeHandler(Handler)
	
	req,err := onyx.NewRequestWithBody("POST", "/api/console", "command=ls")

	onyx.Header(req, "Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		panic(err)
	}

	onyx.Handle(req, h, t)
}

func TestApiTerminal_realtime(t *testing.T) {


}
