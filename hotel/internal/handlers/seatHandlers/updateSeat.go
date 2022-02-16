package seathandlers

import (
	"encoding/json"
	"fmt"
	"hotel/domain/model"
	"hotel/internal/store"
	"hotel/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UpdateSeat(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		req := &model.SeatDTO{}
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

		roomDTO, _ := s.Room().FindByID(req.RoomID)

		room, _ := s.Room().RoomFromDTO(roomDTO)

		SeatDTO, err := s.Seat().FindByID(req.SeatID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find Room. Err msg:%v.", err)
			return
		}

		seat, err := s.Seat().SeatFromDTO(SeatDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find room. Err msg:%v.", err)
			return
		}

		if room != nil {
			if seat.Room.RoomID != req.RoomID {
				seat.Room = *room
			}
		}

		if req.Description != "" {
			seat.Description = req.Description
		}

		if !req.RentFrom.IsZero() {
			seat.RentFrom = req.RentFrom
		}

		if !req.RentTo.IsZero() {
			seat.RentTo = req.RentTo
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
