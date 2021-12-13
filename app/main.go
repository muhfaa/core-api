package main

import (
	"context"

	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"

	"core-data/modules/mysql"

	orderController "core-data/api/v1/order"
	orderService "core-data/business/order"
	kerusakanModule "core-data/modules/kerusakan"
	orderModule "core-data/modules/order"
	teknisiModule "core-data/modules/teknisi"

	config "core-data/config"

	coreDataApp "core-data/api/v1"

	"github.com/labstack/gommon/log"
)

func newDatabaseConnection() *sqlx.DB {
	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.GetConfig().MySQL.User,
		config.GetConfig().MySQL.Password,
		config.GetConfig().MySQL.Host,
		config.GetConfig().MySQL.Port,
		config.GetConfig().MySQL.Name)

	db := mysql.NewDatabaseConnection(uri)

	return db
}

func newOrderService(dbsql *sqlx.DB) orderService.Service {
	orderRepoSQL := orderModule.NewMySQLRepository(dbsql)
	teknisiRepoHTTP := teknisiModule.NewHTTPRequestRepository()
	kerusakanRepoHTTP := kerusakanModule.NewHTTPRequestRepository()

	return orderService.NewService(orderRepoSQL, teknisiRepoHTTP, kerusakanRepoHTTP)
}

func main() {

	// database init
	dbsql := newDatabaseConnection()

	// service init
	orderService := newOrderService(dbsql)

	// controller init
	orderControllerV1 := orderController.NewController(orderService)

	e := echo.New()

	//register API path and handler
	coreDataApp.RegisterPath(
		e,
		*orderControllerV1,
	)
	// run server
	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.GetConfig().BackendPort)

		if err := e.Start(address); err != nil {
			log.Error("failed to start server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Wait for interrupt signal to gracefully shutdown the server with
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("failed to grafefully shutting down server", err)
	}

}
