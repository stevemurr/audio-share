package main

import (
	"flag"
	"murrman/audio-share/handler"
	"murrman/audio-share/storage/inmem"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	listenAddr = flag.String("addr", ":1323", "listen addr")
)

func init() {
	flag.Parse()
}

func main() {
	e := echo.New()
	db := inmem.New()
	h := &handler.Handler{
		DB: db,
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Static("/", "public/index.html")
	e.Static("/:id", "public/client.html")
	e.POST("/:id", h.UpdateRegions)
	e.POST("/upload", h.Upload)
	e.GET("/api/files", h.List)
	e.GET("/api/files/:file", h.GetFile)
	e.GET("/api/files/:file/data", h.GetPayload)
	e.Logger.Fatal(e.Start(*listenAddr))
}
