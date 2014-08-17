package person

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Site     string        `bson:"site"`
	CheckIn  time.Time     `bson:"check-in"`
	CheckOut time.Time     `bson:"check-out"`
}

type Repository interface {
	Insert(p Person) error
	Update(p Person) (Person, error)
}

type MongoRepository struct {
	c *mgo.Collection
}

func NewMongoRepository(c *mgo.Collection) Repository {
	return &MongoRepository{c}
}

func (repository *MongoRepository) Insert(p Person) error {
	return repository.c.Insert(&p)
}

func (repository *MongoRepository) Update(p Person) (Person, error) {
	var newP Person

	_, err := repository.c.FindId(p.ID).Apply(mgo.Change{
		Update: &p,
	}, &newP)

	if err != nil {
		return Person{}, err
	}

	return newP, nil
}
