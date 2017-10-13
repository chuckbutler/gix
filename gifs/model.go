package gifs

import "gopkg.in/mgo.v2/bson"

// Gif represnts an image w/ meta

type Gif struct {
	ID   bson.ObjectId `bson:"_id"`
	Url  string        `json:"Url"`
	Tags string        `json:"tags"`
}

// Gifs is an array of Gif

type Gifs []Gif
