package main

import (
	"net/url"

	"github.com/jmataya/jot-go/restful"
)

type Notes struct {
	restful.StandardRestfulType
}

func (Notes) List(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "LIST for Notes"}
	return 200, data
}

func (Notes) Show(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "SHOW for Notes"}
	return 200, data
}

func (Notes) Create(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "CREATE for Notes"}
	return 200, data
}

func (Notes) Update(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "UPDATE for Notes"}
	return 200, data
}

func (Notes) Destroy(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "DESTROY for Notes"}
	return 200, data
}

func main() {
	notes := new(Notes)
	notes.SetBasePath("/notes/{id}")
	var api = new(restful.API)
	api.RegisterRestfulType("Note", notes)
	api.Start(3000)
}
