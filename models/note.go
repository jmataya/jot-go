package models

import (
	"encoding/json"
	"fmt"

	"github.com/jmataya/jot-go/config"
	"labix.org/v2/mgo/bson"
)

type Note struct {
	// MongoModel `bson:"-"`
	Id    bson.ObjectId "_id,omitempty"
	Title string        "title"
}

func NewNote() *Note {
	n := new(Note)
	// n.Initialize(n)
	return n
}

func (m *Note) Store() ([]byte, error) {
	// if m.model == nil {
	// 	return nil, errors.New("Model not initialized")
	// } else if !m.Id.Valid() {
	// 	m.Id = bson.NewObjectId()
	// }

	m.Id = bson.NewObjectId()
	content, _ := json.Marshal(m)
	fmt.Println(content)

	driver := new(config.MongoDriver)
	driver.SetConnect("mongodb://localhost/jot-go")
	driver.SetDatabase("test")
	driver.SetCollection("foo")

	err := driver.Insert(m)

	return content, err
}
