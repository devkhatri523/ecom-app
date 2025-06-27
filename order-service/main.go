package main

import (
	"net/http"
	"order-service/controller"
	"order-service/repo"
	"order-service/router"
	"order-service/service"

	"github.com/devkhatri523/ecom-app/go-config/v4/config"
	"github.com/devkhatri523/ecom-app/go-config/v4/database"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db, err := OpenDb()
	if err != nil {
		panic(err)
	}
	orderRepository := repo.NewOrderRepositoryImpl(db.OrmInstance)
	orderLineRepository := repo.NewOrderLineRepositoryImpl(db.OrmInstance)
	orderService := service.NewOrderServiceImpl(orderRepository, orderLineRepository)
	orderController := controller.NewOrderController(orderService)
	orderRouters := router.OrderRouter(orderController)
	server := &http.Server{
		Addr:    "0.0.0.0:9090",
		Handler: orderRouters,
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
