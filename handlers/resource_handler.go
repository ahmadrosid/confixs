package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetChartDataHandler(c echo.Context) error {
	usage := ResourceUsageResponse{}

	// Get CPU usage
	cpuPercent, err := cpu.Percent(0, false)
	if err == nil && len(cpuPercent) > 0 {
		usage.CPU = cpuPercent[0]
	}

	// Get memory usage
	memInfo, err := mem.VirtualMemory()
	if err == nil {
		usage.Memory = memInfo.UsedPercent
	}

	// Get disk usage
	diskInfo, err := disk.Usage("/")
	if err == nil {
		usage.Disk = diskInfo.UsedPercent
	}

	return c.JSON(http.StatusOK, usage)
}
