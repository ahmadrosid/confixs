package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/ahmadrosid/confixs/utils"
	"github.com/labstack/echo/v5"
)

func ListConfigHandler(c echo.Context) error {
	files, err := utils.GetFileList("./")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": files})
}

func CheckNginxHandler(c echo.Context) error {
	cmd := exec.Command("nginx", "-v")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"installed": false,
			"error":     err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"installed": true,
		"version":   string(output),
	})
}

func InstallNginxHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().WriteHeader(http.StatusOK)

	cmd := exec.Command("sudo", "apt-get", "update")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Error creating StdoutPipe for Cmd: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error starting Cmd: %v", err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Fprintf(c.Response(), "data: %s\n\n", scanner.Text())
		c.Response().Flush()
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Error waiting for Cmd: %v", err)
	}

	cmd = exec.Command("sudo", "apt-get", "install", "-y", "nginx")
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Error creating StdoutPipe for Cmd: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error starting Cmd: %v", err)
	}

	scanner = bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Fprintf(c.Response(), "data: %s\n\n", scanner.Text())
		c.Response().Flush()
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Error waiting for Cmd: %v", err)
	}

	fmt.Fprintf(c.Response(), "data: Installation completed\n\n")
	return nil
}
