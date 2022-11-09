package main

import (
	"storageproject/cmd/server/routes"
	"storageproject/pkg/db"
)

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")
}
