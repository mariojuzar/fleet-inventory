package cmd

import (
	"fmt"
	"github.com/mariojuzar/fleet-inventory/cmd/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Fleet Inventory Service",
		Short: "Fleet Inventory Management System - Backend Service",
	}
)

func Execute() {
	rootCmd.AddCommand(server.ServeHTTPCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Msg(fmt.Sprint("Error: ", err.Error()))
		os.Exit(-1)
	}
}
