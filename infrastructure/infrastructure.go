package infrastructure

import (
	"log/slog"
	"taskflow/domain/usecases"
	"taskflow/infrastructure/config"
	"taskflow/infrastructure/datastore/repositories"
	"taskflow/infrastructure/modules"
	"taskflow/infrastructure/util"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

func Init(configFilePath string) {
	var cfg config.Config
	_, err := toml.DecodeFile(configFilePath, &cfg)
	if err != nil {
		panic(err)
	}

	// Config database
	err = config.Connect(&cfg)
	if err != nil {
		panic(err)
	}

	// Config utils
	maker := util.NewPasetoMaker(cfg.Crypt.SymmetricKey)
	crypt := util.NewCrypt(maker)

	conn := cfg.Database.Conn

	// Repositories
	authRepository := repositories.NewAuthRepository(conn)

	// Use Cases
	authUseCases := usecases.NewAuthUseCases(authRepository, crypt)

	// Modules
	healthModule := modules.NewHealthModule()
	authModule := modules.NewAuthModule(authUseCases)

	// Assign a router to the server
	cfg.Server.Router = mux.NewRouter()

	// Register routes
	cfg.Server.RegisterModules(healthModule, authModule)

	slog.Info("server is running on", slog.Int("port", cfg.Server.Port))

	err = cfg.Server.Run(cfg)
	if err != nil {
		panic(err)
	}
}
