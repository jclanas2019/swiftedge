package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swiftedge",
	Short: "SwiftEdge - Herramienta CLI modular y extensible",
	Long:  `SwiftEdge es una herramienta de línea de comandos que también funciona como un SDK modular para múltiples propósitos.`,
}

// Execute inicia el comando raíz.
func Execute() error {
	return rootCmd.Execute()
}
