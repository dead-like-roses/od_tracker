package main

import (
	"log"

	"github.com/dead-like-roses/od_tracker/storage"
	//	. "github.com/go-jet/jet/v2/postgres"
	"github.com/labstack/echo/v4"
	// _ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	db := storage.InitDatabase()
	err := storage.RunMigrations(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	//	ah := ActivityHandler{}

	//	e.POST("/register", ah.registerActivity)
	//	e.GET("/all", ah.listDatapoints)
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
