package main

import (
	"net/url"

	"github.com/jmataya/jot/restful"
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

func main() {
	resource := new(NoteResource)

	var api = new(restful.API)
	api.AddResource(resource, "/notes")
	api.Start(3000)
}
