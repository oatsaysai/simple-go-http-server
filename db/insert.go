package db

import (
	"context"
)

type InsertString func(
	data string,
) (stringID int64, err error)

func (db *DB) InsertString(
	data string,
) (stringID int64, err error) {

	txn, err := db.DB.Begin(context.Background())
	if err != nil {
		db.logger.Errorf("error starting transaction: %+v", err)
		return 0, err
	}

	err = txn.QueryRow(context.Background(),
		`
			INSERT INTO simple_string (
				"data"
			)
			VALUES ($1)
			RETURNING "id";
		`,
		data,
	).Scan(&stringID)
	if err != nil {
		db.logger.Errorf("error insert string to DB: %+v", err)
		txn.Rollback(context.Background())
		return 0, err
	}

	err = txn.Commit(context.Background())
	if err != nil {
		txn.Rollback(context.Background())
		return 0, err
	}

	return stringID, err
}
