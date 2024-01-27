package configuration

import (
	atlas "ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	_ "github.com/lib/pq"
	dbconfig "go-patient-history/config/db"
	"go-patient-history/ent"
	"go-patient-history/ent/migrate"
	envconstant "go-patient-history/libs/common/constant/environment"
	logconstant "go-patient-history/libs/common/constant/logger"
	logger "go-patient-history/libs/common/logger/main"
)

func DatabaseConnection() *ent.Client {
	cfg := dbconfig.DBConfigLoad()

	drv, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SslMode))
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DatabaseConfiguration, Message: fmt.Sprintf("Failed Opening Connection: %#v", err.Error())})
		return nil
	}

	newDir, err := atlas.NewLocalDir(envconstant.MigrationDir)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DatabaseConfiguration, Message: fmt.Sprintf("Failed creating atlas migration directory: %v", err.Error())})
		return nil
	}

	opts := []schema.MigrateOption{
		schema.WithDir(newDir),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.Postgres),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	m, err := schema.NewMigrate(drv, opts...)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DatabaseConfiguration, Message: fmt.Sprintf("Failed creating migrate: %#v", err.Error())})
		return nil
	}

	if err = m.VerifyTableRange(context.Background(), migrate.Tables); err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DatabaseConfiguration, Message: fmt.Sprintf("Failed verifyint range allocations: %#v", err.Error())})
		return nil
	}
	client := ent.NewClient(ent.Driver(drv))

	logger.LogDebug(logger.LoggerPayload{FuncName: logconstant.DatabaseConfiguration, Message: "Database connection successfully"})

	return client
}
