package controllers

import "net/url"

type NotesController struct{}

func (NotesController) List(values url.Values, params map[string]string) (int, interface{}) {
	data := map[string]string{"message": "LIST for Notes"}
	return 200, data
}

func (NotesController) Show(values url.Values, params map[string]string) (int, interface{}) {
	data := map[string]string{"id": params["id"]}
	return 200, data
}

func (NotesController) Create(values url.Values, params map[string]string) (int, interface{}) {
	data := map[string]string{"message": "CREATE for Notes"}
	return 200, data
}

func (NotesController) Update(values url.Values, params map[string]string) (int, interface{}) {
	data := map[string]string{"message": "UPDATE for Notes"}
	return 200, data
}

func (NotesController) Destroy(values url.Values, params map[string]string) (int, interface{}) {
	data := map[string]string{"message": "DESTROY for Notes"}
	return 200, data
}
