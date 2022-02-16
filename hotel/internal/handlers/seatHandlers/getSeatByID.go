package seathandlers

import (
	"encoding/json"
	"hotel/internal/store"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetSeatByID ...
func GetSeatByID(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}

		seat, err := s.Seat().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find seat. Err msg:%v.", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(seat)
	}
}