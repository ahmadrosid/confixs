package ui

import (
	"embed"

	"github.com/labstack/echo/v5"
)

//go:embed *
var distDir embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistDirFS = echo.MustSubFS(distDir, "dist")
