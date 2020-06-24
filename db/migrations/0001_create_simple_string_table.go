package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createSimpleStringTableMigration = &Migration{
	Number: 1,
	Name:   "Create simple string table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
			CREATE TABLE simple_string(
				id BIGSERIAL NOT NULL,
				data text PRIMARY KEY NOT NULL,
				timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
			);
		`

		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create processing transfer transactions table")
	},
}

func init() {
	Migrations = append(Migrations, createSimpleStringTableMigration)
}
