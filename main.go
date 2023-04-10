package main

import (
	"fmt"
	"github.com/febryanmz/Data-Diri/utils/database"

	"github.com/febryanmz/Data-Diri/factory"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectToDB()
	defer db.Close()

	e := gin.New()
	e.Use(gin.Recovery())

	factory.InitFactory(e, db)
	e.Run(fmt.Sprintf(":%d", 8080))

}
