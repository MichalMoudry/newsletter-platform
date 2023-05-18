package database

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

// This function starts a database transaction.
func BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return dbContext.BeginTxx(ctx, nil)
}

// This function ends a specific database transaction.
func EndTransaction(transaction *sqlx.Tx, err error) error {
	if err != nil {
		if rollbackErr := transaction.Rollback(); rollbackErr != nil {
			err = errors.Join(rollbackErr, err)
		}
	}
	if commitErr := transaction.Commit(); commitErr != nil {
		err = errors.Join(commitErr, err)
	}
	return err
}
