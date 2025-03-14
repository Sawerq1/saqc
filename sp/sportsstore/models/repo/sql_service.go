package repo

import (
	"context"
	"database/sql"
	"platform/config"
	"platform/logging"
	"platform/services"
	"sportsstore/models"
	"sync"
)

func RegisterSQLRepositoryService() {
	var db *sql.DB
	var commands *SQLCommands
	var needInit bool
	loadOnce := sync.Once{}
	resetOnce := sync.Once{}
	services.AddScoped(func(ctx context.Context, config config.Configuration,
		logger logging.Logger) models.Repository {
		loadOnce.Do(func() {
			db, commands, needInit = openDB(config, logger)
		})
		repo := &SQLRepository{
			Configuration: config,
			Logger:        logger,
			Commands:      *commands,
			DB:            db,
			Context:       ctx,
		}
		resetOnce.Do(func() {
			if needInit || config.GetBoolDefault("sql:always_reset", true) {
				repo.Init()
				repo.Seed()
			}
		})
		return repo
	})
}
