package main

import (
	"fmt"
	"github.com/gavruk/crater"

	"github.com/gavruk/crater-example/models"
)

func main() {
	fmt.Println("Listening on port 8080:")
	fmt.Println("http://localhost:8080")

	app := crater.App{}

	config := crater.Settings{}
	config.ViewPath = "./Views"

	app.Settings(config)

	app.Get("/", func(req *crater.Request, res *crater.Response) {
		res.Render("index", nil)
	})

	app.Get("/hello", func(req *crater.Request, res *crater.Response) {
		user := &models.User{Name: "Your name"}
		res.Render("hello", user)
	})

	server := crater.Server{}
	server.Listen(":8080")
}
