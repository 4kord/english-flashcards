package maindb

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	// bool - commit changes if failed
	ExecTx(ctx context.Context, fn func(Querier) (bool, error)) error
}

type store struct {
	Querier
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &store{
		db:      db,
		Querier: New(db),
	}
}

func (store *store) ExecTx(ctx context.Context, fn func(Querier) (bool, error)) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	var commit bool

	commit, err = fn(q)
	if err != nil {
		if commit {
			if cmErr := tx.Commit(); cmErr != nil {
				return fmt.Errorf("tx err: %v, cm err: %v", err, cmErr)
			}
		} else {
			if rbErr := tx.Rollback(); rbErr != nil {
				return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		}

		return err
	}

	return tx.Commit()
}
