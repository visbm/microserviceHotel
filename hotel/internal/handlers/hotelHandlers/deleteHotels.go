package hotelhandlers

import (
	"encoding/json"
	"fmt"
	"hotel/internal/apperror"
	"hotel/internal/store"
	"hotel/pkg/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteHotel ...
func DeleteHotels(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
		err = s.Hotel().Delete(id)
		if err != nil {
			s.Logger.Errorf("Can't delete hotel. Err msg:%v.", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't delete hotel.", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't delete hotel. Err msg:%v.", err)))
			return
		}

		s.Logger.Info("Delete hotel with id = %d", id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Delete hotel with id = %d", id)})

	}
}
