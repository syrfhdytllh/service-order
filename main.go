package main

import (
	"service-order/database"
	"service-order/routers"
)

func main() {
	database.StartDB()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
