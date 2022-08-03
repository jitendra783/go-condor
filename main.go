package main

import (
	"database/sql"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fetal("cannot load config:", err)
	}

	conn, err := sql.Open(config.dbDriver, config.dbSource)
	if err != nil {
		log.Fetal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.serverAddress)
	if err != nil {
		log.Fetal("cannot start server:", err)
	}
}
