package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"ticket-booking/constants"
	"ticket-booking/models"
	"ticket-booking/service"
)

// BookTicketHandler handles all the ticket Booking requests from clients.
// This writes response to the ResponseWriter object based on return
// value from service.
func BookTicketHandler(w http.ResponseWriter, r *http.Request) {
	ts := service.TicketService{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	UserAndTime := models.BookTicketRequest{}
	err = json.Unmarshal(body, &UserAndTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	User := UserAndTime.User
	time := UserAndTime.Timing
	SeatNo := UserAndTime.SeatNo
	err, tId := ts.BookTicket(User, time, SeatNo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Your Ticket is booked Succesfully with TicketId " + *tId))

}

// UpdateTicketTimingHandler handles updating all the ticket timing update requests from clients.
// This writes response to the ResponseWriter object based on return
// value from service.
func UpdateTicketTimingHandler(w http.ResponseWriter, r *http.Request) {
	ts := service.TicketService{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	tIdAndTime := models.UpdateTktTimeRequest{}
	err = json.Unmarshal(body, &tIdAndTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	time := tIdAndTime.Time
	tId := tIdAndTime.TicketID
	// fmt.Println(tId, time)
	err = ts.UpdateTicketTiming(tId, time)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(constants.SUCCESS))
}

// ViewAllTicketsHandler handles request for viewing Tickets at a particular time.
// This writes response to the ResponseWriter object based on return
// value from service.
func ViewAllTicketsHandler(w http.ResponseWriter, r *http.Request) {
	ts := service.TicketService{}
	time := r.FormValue("at_time")
	tickets := ts.ViewAllTickets(time)
	if len(tickets) == 0 {
		w.Write([]byte("No Tickets exists at time" + " " + time))
	} else {
		bs, err := json.Marshal(tickets)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(bs)
		}
	}
}

// DeleteTicketHandler handles request for deleting a Particular Ticket by taking ticket_id as query Parameter.
// This writes response to the ResponseWriter object based on return
// value from service.
func DeleteTicketHandler(w http.ResponseWriter, r *http.Request) {
	tId := r.FormValue("ticket_id")
	ts := service.TicketService{}
	err := ts.DeleteTicket(tId)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Ticket Deleted Successfully."))
	}
}

// ViewUSerDetailsByTicketIDHandler handles request for viewing User details based on ticket_id
// requested as query parameter from client.
// This writes response to the ResponseWriter object based on return
// value from service.
func ViewUSerDetailsByTicketIDHandler(w http.ResponseWriter, r *http.Request) {
	ts := service.TicketService{}
	tId := r.FormValue("ticket_id")

	gr := ts.ViewUSerDetailsByTicketID(tId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(gr.StatusCode)
	b, err := json.Marshal(gr.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(constants.MARSHAL_ERROR))
		return
	}
	w.Write(b)
}
