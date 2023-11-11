package server

import (
	"github.com/mariojuzar/fleet-inventory/config"
	"github.com/mariojuzar/fleet-inventory/internal/handler/api"
	"github.com/mariojuzar/fleet-inventory/internal/infrastructures/mysql"
	"github.com/mariojuzar/fleet-inventory/internal/interfaces/dao"
	"github.com/mariojuzar/fleet-inventory/internal/usecases/spacecraft"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:   "serve-http",
		Short: "Fleet Inventory Management System HTTP",
		Long:  "Serve Fleet Inventory Management System through HTTP",
		RunE:  run,
	}
)

func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}

func run(cmd *cobra.Command, args []string) error {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	DB, err := mysql.New(&mysql.Options{Cfg: cfg.Database})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize db connection")
	}

	spaceCraftRepo := dao.NewSpaceCraft(&dao.SpaceCraftRepositoryOpts{DB: DB})
	spaceCraftUc := spacecraft.New(&spacecraft.Opts{
		SpaceCraftRepo: spaceCraftRepo,
	})

	server := api.New(&api.Options{
		Cfg:  cfg.Server,
		ScUc: spaceCraftUc,
	})

	server.Run()

	return nil
}
