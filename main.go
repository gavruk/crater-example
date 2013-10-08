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
	config.ViewsPath = "./Views"
	config.StaticFilesPath = "./Content"

	app.Settings(config)

	app.HandleStaticFiles("/content")

	app.Get("/", func(req *crater.Request, res *crater.Response) {
		res.Render("index", nil)
	})

	app.Get("/hello", func(req *crater.Request, res *crater.Response) {

		user := &models.User{}

		if err := req.Parse(user); err != nil {
			fmt.Println(err.Error())
		}

		res.Render("hello", user)
	})

	server := crater.Server{}
	server.Listen(":8080")
}
