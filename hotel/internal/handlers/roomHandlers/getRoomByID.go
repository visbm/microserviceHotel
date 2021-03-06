package roomhandlers

import (
	"encoding/json"
	"fmt"
	"hotel/internal/apperror"
	"hotel/internal/store"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetRoomByID ...
func GetRoomByID(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))
			json.NewEncoder(w).Encode(apperror.NewAppError(fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id")), fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Bad request. Err msg:%v. Requests body: %v", err, ps.ByName("id"))))
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't open DB. Err msg:%v.", err)))
			return
		}

		room, err := s.Room().FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't find room.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't find room. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(room)
	}
}
