package main

import (
	"fmt"

	"github.com/gavruk/crater"
	"github.com/gavruk/crater/middleware"

	"github.com/gavruk/crater-example/models"
)

func main() {
	fmt.Println("Listening on port 8080:")
	fmt.Println("http://localhost:8080")

	settings := &crater.Settings{}
	settings.ViewsPath = "./Views"
	settings.StaticPath = "./Content"

	app := crater.NewApp()
	app.Settings(settings)

	app.Static("/content")

	app.Use(middleware.InMemorySession)

	app.Get("/signin", func(req *crater.Request, res *crater.Response) {
		res.RenderTemplate("signin", nil)
	})

	app.Post("/signin", func(req *crater.Request, res *crater.Response) {
		signInModel := new(models.SignInModel)
		if err := req.Parse(signInModel); err != nil {
			res.Json(&models.JsonResponse{false, err.Error()})
			return
		}
		if signInModel.Password != "123" {
			res.Json(&models.JsonResponse{false, "Credentials are not valid!"})
			return
		}
		req.Session.Value = signInModel.Username
		res.Json(&models.JsonResponse{true, ""})
	})

	app.Get("/", func(req *crater.Request, res *crater.Response) {
		userFromSession := req.Session.Value
		if userFromSession == nil {
			res.Redirect("/signin")
			return
		}
		username := userFromSession.(string)
		model := new(models.ViewModel)
		model.User = models.User{Name: username}
		res.RenderTemplate("index", model)
	})

	app.Get("/signout", func(req *crater.Request, res *crater.Response) {
		if req.Session != nil {
			req.Session.Abandon()
		}
		res.Redirect("/signin")
	})

	app.Get("/about", func(req *crater.Request, res *crater.Response) {
		userFromSession := req.Session.Value
		if userFromSession == nil {
			res.Redirect("/signin")
			return
		}
		username := userFromSession.(string)
		model := new(models.ViewModel)
		model.User = models.User{Name: username}
		res.RenderTemplate("about", model)
	})

	app.Get("/string", func(req *crater.Request, res *crater.Response) {
		res.Send("<h1>Hello World</h1>")
	})

	// example: localhost:8080/hello/John
	app.Get("/hello/{name}", func(req *crater.Request, res *crater.Response) {
		name := req.RouteParams["name"]
		res.Send(fmt.Sprintf("<h1>Hello, %s</h1>", name))
	})

	app.Listen(":8080")
}
