package fiber

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

var useEmbed = true

func Start(viewsfs embed.FS, embedDirStatic embed.FS, publicFs embed.FS) {
	fmt.Println("hello")

	var app *fiber.App

	if useEmbed {
		engine := html.NewFileSystem(http.FS(viewsfs), ".html")

		// Pass the engine to the Views
		app = fiber.New(fiber.Config{
			Views: engine,
		})
		app.Use("/assets", filesystem.New(filesystem.Config{
			Root:       http.FS(embedDirStatic),
			PathPrefix: "client/dist/assets",
			Browse:     true,
		}))
		app.Use("/vite.svg", filesystem.New(filesystem.Config{
			Root:       http.FS(publicFs),
			PathPrefix: "client/dist/vite.svg",
			Browse:     true,
		}))
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Render("client/dist/index", fiber.Map{})
		})

	} else {
		engine := html.New("./client/dist", ".html")
		app = fiber.New(fiber.Config{
			Views: engine,
		})
		app.Static("/assets", "./client/dist/assets")
		app.Static("/vite.svg", "./client/dist/vite.svg")
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Render("index", fiber.Map{})
		})
	}

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3007")
	if err != nil {
		log.Fatal(err)
	}
}
