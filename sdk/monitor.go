package cmd

import (
	"log"
	"swiftedge/sdk/monitor"

	"github.com/spf13/cobra"
)

var (
	dbUser     string
	dbPassword string
	dbHost     string
	dbName     string
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Ejecuta el monitor de base de datos",
	Long:  `Monitor de base de datos MySQL/MariaDB con métricas de sistema, conexiones y replicación.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := monitor.Run(dbUser, dbPassword, dbHost, dbName)
		if err != nil {
			log.Fatalf("Error ejecutando el monitor: %v", err)
		}
	},
}

func init() {
	monitorCmd.Flags().StringVarP(&dbUser, "user", "u", "", "Usuario de la base de datos (requerido)")
	monitorCmd.Flags().StringVarP(&dbPassword, "password", "p", "", "Contraseña de la base de datos (requerido)")
	monitorCmd.Flags().StringVarP(&dbHost, "host", "H", "localhost", "Host de la base de datos (por defecto: localhost)")
	monitorCmd.Flags().StringVarP(&dbName, "database", "d", "", "Nombre de la base de datos (requerido)")

	rootCmd.AddCommand(monitorCmd)
}
