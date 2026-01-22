package usecases

import (
	"context"
	"taskflow/domain/entities"
	"taskflow/infrastructure/datastore"
)

type BoardUseCases struct {
	repository datastore.BoardRepository
}

func NewBoardUseCases(repository datastore.BoardRepository) BoardUseCases {
	return BoardUseCases{
		repository: repository,
	}
}

func (b BoardUseCases) GetBoards(ctx context.Context) ([]entities.Board, error) {
	return b.GetBoards(ctx)
}
