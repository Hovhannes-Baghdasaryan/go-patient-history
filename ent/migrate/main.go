//go:build ignore

package main

import (
	atlas "ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	config "github.com/Hovhannes-Baghdasaryan/go-patient-history/config/db"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent/migrate"
	envconstant "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/constant/environment"
	logconstant "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/constant/logger"
	logger "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/logger/main"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	if err := os.MkdirAll(envconstant.MigrationDir, 0755); err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.MigrationMain, Message: fmt.Sprintf("Creating migration directory: %v", err.Error())})
	}

	newDir, err := atlas.NewLocalDir(envconstant.MigrationDir)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.MigrationMain, Message: fmt.Sprintf("Failed creating atlas migration directory: %v", err.Error())})
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(newDir),                      // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	if len(os.Args) != 2 {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.MigrationMain, Message: "Migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'"})
	}

	cfg := config.DBConfigLoad()

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SslMode), os.Args[1], opts...)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.MigrationMain, Message: fmt.Sprintf("Failed generating migration file: %v", err)})
	}
}
