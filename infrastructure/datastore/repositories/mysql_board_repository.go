package repositories

import (
	"context"
	"database/sql"
	"errors"
	"taskflow/domain/entities"
	"taskflow/infrastructure/datastore"
)

type boardRepository struct {
	conn func() *sql.DB
}

func NewBoardRepository(settings datastore.RepositorySettings) datastore.BoardRepository {
	return boardRepository{
		conn: settings.Connection,
	}
}

func (r boardRepository) GetBoards(ctx context.Context) ([]entities.Board, error) {
	const query = `
	SELECT id, 
	       uuid, 
	       title,
	       description,
	       u.id,
	       u.uuid,
	       u.email,
	       u.created_at,
	       u.modified_at,
	       status_code, 
	       modified_at, 
	       created_at 
	FROM boards 
	    LEFT JOIN users u ON u.id = user_id		
	`
	boards := make([]entities.Board, 0)
	rows, err := r.conn().QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Join(entities.ErrExecuteQuery, err)
	}
	defer rows.Close()

	for rows.Next() {
		var board entities.Board
		err = rows.Scan(
			board.ID,
			board.Title,
			board.Description,
			board.CreatedBy.ID,
			board.CreatedBy.UUID,
			board.CreatedBy.Email,
			board.CreatedBy.CreatedAt,
			board.CreatedBy.ModifiedAt,
			board.StatusCode,
			board.ModifiedAt,
			board.CreatedAt,
		)
		if err != nil {
			return nil, errors.Join(entities.ErrScan, err)
		}
		boards = append(boards, board)
	}

	return boards, nil
}
