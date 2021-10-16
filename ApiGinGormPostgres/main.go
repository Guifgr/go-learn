package main

import (
	"github.com/guifgr/GO-LEARN/ApiRestDatabase/database"
	"github.com/guifgr/GO-LEARN/ApiRestDatabase/server"
)

func main() {
	database.StartDB()

	newServer := server.NewServer()
	newServer.Run()
}
