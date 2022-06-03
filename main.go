package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gym-app/api"
	"gym-app/app/modules"
	"gym-app/config"
	"gym-app/util"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

var DB *gorm.DB

func main() {

	config := config.GetConfig()
	dbCon := util.NewConnectionDatabase(config)

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!!")
	})

	api.RegisterRoutes(e, &controllers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}
	e.Logger.Fatal(e.Start(":" + port))

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	defer dbCon.CloseConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
