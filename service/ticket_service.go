package service

import (
	"MyProjects/Zomentum/constants"
	"MyProjects/Zomentum/dao"
	"MyProjects/Zomentum/models"
	"net/http"
)

type TicketService struct {
}

//BookTicket recieves request from handler and Perform business Logic for booking a ticket.
func (t *TicketService) BookTicket(u1 models.User, time string, seat_no string) (error, *string) {
	err, tId := dao.BookTicket(u1, time, seat_no)
	return err, tId
}

//UpdateTicketTiming recieves request from handler and Perform business Logic for Updating
//a ticket time.
func (t *TicketService) UpdateTicketTiming(tId string, time string) error {
	err := dao.UpdateTicketTiming(tId, time)
	return err
}

//ViewAllTickets recieves request from handler and Perform business Logic for viewing
// all tickets at a particular time.
func (t *TicketService) ViewAllTickets(time string) []models.Ticket {
	tickets := dao.ViewAllTicketsAttime(time)
	return tickets
}

//DeleteTicket recieves request from handler and Perform business Logic to delete
// a ticket corresponding to tickedId
func (t *TicketService) DeleteTicket(ticketId string) error {
	err := dao.DeleteTicket(ticketId)
	return err
}

//ViewUSerDetailsByTicketID recieves request from handler and Perform business Logic to view
// User details of the person who holds right of that ticket.
func (t *TicketService) ViewUSerDetailsByTicketID(ticketId string) models.GenericResponse {
	u, err := dao.ViewUserDetails(ticketId)
	gr := models.GenericResponse{
		StatusCode: http.StatusInternalServerError,
		// Data:       "",
	}
	if err != nil {
		errMsg := err.Error()
		gr.Message = errMsg
		if errMsg == constants.INVALID_TICKET {
			gr.StatusCode = http.StatusBadRequest
		} else if errMsg == constants.NO_USER_FOUND {
			gr.StatusCode = http.StatusNotFound
		}

	} else {
		gr.StatusCode = http.StatusOK
		gr.Data = u
		gr.Message = constants.SUCCESS
	}

	return gr
}

//CheckExpiryAndDeleteTickets recieves request from handler and Perform business Logic for
//checking and deleting expired tickets.
func (t *TicketService) CheckExpiryAndDeleteTickets() {
	dao.CheckExpiryAndDeleteTickets()
}
