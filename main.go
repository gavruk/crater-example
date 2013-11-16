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
	settings.ViewsPath = "./views"
	settings.StaticPath = "./content"

	app := crater.NewApp()
	app.Settings(settings)

	app.Static("/content")

	app.Use(middleware.InMemorySession)

	// Check if authorized. If not - redirect to signin page
	app.Use(func(req *crater.Request, res *crater.Response) {
		userFromSession := req.Session.Value
		if userFromSession == nil && req.URL.String() != "/signin" {
			res.Redirect("/signin")
		}
	})

	// =============
	// Auth
	// =============

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

	app.Get("/signout", func(req *crater.Request, res *crater.Response) {
		if req.Session != nil {
			req.Session.Abandon()
		}
		res.Redirect("/signin")
	})

	// =============
	// CRUD
	// =============

	books := make(map[string]*models.BookModel)

	app.Get("/", func(req *crater.Request, res *crater.Response) {
		model := models.NewBookListViewModel(req.Session.Value)
		for _, v := range books {
			model.Books = append(model.Books, v)
		}
		res.RenderTemplate("index", model)
	})

	app.Get("/add", func(req *crater.Request, res *crater.Response) {
		model := models.NewAddBookViewModel(req.Session.Value)
		res.RenderTemplate("add", model)
	})

	app.Post("/add", func(req *crater.Request, res *crater.Response) {
		book := new(models.BookModel)
		if err := req.Parse(book); err != nil {
			res.Redirect("/add")
		}
		book.Id = GenerateId()
		books[book.Id] = book
		res.Redirect("/")
	})

	app.Get("/edit/{id}", func(req *crater.Request, res *crater.Response) {
		model := models.NewAddBookViewModel(req.Session.Value)
		bookId := req.RouteParams["id"]
		book := books[bookId]
		if book == nil {
			res.Redirect("/")
			return
		}
		model.Book = book
		res.RenderTemplate("edit", model)
	})

	app.Post("/edit", func(req *crater.Request, res *crater.Response) {
		book := new(models.BookModel)
		if err := req.Parse(book); err != nil {
			res.Redirect("/edit")
		}
		books[book.Id] = book
		res.Redirect("/")
	})

	app.Get("/remove/{id}", func(req *crater.Request, res *crater.Response) {
		bookId := req.RouteParams["id"]
		delete(books, bookId)
		res.Redirect("/")
	})

	// ===========
	// About
	// ============
	app.Get("/about", func(req *crater.Request, res *crater.Response) {
		model := models.NewIndexViewModel(req.Session.Value)
		res.RenderTemplate("about", model)
	})

	app.Listen(":8080")
}
