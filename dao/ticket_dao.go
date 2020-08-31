package dao

import (
	"MyProjects/Zomentum/constants"
	"MyProjects/Zomentum/models"
	"errors"
	"fmt"
)

var Userdata map[string]models.User = map[string]models.User{}
var TicketData map[string]models.Ticket = map[string]models.Ticket{}

// InitiateDatabase initiates the database and make it ready-to-use for server.
func InitiateDatabase() {
	t1 := models.Ticket{
		true,
		100,
		"001",
		"10:00", //"31.08.2020@10:00"
		true,
		"1234567890",
	}
	t2 := models.Ticket{
		true,
		200,
		"002",
		"10:00",
		false,
		"9876543210",
	}
	// fmt.Println(t1, t2)
	m1 := map[string]models.Ticket{}
	m1[t1.ID] = t1
	m2 := map[string]models.Ticket{}
	m2[t2.ID] = t2
	u1 := models.User{
		"Rohit",
		"Khatri",
		"1234567890",
		m1,
	}
	u2 := models.User{
		"Sunil",
		"Khatri",
		"9876543210",
		m2,
	}
	Userdata[u1.PhoneNum] = u1
	Userdata[u2.PhoneNum] = u2
	TicketData[t1.ID] = t1
	TicketData[t2.ID] = t2
}

//ViewAllTicketsAttime recieves request from service and reads from Database and returns
// slice of tickets corresponding to a particular time
func ViewAllTicketsAttime(time string) []models.Ticket {
	tickets := []models.Ticket{}
	for _, tkt := range TicketData {
		if tkt.Timing == time {
			tickets = append(tickets, tkt)
		}
	}
	return tickets
}

//ViewUserDetails recieves request from service and reads from Database and returns a
// pointer to the user who holds rights to a particular ticket_Id
func ViewUserDetails(ticketId string) (*models.User, error) {
	if v, ok := TicketData[ticketId]; ok {
		if ud, ok := Userdata[v.UserPhoneNum]; ok {
			return &ud, nil
		}
		err := errors.New(constants.NO_USER_FOUND)
		return nil, err
	}

	fmt.Println(constants.INVALID_TICKET)
	err := errors.New(constants.INVALID_TICKET)
	return nil, err
}

//DeleteTicket recieves request from service and deletes a ticket from Database if the ticket
// with requested tId is available.
func DeleteTicket(tId string) error {
	if v, ok := TicketData[tId]; ok {
		delete(Userdata[v.UserPhoneNum].Tktstatus, tId)
		delete(TicketData, tId)
		return nil
	}
	return errors.New(constants.INVALID_TICKET)
}

//UpdateTicketTiming recieves request from service and modifies the Database for the requested
// tId with new timing.
func UpdateTicketTiming(tId string, time string) error {
	if v, ok := TicketData[tId]; ok {
		v.Timing = time
		TicketData[tId] = v
		return nil
	}
	return errors.New(constants.INVALID_TICKET)
}

// CountOfTickets returns the count Of total Tickets booked for a show time.
func CountOfTickets(time string) int8 {
	count := int8(0)
	for _, v := range TicketData {
		if v.Timing == time {
			count++
		}
	}
	return (count)
}

//BookTicket recieves request from service and creates newly booked tickets
//to Database.
func BookTicket(user models.User, time string, seat_no string) (error, *string) {
	if CountOfTickets(time) > 19 {
		return errors.New(constants.MAX_TICKET_CAPACITY_REACHED), nil
	}
	if v, ok := Userdata[user.PhoneNum]; ok {
		ticket := models.Ticket{
			IsBooked:     true,
			Price:        100,
			ID:           "00" + seat_no,
			Timing:       time,
			Expired:      false,
			UserPhoneNum: user.PhoneNum,
		}
		if _, ok = TicketData[ticket.ID]; ok && TicketData[ticket.ID].Timing == time {
			return errors.New(constants.SEAT_ALREADY_BOOKED), nil
		}
		TicketData[ticket.ID] = ticket
		v.Tktstatus[ticket.ID] = ticket
		Userdata[user.PhoneNum] = v
		return nil, &ticket.ID
	}
	return errors.New(constants.NO_USER_FOUND_), nil
}

//CheckExpiryAndDeleteTickets recieves request from service and deletes the expired
// tickets from the Database
func CheckExpiryAndDeleteTickets() {
	for key, val := range TicketData {
		if val.Expired {
			delete(Userdata[val.UserPhoneNum].Tktstatus, key)
			fmt.Println("Deleted Ticket that had ticket Id", key, "as it was expired.")
			delete(TicketData, key)
		}
	}
}
