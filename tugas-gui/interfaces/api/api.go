package api

import (
	"context"
	"tugas-gui/internal/config"
	"tugas-gui/internal/sensor"
	"fmt"
)

type Api struct {
	cfg    *config.Config
	// log    *utils.Logger
	// Client *Client
	// Server *Server
}

func NewApi(cfg *config.Config) *Api {
	return &Api{
		cfg: cfg,
	}
}

func (a *Api) GetSensorData() string {
	cpu, _ := sensor.GetCPUUsage()
	memUsed, memTotal, _ := sensor.GetMemoryUsage()
	diskUsed, diskTotal, err := sensor.GetDiskUsage(a.cfg.Sensors.DiskPath)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(a.cfg.Sensors.DiskPath)

	return fmt.Sprintf("Cpu: %.2f%%\nMemory: %.2f%% dari %d bytes\nDisk: %.2f%% dari %d bytes\n", cpu, memUsed, memTotal, diskUsed, diskTotal)
}

func (a *Api) Startup(ctx context.Context) {
	fmt.Println("API started")
}
