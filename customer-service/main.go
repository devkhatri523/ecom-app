package main

import (
	"customer-service/controller"
	"customer-service/repo"
	"customer-service/router"
	"customer-service/service"
	"fmt"
	"net/http"

	"github.com/devkhatri523/ecom-app/go-config/v4/config"
	"github.com/devkhatri523/ecom-app/go-config/v4/database"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db, err := OpenDb()
	if err != nil {
		fmt.Sprintf("Error while connecting  database: %s", err)
	}
	customerRepo := repo.NewCustomerRepositoryImpl(db.OrmInstance)
	customerService := service.NewCustomerServiceImpl(customerRepo)
	customerController := controller.NewCustomerController(customerService)
	routers := router.CustomerRouter(customerController)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: routers,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func OpenDb() (*database.OrmDB, error) {
	host := config.Default().GetString("db.postgres.host")
	port := config.Default().GetInt("db.postgres.port")
	userName := config.Default().GetString("db.postgres.username")
	password := config.Default().GetString("db.postgres.password")
	dbname := config.Default().GetString("db.postgres.database")

	orm, err := database.OpenOrm(host, port, userName, password, dbname)
	if err != nil {
		return nil, err
	}
	return orm, nil

}
