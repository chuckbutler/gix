package gifs

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Controller ...
type Controller struct {
	Repository Repository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	gifs := c.Repository.GetGifs()
	log.Println(gifs)
	data, _ := json.Marshal(gifs)
	w.Header().Set("ContentType", "applications.json; charset=UTF=8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddGif POST /
func (c *Controller) AddGif(w http.ResponseWriter, r *http.Request) {
	var gif Gif
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the request body
	if err != nil {
		log.Fatalln("Error AddGif", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddGif", err)
	}
	if err := json.Unmarshal(body, &gif); err != nil {
		// unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddGif unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.AddGif(gif)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applications/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

func (c *Controller) UpdateGif(w http.ResponseWriter, r *http.Request) {
	var gif Gif
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("error UpdateGif", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddUpdateGif", err)
	}
	if err := json.Unmarshal(body, &gif); err != nil {
		// unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateGif unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.UpdateGif(gif)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

func (c *Controller) DeleteGif(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] // param id

	if err := c.Repository.DeleteGif(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	return
}
