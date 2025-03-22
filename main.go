package main

// import "github.com/gofiber/fiber/v2"

import (
	"embed"
	"server/fiber"
)

//go:embed client/dist/*
var viewsfs embed.FS

//go:embed client/dist/*
var publicFs embed.FS

//go:embed client/dist/assets/*
var embedDirStatic embed.FS

func main() {
	fiber.Start(viewsfs, embedDirStatic, publicFs)
}
