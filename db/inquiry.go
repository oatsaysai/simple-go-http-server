package db

import (
	"context"
)

type CheckStringIsAlreadyExistFunc func(data string) bool

func (db *DB) CheckStringIsAlreadyExist(data string) bool {
	var exists bool

	err := db.DB.QueryRow(context.Background(),
		`
		SELECT EXISTS(SELECT data FROM simple_string WHERE data = $1)
		`, data).Scan(&exists)
	if err != nil {
		db.logger.Errorf("%+v", err)
		return false
	}

	return exists
}

type StringData struct {
	ID   int64  `json:"id"`
	Data string `json:"data"`
}

type GetAllStringFunc func() ([]StringData, error)

func (db *DB) GetAllString() ([]StringData, error) {
	strList := []StringData{}
	rows, err := db.DB.Query(context.Background(), `
		SELECT "id", "data" 
		FROM simple_string
	`)
	defer rows.Close()
	for rows.Next() {
		var str StringData
		err = rows.Scan(&str.ID, &str.Data)
		strList = append(strList, str)
	}

	if err != nil {
		db.logger.Errorf("%+v", err)
		return nil, err
	}
	return strList, nil
}
