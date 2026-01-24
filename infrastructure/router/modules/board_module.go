package modules

import (
	"log/slog"
	"net/http"
	"taskflow/domain/usecases"
	"taskflow/infrastructure/router"

	"github.com/gorilla/mux"
)

type boardModule struct {
	boardUseCases usecases.BoardUseCases
	name          string
	path          string
}

func NewBoardModule(boardUseCases usecases.BoardUseCases) router.Module {
	return boardModule{
		boardUseCases: boardUseCases,
		name:          "Boards",
		path:          "/boards",
	}
}

func (b boardModule) Name() string {
	return b.name
}

func (b boardModule) Path() string {
	return b.path
}

func (b boardModule) Setup(r *mux.Router) ([]router.RouteDefinition, *mux.Router) {
	defs := []router.RouteDefinition{
		{
			Path:        "/list",
			Description: "List boards",
			Handler:     b.list,
			HttpMethods: []string{http.MethodGet},
		},
		{
			Path:        "/create",
			Description: "Create a new board",
			Handler:     b.create,
			HttpMethods: []string{http.MethodPost},
		},
	}

	for _, d := range defs {
		r.HandleFunc(b.path+d.Path, d.Handler).Methods(d.HttpMethods...)
	}

	return defs, r
}

func (b boardModule) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	boards, err := b.boardUseCases.GetBoards(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get boards", "cause", err)
		router.WriteInternalError(w)
		return
	}

	err = router.Write(w, boards)
	if err != nil {
		slog.ErrorContext(ctx, "failed to write response", "cause", err)
	}
}

func (b boardModule) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	boards, err := b.boardUseCases.GetBoards(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get boards", "cause", err)
		router.WriteInternalError(w)
		return
	}

	err = router.Write(w, boards)
	if err != nil {
		slog.ErrorContext(ctx, "failed to write response", "cause", err)
	}
}
