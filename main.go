package main

import (
	"MyProjects/Zomentum/dao"
	"MyProjects/Zomentum/routes"
	"MyProjects/Zomentum/service"
	"fmt"
	"net/http"
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
