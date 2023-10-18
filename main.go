package main

import (
	"log"
)

func main()  {
	store, err := newPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.init(); err != nil {
		log.Fatal(err)
	}

	server := newApiServer(":3000", store)
	server.run()
}