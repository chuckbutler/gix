package gifs

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repository ...
type Repository struct{}

// SERVER the DB Server
const SERVER = "localhost:27017"

// DBName the name of the DB instance
const DBNAME = "gix"

// DOCNAME the name of the document
const DOCNAME = "gifs"

//GetGifs returns the list of Gifs
func (r Repository) GetGifs() Gifs {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to the Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Gifs{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddGif insters a Gif in the DB
func (r Repository) AddGif(gif Gif) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	gif.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(gif)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateGif updates a Gif ini the DB (not used for now)
func (r Repository) UpdateGif(gif Gif) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	gif.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(gif.ID, gif)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteGif deletes a gif (not used for now)
func (r Repository) DeleteGif(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}

	// Grab Id
	oid := bson.ObjectIdHex(id)

	// Remove Gif
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	return "OK"
}
