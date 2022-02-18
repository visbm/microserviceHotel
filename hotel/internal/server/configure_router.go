package server

import (
	hotelhandlers "hotel/internal/handlers/hotelHandlers"
	roomhandlers "hotel/internal/handlers/roomHandlers"
	seathandlers "hotel/internal/handlers/seatHandlers"
	"hotel/internal/store"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {

	s.Router.Handle("GET", "/hotel/hotels", hotelhandlers.AllHotelsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/hotel/hotels/id", hotelhandlers.GetHotelByID(store.New(s.Config)))
	s.Router.Handle("POST", "/hotel/hotels/delete", hotelhandlers.DeleteHotels(store.New(s.Config)))

	s.Router.Handle("GET", "/hotel/rooms", roomhandlers.AllRoomsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/hotel/rooms/id", roomhandlers.GetRoomByID(store.New(s.Config)))
	s.Router.Handle("POST", "/hotel/rooms/delete", roomhandlers.DeleteRooms(store.New(s.Config)))

	s.Router.Handle("GET", "/hotel/seats", seathandlers.AllSeatsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/hotel/seats/id", seathandlers.GetSeatByID(store.New(s.Config)))
	s.Router.Handle("POST", "/hotel/seats/delete", seathandlers.DeleteSeats(store.New(s.Config)))

}
