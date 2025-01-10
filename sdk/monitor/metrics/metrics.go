package metrics

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var (
	CPUUsage    string
	MemoryUsage string
	DiskUsage   string
	Connections int
)

// CollectMetrics recolecta métricas del sistema y de la base de datos.
func CollectMetrics(db *sql.DB) {
	for {
		// Métricas de CPU
		cpuPercent, _ := cpu.Percent(0, false)
		CPUUsage = fmt.Sprintf("%.2f%%", cpuPercent[0])

		// Métricas de memoria
		vm, _ := mem.VirtualMemory()
		MemoryUsage = fmt.Sprintf("%.2f%% usado de %.2f GB", vm.UsedPercent, float64(vm.Total)/1024/1024/1024)

		// Métricas de disco
		diskInfo, _ := disk.Usage("/")
		DiskUsage = fmt.Sprintf("%.2f%% usado, %.2f GB libres", diskInfo.UsedPercent, float64(diskInfo.Free)/1024/1024/1024)

		// Conexiones activas a la base de datos
		updateDatabaseConnections(db)

		time.Sleep(3 * time.Second)
	}
}

func updateDatabaseConnections(db *sql.DB) {
	if err := db.QueryRow("SHOW STATUS LIKE 'Threads_connected'").Scan(new(string), &Connections); err != nil {
		log.Printf("Error obteniendo conexiones: %v", err)
	}
}
