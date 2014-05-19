package config

import "labix.org/v2/mgo"

type MongoDriver struct {
	session    *mgo.Session
	database   *mgo.Database
	collection *mgo.Collection

	connectionStr string
	databaseStr   string
	collectionStr string
}

func (m *MongoDriver) connect() (*mgo.Session, *mgo.Collection, error) {
	sess, err := mgo.Dial(m.connectionStr)
	if err != nil {
		return nil, nil, err
	}

	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB(m.databaseStr).C(m.collectionStr)
	return sess, collection, err
}

func (m *MongoDriver) SetConnect(path string) {
	m.connectionStr = path
}

func (m *MongoDriver) SetDatabase(dbName string) {
	m.databaseStr = dbName
}

func (m *MongoDriver) SetCollection(collectionName string) {
	m.collectionStr = collectionName
}

func (m *MongoDriver) Insert(docs ...interface{}) error {
	session, collection, err := m.connect()
	if err != nil {
		return err
	} else {
		defer session.Close()
		err = collection.Insert(docs)
		return err
	}
}
