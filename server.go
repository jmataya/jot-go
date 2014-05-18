package main

import (
	"net/url"

	"github.com/jmataya/jot-go/restful"
)

type NoteResource struct {
	restful.PostNotSupported
	restful.PutNotSupported
	restful.DeleteNotSupported
}

func (NoteResource) Get(values url.Values) (int, interface{}) {
	data := map[string]string{"hello": "world"}
	return 200, data
}

type GlobalResource struct {
	restful.PostNotSupported
	restful.PutNotSupported
	restful.DeleteNotSupported
}

func (GlobalResource) Get(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "always works"}
	return 200, data
}

func main() {
	noteResource := new(NoteResource)
	globalResource := new(GlobalResource)

	var api = new(restful.API)
	api.AddResource(noteResource, "/notes")
	api.AddResource(globalResource, "/")
	api.Start(3000)
}
