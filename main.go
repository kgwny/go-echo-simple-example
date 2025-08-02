// Command main runs the Echo web server.
package main

import (
	"go-echo-simple-example/models"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}
	renderer := &TemplateRenderer{templates: tmpl}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.GET("/list", list)
	e.POST("/create", create)
	e.GET("/edit", edit)
	e.POST("/update", update)
	e.GET("/destroy", destroy)

	e.Logger.Fatal(e.Start(":8080"))
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	} else {
		log.Printf("expected map[string]interface{} for template data, got %T", data)
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

// delete
func destroy(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Fatalln(err)
	}
	models.DeleteTodo(id)
	return c.Redirect(http.StatusSeeOther, "/list")
}

// update
func update(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	content := c.FormValue("content")
	if content == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Content is required")
	}
	todo, err := models.GetTodo(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve todo")
	}
	todo.Content = content
	models.UpdateTodo(todo)
	return c.Redirect(http.StatusSeeOther, "/list")
}

// edit
func edit(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	todo, err := models.GetTodo(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve todo")
	}
	data := map[string]interface{}{
		"todo": todo,
	}
	return c.Render(http.StatusOK, "edit.html", data)
}

// create
func create(c echo.Context) error {
	content := c.FormValue("content")
	if content == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Content is required")
	}
	models.CreateTodo(content)
	return c.Redirect(http.StatusSeeOther, "/list")
}

// list
func list(c echo.Context) error {
	var todos []models.Todo
	models.DB.Find(&todos)
	data := map[string]interface{}{
		"todos": todos,
	}
	return c.Render(http.StatusOK, "list.html", data)
}
