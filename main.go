package main

import (
	"fmt"
	"net/http"
	"ticket-booking/dao"
	"ticket-booking/routes"
	"ticket-booking/service"
	"time"
)

func main() {
	routes.InitiateRoutes()
	dao.InitiateDatabase()
	go checkExpiryTickets()
	http.ListenAndServe(":8080", routes.AllRoutes)

}

/*
checkExpiryTickets checks if any ticket in database is
expired and deletes that ticket from db.
*/
func checkExpiryTickets() {
	ts := service.TicketService{}
	for {
		time.Sleep(20 * time.Second)
		fmt.Println("Checking And Deleting Expired Tickets.")
		ts.CheckExpiryAndDeleteTickets()
	}
}
