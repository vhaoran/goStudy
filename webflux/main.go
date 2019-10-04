package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type (
	Geolocation struct {
		Altitude  float64
		Latitude  float64
		Longitude float64
	}
)

var (
	locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		for _, l := range locations {
			for j := 0; j < 30; j++ {
				l.Latitude += float64(j)
				if err := json.NewEncoder(c.Response()).Encode(l); err != nil {
					return err
				}
				c.Response().Flush()
				time.Sleep(100 * time.Millisecond)
			}
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
