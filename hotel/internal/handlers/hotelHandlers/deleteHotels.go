package hotelhandlers

import (
	"encoding/json"
	"fmt"
	"hotel/internal/store"
	"hotel/pkg/response"
	"log"
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
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			return
		}
		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}
		err = s.Hotel().Delete(id)
		if err != nil {
			log.Print(err)
			s.Logger.Errorf("Can't delete hotel. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Delete hotel with id = %d", id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Delete hotel with id = %d", id)})

	}
}