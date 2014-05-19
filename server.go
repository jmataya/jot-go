package main

import (
	"github.com/jmataya/jot-go/controllers"
	"github.com/jmataya/jot-go/restful"
)

func main() {
	notes := new(controllers.NotesController)
	var api = new(restful.API)
	api.RegisterRestfulController("/notes/{id}", notes)
	api.Start(3000)
}
