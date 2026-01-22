package infrastructure

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"taskflow/domain/entities"
	"taskflow/domain/usecases"
	"taskflow/infrastructure/datastore/repositories"
	"taskflow/infrastructure/router/modules"

	"github.com/gorilla/mux"
)

func SetupModules(r *mux.Router, config entities.Config) error {
	// Repository repoSettings
	repoSettings, err := repositories.NewRepositorySettings(config)
	if err != nil {
		return errors.Join(errors.New("failed to create settings repository"), err)
	}

	// Repositories
	authRepository := repositories.NewAuthRepository(repoSettings)

	// Use Cases
	authUseCases := usecases.NewAuthUseCases(authRepository, config.Paseto.SecurityKey)

	// Modules
	authModule := modules.NewAuthModule(authUseCases)

	_, _ = authModule.Setup(r)

	// Home URL handler returns the current server time
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serverTime, err := repoSettings.ServerTime(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = fmt.Fprintf(w, "%v", serverTime.UTC().Unix())

		if err != nil {
			log.Println(err)
		}
	})

	return nil
}
