package app

import (
	"errors"
	"testing"
)

func TestInsertString(t *testing.T) {

	var expectedID int64
	expectedID = 1

	logger, err := createLoggerForTest()
	if err != nil {
		t.Fatalf("create logger err: %+v", err)
	}

	appConfig := &Config{
		MaxStringLength: 10,
	}

	params := &InsertStringParams{
		Data: "test",
	}

	insertStringToDB := func(data string) (stringID int64, err error) {
		return expectedID, nil
	}

	res, err := insertString(
		appConfig,
		logger,
		params,
		insertStringToDB,
	)
	if err != nil {
		t.Errorf("InsertString err: %+v", err)
	}
	if res.ID != expectedID {
		t.Errorf(
			"Invalid ID Expected: %d, got: %d",
			expectedID,
			res.ID,
		)
	}
}

func TestInsertString_With_Invalid_Params(t *testing.T) {

	var expectedID int64
	expectedID = 1

	logger, err := createLoggerForTest()
	if err != nil {
		t.Fatalf("create logger err: %+v", err)
	}

	appConfig := &Config{
		MaxStringLength: 10,
	}

	params := &InsertStringParams{
		Data: "",
	}

	insertStringToDB := func(data string) (stringID int64, err error) {
		return expectedID, nil
	}

	_, err = insertString(
		appConfig,
		logger,
		params,
		insertStringToDB,
	)
	if err == nil {
		t.Errorf("Invalid params should error")
	}
}

func TestInsertString_With_Too_Long_Data(t *testing.T) {

	var expectedID int64
	expectedID = 1

	logger, err := createLoggerForTest()
	if err != nil {
		t.Fatalf("create logger err: %+v", err)
	}

	appConfig := &Config{
		MaxStringLength: 10,
	}

	params := &InsertStringParams{
		Data: "testtesttest",
	}

	insertStringToDB := func(data string) (stringID int64, err error) {
		return expectedID, nil
	}

	_, err = insertString(
		appConfig,
		logger,
		params,
		insertStringToDB,
	)
	if err == nil {
		t.Errorf("Too long data should error")
	}
}

func TestInsertString_InsertDB_Error(t *testing.T) {

	var expectedID int64
	expectedID = 1

	logger, err := createLoggerForTest()
	if err != nil {
		t.Fatalf("create logger err: %+v", err)
	}

	appConfig := &Config{
		MaxStringLength: 10,
	}

	params := &InsertStringParams{
		Data: "test",
	}

	insertStringToDB := func(data string) (stringID int64, err error) {
		return expectedID, errors.New("duplicate key value violates")
	}

	_, err = insertString(
		appConfig,
		logger,
		params,
		insertStringToDB,
	)
	if err == nil {
		t.Errorf("duplicate key case should error")
	}
}
