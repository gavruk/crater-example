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
		var username string
		if value, ok := req.GetString("name"); ok {
			username = value
		}

		user := &models.User{Name: username}
		res.Render("hello", user)
	})

	server := crater.Server{}
	server.Listen(":8080")
}
