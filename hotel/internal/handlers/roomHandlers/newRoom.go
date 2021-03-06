package roomhandlers

import (
	"encoding/json"
	"fmt"
	"hotel/domain/model"
	"hotel/internal/apperror"
	"hotel/internal/store"
	"hotel/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRoom(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.RoomDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't open DB", fmt.Sprintf("%d", http.StatusInternalServerError), fmt.Sprintf("Can't open DB. Err msg:%v.", err)))
			return
		}

		hotel, err := s.Hotel().FindByID(req.HotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode(apperror.NewAppError("Cant find hotel.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Cant find hotel. Err msg:%v.", err)))
			return
		}

		room := model.Room{
			RoomID:       0,
			RoomNumber:   req.RoomNumber,
			PetType:      model.PetType(req.PetType),
			Hotel:        *hotel,
			RoomPhotoURL: req.RoomPhotoURL,
		}

		err = room.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			json.NewEncoder(w).Encode(apperror.NewAppError("Data is not valid.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Data is not valid. Err msg:%v.", err)))
			return
		}

		_, err = s.Room().Create(&room)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(apperror.NewAppError("Can't create room.", fmt.Sprintf("%d", http.StatusBadRequest), fmt.Sprintf("Can't create room. Err msg:%v.", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Creat room with id = %d", room.RoomID)})
	}
}
