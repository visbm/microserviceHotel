package server

import (
	hotelhandlers "hotel/hotel/internal/handlers/hotelHandlers"
	roomhandlers "hotel/hotel/internal/handlers/roomHandlers"
	seathandlers "hotel/hotel/internal/handlers/seatHandlers"
	"hotel/hotel/internal/store"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {

	s.Router.Handle("GET", "/admin/hotels", hotelhandlers.AllHotelsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/admin/hotels/id", hotelhandlers.GetHotelByID(store.New(s.Config)))
	s.Router.Handle("POST", "/admin/hotels/delete", hotelhandlers.DeleteHotels(store.New(s.Config)))

	s.Router.Handle("GET", "/admin/rooms", roomhandlers.AllRoomsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/admin/rooms/id", roomhandlers.GetRoomByID(store.New(s.Config)))
	s.Router.Handle("POST", "/admin/rooms/delete", roomhandlers.DeleteRooms(store.New(s.Config)))

	s.Router.Handle("GET", "/admin/seats", seathandlers.AllSeatsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/admin/seats/id", seathandlers.GetSeatByID(store.New(s.Config)))
	s.Router.Handle("POST", "/admin/seats/delete", seathandlers.DeleteSeats(store.New(s.Config)))

}
