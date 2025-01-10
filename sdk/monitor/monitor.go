package monitor

import (
	"swiftedge/sdk/monitor/helpers"
	"swiftedge/sdk/monitor/metrics"
	"swiftedge/sdk/monitor/ui"
)

// Run inicia el monitor.
func Run(user, password, host, database string) error {
	db, err := helpers.ConnectDB(user, password, host, database)
	if err != nil {
		return err
	}
	defer db.Close()

	// Inicia la recolección de métricas
	go metrics.CollectMetrics(db)

	// Dibuja la interfaz gráfica
	ui.DrawUI()

	return nil
}
