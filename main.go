package main

import (
	"mini-project/middlewares"
	routes "mini-project/route"
)

func main() {
	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
