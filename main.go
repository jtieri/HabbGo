package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"time"

	"github.com/jtieri/habbgo/config"
	"github.com/jtieri/habbgo/transport"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go listenForTerminationSignals(cancel)

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, opts))

	cfg, err := readConfigFile(log, config.DefaultConfigPath)
	if err != nil {
		panic(err)
	}

	if cfg.Global.Debug {
		opts.Level = slog.LevelDebug
		log = slog.New(slog.NewTextHandler(os.Stdout, opts))
	}

	db, err := connectToDatabase(log, cfg)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error("Failed to close database", "error", err)
		}
	}(db)

	// Start the TCP game server.
	gameServer := transport.NewGameServer(log, cfg, db)
	errChan := gameServer.Start(ctx)

	// Listen for errors from the game server.
	if err := <-errChan; err != nil && !errors.Is(err, context.Canceled) {
		log.Error("Failed to start game server", "error", err)
	}
}

// listenForTerminationSignals listens for OS termination signals and attempts to cancel the root context.
// If the application does not gracefully shut down, it will force a hard shutdown after one minute.
func listenForTerminationSignals(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt) // Using signal.Notify, instead of signal.NotifyContext, in order to see details of signal.

	sig := <-sigCh // Wait for interrupt signal.

	cancel() // Cancel the root context.

	// Short delay before printing the received signal message.
	// This should result in cleaner output from non-interactive commands that stop quickly.
	time.Sleep(250 * time.Millisecond)
	fmt.Fprintf(os.Stderr, "Received signal %v. Attempting clean shutdown. Send interrupt again to force hard shutdown.\n", sig)

	debug.SetTraceback("all") // Dump all goroutines on panic, not just the current one.

	// Block waiting for a second interrupt or a timeout.
	// The main goroutine ought to finish before either case is reached.
	// But if a case is reached, panic so that we get a non-zero exit and a dump of remaining goroutines.
	select {
	case <-time.After(time.Minute):
		panic(errors.New("habbgo did not shut down within one minute of interrupt"))
	case sig := <-sigCh:
		panic(fmt.Errorf("received signal %v; forcing quit", sig))
	}
}

// readConfigFile attempts to read the specified configuration file from disk.
func readConfigFile(log *slog.Logger, cfgFile string) (*config.Config, error) {
	log.Info(
		"Attempting to read config file from disk",
		"path", cfgFile,
	)

	cfgBz, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	cfg := &config.Config{}

	err = yaml.Unmarshal(cfgBz, cfg)
	if err != nil {
		return nil, err
	}

	log.Info(
		"Read config file from disk",
		"path", cfgFile,
	)

	return cfg, nil
}

// connectToDatabase attempts to open a connection to the database, it also checks that the connection is alive.
func connectToDatabase(log *slog.Logger, c *config.Config) (*sql.DB, error) {
	log.Info(
		"Attempting to connect to the database",
		"host", c.DB.Host,
		"port", c.DB.Port,
		"username", c.DB.Username,
		"password", c.DB.Password,
		"db_name", c.DB.Name,
		"db_driver", c.DB.Driver,
		"sslmode", c.DB.SSLMode,
	)

	// Open the connection to the database.
	db, err := sql.Open(c.DB.Driver, c.DB.ConnectionString())
	if err != nil {
		return nil, err
	}

	// Check that the connection to the database is alive.
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Info(
		"Connected to the database",
		"host", c.DB.Host,
		"port", c.DB.Port,
		"username", c.DB.Username,
		"password", c.DB.Password,
		"db_name", c.DB.Name,
		"db_driver", c.DB.Driver,
		"sslmode", c.DB.SSLMode,
	)

	return db, nil
}
