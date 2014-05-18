package main

import (
	"net/url"

	"github.com/jmataya/jot-go/restful"
)

type Notes struct{}

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

type Tasks struct{}

func (Tasks) List(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "LIST for Tasks"}
	return 200, data
}

func (Tasks) Show(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "SHOW for Tasks"}
	return 200, data
}

func (Tasks) Create(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "CREATE for Tasks"}
	return 200, data
}

func (Tasks) Update(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "UPDATE for Tasks"}
	return 200, data
}

func (Tasks) Destroy(values url.Values) (int, interface{}) {
	data := map[string]string{"message": "DESTROY for Tasks"}
	return 200, data
}

func main() {
	notes := new(Notes)
	tasks := new(Tasks)
	var api = new(restful.API)
	api.RegisterRestfulType("/notes/{id}", notes)
	api.RegisterRestfulType("/tasks/{id}", tasks)
	api.Start(3000)
}
