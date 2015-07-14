package datastores

import (
	"commonobjects"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoStorer struct {
	name        string
	connection  *mgo.Session
	collection  string
	database    string
	creatorFunc commonobjects.StorableCreator
}

func NewMongoStorer(name, collection, connectionString, database string, creatorFunc commonobjects.StorableCreator) (*mongoStorer, error) {
	sess, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	mongostorer := &mongoStorer{name: name, collection: collection, connection: sess, database: database, creatorFunc: creatorFunc}
	return mongostorer, nil
}

func (ms *mongoStorer) GetName() string {
	return ms.name
}

func (ms *mongoStorer) Put(id string, item interface{}) error {
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	err := connCopy.DB(ms.database).C(ms.collection).Insert(item)
	if err != nil {
		return err
	}
	return nil
}

func (ms *mongoStorer) GetById(id string) (interface{}, error) {
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	object := ms.creatorFunc().(commonobjects.Storable)
	idkey := object.GetIdField()
	condition := bson.M{}
	condition[idkey] = id
	err := connCopy.DB(ms.database).C(ms.collection).Find(condition).One(object)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (ms *mongoStorer) Get(conditions interface{}) (interface{}, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (ms *mongoStorer) Delete(id string) error {
	return nil
}

func (ms *mongoStorer) GetList() (interface{}, error) {
	return nil, nil
}
