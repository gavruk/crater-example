package main

import (
	"fmt"
	"time"

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
		req.Session.Value = &models.User{Name: "Bob"}
		req.Cookie.Set("hello", "world", time.Now().Add(time.Hour))
		res.Render("index", nil)
	})

	app.Get("/hello", func(req *crater.Request, res *crater.Response) {
		fmt.Println(req.Session.Value)
		req.Session.Abandon()
		fmt.Println(req.Cookie.Get("hello"))

		user := new(models.User)

		if err := req.Parse(user); err != nil {
			fmt.Println(err.Error())
		}

		res.Render("hello", user)
	})

	app.Get("/post", func(req *crater.Request, res *crater.Response) {
		res.Render("post", nil)
	})

	app.Post("/post", func(req *crater.Request, res *crater.Response) {
		user := new(models.User)

		if err := req.Parse(user); err != nil {
			fmt.Println(err.Error())
		}

		res.Render("post", user)
	})

	app.Post("/postjson", func(req *crater.Request, res *crater.Response) {
		user := &models.User{}

		if err := req.Parse(user); err != nil {
			fmt.Println(err.Error())
		}

		res.Json(user)
	})

	app.Get("/redirect", func(req *crater.Request, res *crater.Response) {
		res.Redirect("/post")
	})

	server := crater.Server{}
	server.Listen(":8080")
}
