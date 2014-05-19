package main

import (
	"fmt"
	"os"

	"labix.org/v2/mgo/bson"

	"github.com/jmataya/jot-go/config"
	"github.com/jmataya/jot-go/models"
)

func main() {
	driver := new(config.MongoDriver)
	driver.SetConnect("mongodb://localhost/jot-go")
	driver.SetDatabase("test")
	driver.SetCollection("foo")

	doc := models.Note{Id: bson.NewObjectId(), Title: "Hello from go"}
	err := driver.Insert(doc)
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		os.Exit(1)
	}

	n := models.NewNote()
	n.Title = "Test"
	content, err := n.Store()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%s\n", content)
	}

	// 	var updatednote models.Note
	// 	err = session.DB("test").C("foo").Find(bson.M{}).One(&updatednote)
	// 	if err != nil {
	// 		fmt.Printf("git an error when finding a doc %v\n", err)
	// 		os.Exit(1)
	// 	}
	//
	// 	fmt.Printf("Found document: %+v\n", updatednote)

	// notes := new(controllers.NotesController)
	// var api = new(restful.API)
	// api.RegisterRestfulController("/notes/{id}", notes)
	// api.Start(3000)
}
