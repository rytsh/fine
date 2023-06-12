package args

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/worldline-go/logz"

	"github.com/rytsh/fine/internal/config"
	"github.com/rytsh/fine/internal/server"
	"github.com/rytsh/liz/shutdown"
)

var ErrShutdown = errors.New("shutting down signal received")

var rootCmd = &cobra.Command{
	Use:           "fine",
	Short:         "file management service",
	Long:          config.Logo + "rest api for file management",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := logz.SetLogLevel(config.App.Log.Level); err != nil {
			log.Error().Err(err).Msg("failed to set log level")
		}

		log.Info().Msgf("fine [%s]", cmd.Version)
	},

	Example: "fine -c ./config.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		// load configuration
		overrideValues := make(map[string]config.OverrideHold)
		override(overrideValues)

		visit := func() {
			// override used cmd values
			cmd.Flags().Visit(func(f *pflag.Flag) {
				if v, ok := overrideValues[f.Name]; ok {
					*v.Memory = v.Value
				}
			})
		}

		if err := config.Load(ctx, visit); err != nil {
			return err
		}

		return fine(ctx)
	},
}

// Execute is the entry point for the application.
func Execute(ctx context.Context) error {
	rootCmd.Version = config.Info.Version
	rootCmd.Long += "\nversion: " + config.Info.Version + " commit: " + config.Info.BuildCommit + " buildDate:" + config.Info.BuildDate

	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.Flags().StringVar(&config.App.Log.Level, "log-level", config.App.Log.Level, "log level (debug, info, warn, error, fatal, panic)")
	rootCmd.Flags().StringVarP(&config.File, "config", "c", config.File, "config file path")
}

// override function hold first values of definitions.
// Use with pflag visit function.
func override(ow map[string]config.OverrideHold) {
	ow["log-level"] = config.OverrideHold{Memory: &config.App.Log.Level, Value: config.App.Log.Level}
}

func fine(ctx context.Context) (err error) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, ctxCancel := context.WithCancel(ctx)
	defer ctxCancel()

	wg.Add(1)
	go shutdown.Global.WatchCtx(ctx, wg)

	wg.Add(1)
	go func() {
		defer wg.Done()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-sig:
			log.Warn().Msg("received shutdown signal")
			ctxCancel()

			if err != nil {
				err = ErrShutdown
			}
		case <-ctx.Done():
		}
	}()

	// start service
	log.Info().Msg("starting service")

	server.Start()

	log.Warn().Msg("service stopped")

	return nil
}
