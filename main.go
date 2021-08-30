package main

import (
	"github.com/youhy/youhypm/models"
	"github.com/youhy/youhypm/server"
)

func main() {
	models.StartDB()

	server := server.NewServer()

	server.Run()
}
