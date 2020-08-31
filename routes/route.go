package routes

import (
	"MyProjects/Zomentum/handler"
	"github.com/gorilla/mux"
)

var AllRoutes *mux.Router = mux.NewRouter()

//InititateRoutes is creating all the routes for the server.
func InitiateRoutes() {
	AllRoutes.Methods("POST").Path("/bookticket").
		Name("BookTicket").HandlerFunc(handler.BookTicketHandler)

	AllRoutes.Methods("PUT").Path("/updatetiming").
		Name("UpdateTicketTiming").HandlerFunc(handler.UpdateTicketTimingHandler)

	AllRoutes.Methods("GET").Path("/viewalltickets").
		Name("ViewAllTickets").HandlerFunc(handler.ViewAllTicketsHandler)

	AllRoutes.Methods("DELETE").Path("/Deleteticket").
		Name("DeleteTicket").HandlerFunc(handler.DeleteTicketHandler)

	AllRoutes.Methods("GET").Path("/viewuserdetails").
		Name("ViewUserDetails").HandlerFunc(handler.ViewUSerDetailsByTicketIDHandler)
}
