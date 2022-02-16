package roomhandlers

import (
	"encoding/json"
	"fmt"
	"hotel/domain/model"
	"hotel/internal/store"
	"hotel/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UpdateRoom(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}

		hotel, _ := s.Hotel().FindByID(req.HotelID)

		roomDTO, _ := s.Room().FindByID(req.RoomID)

		room, err := s.RoomRepository.RoomFromDTO(roomDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find room. Err msg:%v.", err)
			return
		}

		if req.RoomNumber != 0 {
			room.RoomNumber = req.RoomNumber
		}

		if req.PetType != "" {
			room.PetType = req.PetType
		}

		if hotel != nil {
			if hotel.HotelID != req.HotelID {
				room.Hotel = *hotel
			}
		}

		if req.RoomPhotoURL != "" {
			room.RoomPhotoURL = req.RoomPhotoURL
		}

		err = room.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.Room().Update(room)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't update Room. Err msg:%v.", err)
			return
		}

		s.Logger.Info("Update room with id = %d", room.RoomID)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Update room with id = %d", room.RoomID)})

	}
}
