package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mashbens/cps/api"
	"github.com/mashbens/cps/app/modules"
	"github.com/mashbens/cps/config"

	"github.com/mashbens/cps/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config := config.GetConfig()
	dbCon := util.NewConnectionDatabase(config)

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
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
