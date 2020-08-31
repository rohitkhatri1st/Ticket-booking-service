package models

type User struct {
	FirstName string
	LastName  string
	PhoneNum  string
	Tktstatus map[string]Ticket
}

type Ticket struct {
	IsBooked     bool
	Price        int16
	ID           string
	Timing       string
	Expired      bool
	UserPhoneNum string
}
type UpdateTktTimeRequest struct {
	TicketID string `json:"tId"`
	Time     string `json:"time"`
}
type BookTicketRequest struct {
	User   User
	Timing string
	SeatNo string
}
