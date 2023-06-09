package main

import (
	"crud/modules/account"
	"crud/modules/customer"
	"crud/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := gin.New()

	// open connection db
	dbCrud := db.GormMysql()

	////check connection
	//checkdb, err := dbCrud.DB()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////ping to database
	//errconn := checkdb.Ping()
	//if err != nil {
	//	log.Fatal(errconn)
	//}

	fmt.Println("database connected..!")

	adminHandler := account.NewRouter(dbCrud)
	adminHandler.Handle(router)

	customerHandler := customer.NewRouter(dbCrud)
	customerHandler.Handle(router)

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
