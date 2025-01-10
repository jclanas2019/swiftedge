package ui

import (
	"fmt"
	"log"
	"swiftedge/sdk/monitor/metrics"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func DrawUI() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Error inicializando la UI: %v", err)
	}
	defer ui.Close()

	// Define los widgets de la UI
	cpuBox := widgets.NewParagraph()
	cpuBox.Title = "CPU Usage"
	cpuBox.Text = metrics.CPUUsage
	cpuBox.SetRect(0, 0, 30, 3)

	memoryBox := widgets.NewParagraph()
	memoryBox.Title = "Memory Usage"
	memoryBox.Text = metrics.MemoryUsage
	memoryBox.SetRect(0, 3, 30, 6)

	diskBox := widgets.NewParagraph()
	diskBox.Title = "Disk Usage"
	diskBox.Text = metrics.DiskUsage
	diskBox.SetRect(0, 6, 30, 9)

	connectionsBox := widgets.NewParagraph()
	connectionsBox.Title = "Connections"
	connectionsBox.Text = fmt.Sprintf("%d", metrics.Connections)
	connectionsBox.SetRect(0, 9, 30, 12)

	// Renderiza la UI inicial
	ui.Render(cpuBox, memoryBox, diskBox, connectionsBox)

	// Actualiza la UI periódicamente
	ticker := time.NewTicker(3 * time.Second).C
	uiEvents := ui.PollEvents()

	for {
		select {
		case <-ticker:
			cpuBox.Text = metrics.CPUUsage
			memoryBox.Text = metrics.MemoryUsage
			diskBox.Text = metrics.DiskUsage
			connectionsBox.Text = fmt.Sprintf("%d", metrics.Connections)

			ui.Render(cpuBox, memoryBox, diskBox, connectionsBox)

		case e := <-uiEvents:
			if e.Type == ui.KeyboardEvent && (e.ID == "q" || e.ID == "<C-c>") {
				return
			}
		}
	}
}
