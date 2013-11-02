package main

import (
	"fmt"
	"time"

	"github.com/gavruk/crater"
	"github.com/gavruk/crater/session"

	"github.com/gavruk/crater-example/models"
)

func main() {
	fmt.Println("Listening on port 8080:")
	fmt.Println("http://localhost:8080")

	settings := &crater.Settings{}
	settings.ViewsPath = "./Views"
	settings.StaticFilesPath = "./Content"

	app := crater.NewApp(settings)

	app.UseSessionStore(session.NewInMemorySessionStore(), time.Hour)

	app.HandleStaticContent("/content")

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
		req.Session.Value = signInModel
		res.Json(&models.JsonResponse{true, ""})
	})

	app.Get("/", func(req *crater.Request, res *crater.Response) {
		userFromSession := req.Session.Value
		if userFromSession == nil {
			res.Redirect("/signin")
			return
		}
		res.RenderTemplate("index", userFromSession)
	})

	app.Get("/signout", func(req *crater.Request, res *crater.Response) {
		req.Session.Abandon()
		res.Redirect("/signin")
	})

	app.Get("/about", func(req *crater.Request, res *crater.Response) {
		res.RenderTemplate("about", nil)
	})

	app.Get("/string", func(req *crater.Request, res *crater.Response) {
		res.RenderString("<h1>Hello World</h1>")
	})

	app.Listen(":8080")
}
