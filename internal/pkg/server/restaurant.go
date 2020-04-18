package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	"github.com/mdcurran/palatable/internal/pkg/db"
)

func (s *Server) getRestaurant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    name := p.ByName("name")
    if name == "" {
        Error(w, errors.New("no name query parameter provided"), http.StatusBadRequest)
        return
    }

    restaurant, err := db.GetRestaurant(name)
    if err == gorm.ErrRecordNotFound {
        Error(w, err, http.StatusNotFound)
        return
    }
    if err != nil {
        Error(w, err, http.StatusInternalServerError)
        return
    }

    s.encodeJSON(w, &restaurant)
}

func (s *Server) postRestaurant(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        Error(w, err, http.StatusBadRequest)
        return
    }

    var restaurant db.Restaurant
    err = json.Unmarshal(b, &restaurant)
    if err != nil {
        Error(w, err, http.StatusUnprocessableEntity)
        return
    }

    err = db.PostRestaurant(&restaurant)
    if err != nil {
        Error(w, err, http.StatusInternalServerError)
        return
    }

    s.encodeJSON(w, &restaurant)
}
