package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jmataya/jot-go/config"

	"labix.org/v2/mgo/bson"
)

type GoModel interface {
	Save() (json []byte, err error)
}

type MongoModel struct {
	Id    bson.ObjectId "_id,omitempty"
	model GoModel
}

func (m *MongoModel) Initialize(model GoModel) {
	m.model = model
}

func (m *MongoModel) Save() ([]byte, error) {
	if m.model == nil {
		return nil, errors.New("Model not initialized")
	} else if !m.Id.Valid() {
		m.Id = bson.NewObjectId()
	}

	content, _ := json.Marshal(m.model)
	fmt.Println(content)

	driver := new(config.MongoDriver)
	driver.SetConnect("mongodb://localhost/jot-go")
	driver.SetDatabase("test")
	driver.SetCollection("foo")

	err := driver.Insert(m.model)

	return content, err
}
