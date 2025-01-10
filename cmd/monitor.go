package cmd

import (
	"log"
	"swiftedge/sdk/monitor"

	"github.com/spf13/cobra"
)

var (
	user     string
	password string
	host     string
	database string
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Ejecuta el monitor de base de datos",
	Long:  `Monitor de base de datos MySQL/MariaDB con métricas del sistema, conexiones activas y replicación.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := monitor.Run(user, password, host, database)
		if err != nil {
			log.Fatalf("Error ejecutando el monitor: %v", err)
		}
	},
}

func init() {
	monitorCmd.Flags().StringVarP(&user, "user", "u", "", "Usuario de la base de datos (requerido)")
	monitorCmd.Flags().StringVarP(&password, "password", "p", "", "Contraseña de la base de datos (requerido)")
	monitorCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host de la base de datos (por defecto: localhost)")
	monitorCmd.Flags().StringVarP(&database, "database", "d", "", "Nombre de la base de datos (requerido)")

	rootCmd.AddCommand(monitorCmd)
}
