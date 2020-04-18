package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	"github.com/mdcurran/palatable/internal/pkg/db"
)

func (s *Server) listItems(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

    items, err := db.GetItems(restaurant.ID)
    if err != nil && err != gorm.ErrRecordNotFound {
        Error(w, err, http.StatusInternalServerError)
        return
    }

    s.encodeJSON(w, items)
}

func (s *Server) postItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    type new struct {
        Title       string
        Restaurant  string
        Description string
        Review      string
        Date        time.Time
    }

    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        Error(w, err, http.StatusBadRequest)
        return
    }

    var n new
    err = json.Unmarshal(b, &n)
    if err != nil {
        Error(w, err, http.StatusUnprocessableEntity)
        return
    }

    restaurant, err := db.GetRestaurant(n.Restaurant)
    if err == gorm.ErrRecordNotFound {
        Error(w, err, http.StatusNotFound)
        return
    }
    if err != nil {
        Error(w, err, http.StatusInternalServerError)
        return
    }

    var i db.Item
    i.RestaurantID = restaurant.ID
    i.Title = n.Title
    i.Description = n.Description
    i.Review = n.Review
    i.Date = n.Date

    err = db.PostItem(&i)
    if err != nil {
        Error(w, err, http.StatusInternalServerError)
        return
    }

    s.encodeJSON(w, &i)
}
