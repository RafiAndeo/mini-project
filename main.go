package main

import (
	"mini-project/middlewares"
	"mini-project/routes"
)

func main() {
	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
