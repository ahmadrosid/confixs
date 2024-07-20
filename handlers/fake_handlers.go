package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

func FakeInstallNginxHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().WriteHeader(http.StatusOK)

	fakeOutputs := []string{
		"Updating package lists...",
		"Reading package lists... Done",
		"Building dependency tree... Done",
		"Reading state information... Done",
		"Installing nginx...",
		"Unpacking nginx (1.18.0-0ubuntu1.2) ...",
		"Setting up nginx (1.18.0-0ubuntu1.2) ...",
		"Processing triggers for systemd (245.4-4ubuntu3.13) ...",
		"Processing triggers for man-db (2.9.1-1) ...",
		"Processing triggers for ufw (0.36-6ubuntu1) ...",
		"Installation completed",
	}

	for _, output := range fakeOutputs {
		fmt.Fprintf(c.Response(), "data: %s\n\n", output)
		c.Response().Flush()
		time.Sleep(300 * time.Millisecond) // Simulate delay between outputs
	}

	// c.Response().Flush()
	fmt.Fprintf(c.Response(), fmt.Sprintf("event: close\ndata: %s\n\n", "{ \"status\": \"success\" }"))

	return nil
}
