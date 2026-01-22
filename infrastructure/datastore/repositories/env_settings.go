package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"taskflow/domain/entities"
	"taskflow/infrastructure/datastore"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type repositorySettings struct {
	connection *sql.DB
}

func NewRepositorySettings(config entities.Config) (datastore.RepositorySettings, error) {
	db, err := setupConnection(config)
	if err != nil {
		return nil, err
	}
	return repositorySettings{
		connection: db,
	}, nil
}

func (s repositorySettings) Connection() *sql.DB {
	return s.connection
}

func (s repositorySettings) Dismount() error {
	err := s.connection.Close()
	if err != nil {
		log.Printf("error in [Close]: %v", err)
		return fmt.Errorf("error in [Close]: %v", err)
	}

	return nil
}

func (s repositorySettings) ServerTime(
	ctx context.Context,
) (*time.Time, error) {
	//language=sql
	query := "SELECT CURRENT_TIMESTAMP"

	var serverTime time.Time
	err := s.connection.QueryRowContext(ctx, query).Scan(&serverTime)
	if err != nil {
		log.Printf("error in [Scan]: %v", err)
		return nil, fmt.Errorf("error in [Scan]: %v", err)
	}

	return &serverTime, nil
}

func setupConnection(config entities.Config) (*sql.DB, error) {
	//err := migrateDatabase(config)
	//if err != nil {
	//	return nil, err
	//}

	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, errors.Join(errors.New("failed to open database connection"), err)
	}

	db.SetMaxOpenConns(800)
	db.SetConnMaxLifetime(20 * time.Minute)
	db.SetConnMaxIdleTime(20 * time.Minute)

	return db, nil
}

//func migrateDatabase(config entities.Config) error {
//	connection := fmt.Sprintf(
//		"mysql://%s:%s@tcp(%s:%s)/%s?parseTime=true",
//		config.Database.User,
//		config.Database.Password,
//		config.Database.Host,
//		config.Database.Port,
//		config.Database.Name,
//	)
//
//	fsMigrations, err := iofs.New(fs, "migrations")
//	if err != nil {
//		return errors.Join(errors.New("failed to create migration using fs"), err)
//	}
//
//	// Defining the settings used by the migration service
//	migration, err := migrate.NewWithSourceInstance(
//		"iofs",
//		fsMigrations,
//		connection,
//	)
//	if err != nil {
//		return errors.Join(errors.New("failed to create migration with iofs and connection"), err)
//	}
//	defer migration.Close()
//
//	currentVersion, isDirty, err := migration.Version()
//	if err != nil && err.Error() != "no migration" {
//		return errors.Join(errors.New("failed to load the database version"), err)
//	}
//
//	if isDirty {
//		return errors.New("the database is in a dirty state, clear the errors and restart the system")
//	}
//
//	// Checks the need to apply a database migration
//	if DatabaseVersion > currentVersion {
//		slog.Info(
//			"apply database migrations",
//			slog.Int64("from", int64(currentVersion)),
//			slog.Int64("to", int64(DatabaseVersion)),
//		)
//		err = migration.Up()
//		if err != nil {
//			slog.Error(
//				"error running migration",
//				slog.String("error", err.Error()),
//				slog.Int64("from", int64(currentVersion)),
//				slog.Int64("to", int64(DatabaseVersion)),
//			)
//			return errors.Join(errors.New("failed to run up migrations"), err)
//		}
//	} else if DatabaseVersion < currentVersion {
//		slog.Info(
//			"downgrade database",
//			slog.Int64("from", int64(currentVersion)),
//			slog.Int64("to", int64(DatabaseVersion)),
//		)
//		err = migration.Down()
//		if err != nil {
//			slog.Error(
//				"error running downgrade",
//				slog.String("error", err.Error()),
//				slog.Int64("from", int64(currentVersion)),
//				slog.Int64("to", int64(DatabaseVersion)),
//			)
//			return errors.Join(errors.New("failed to run down migrations"), err)
//		}
//	} else {
//		slog.Info("database is up to date", slog.Int64("version", int64(DatabaseVersion)))
//	}
//
//	return nil
//}
