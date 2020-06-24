package app

import (
	"errors"
	"testing"

	"github.com/oatsaysai/simple-go-http-server/db"
)

func TestCheckStringIsAlreadyExist(t *testing.T) {

	expected := true

	logger, err := createLoggerForTest()
	if err != nil {
		t.Fatalf("create logger err: %+v", err)
	}

	params := &CheckStringIsAlreadyExistParams{
		Data: "test",
	}

	checkStringIsAlreadyExistInDB := func(data string) bool {
		return expected
	}

	res, err := checkStringIsAlreadyExist(
		logger,
		params,
		checkStringIsAlreadyExistInDB,
	)
	if err != nil {
		t.Errorf("CheckStringIsAlreadyExist err: %+v", err)
	}
	if res.Exist != expected {
		t.Errorf(
			"Invalid ID Expected: %v, got: %v",
			expected,
			res.Exist,
		)
	}
}

func TestCheckStringIsAlreadyExist_With_Invalid_Params(t *testing.T) {

	expected := true

	logger, err := createLoggerForTest()
	if err != nil {
		t.Fatalf("create logger err: %+v", err)
	}

	params := &CheckStringIsAlreadyExistParams{
		Data: "",
	}

	checkStringIsAlreadyExistInDB := func(data string) bool {
		return expected
	}

	_, err = checkStringIsAlreadyExist(
		logger,
		params,
		checkStringIsAlreadyExistInDB,
	)
	if err == nil {
		t.Errorf("Invalid params should error")
	}
}

func TestGetAllString(t *testing.T) {

	logger, err := createLoggerForTest()
	if err != nil {
		t.Fatalf("create logger err: %+v", err)
	}

	getAllStringInDB := func() ([]db.StringData, error) {
		return nil, nil
	}

	_, err = getAllString(
		logger,
		getAllStringInDB,
	)
	if err != nil {
		t.Errorf("GetAllString err: %+v", err)
	}
}

func TestGetAllString_With_DB_Error(t *testing.T) {

	logger, err := createLoggerForTest()
	if err != nil {
		t.Fatalf("create logger err: %+v", err)
	}

	getAllStringInDB := func() ([]db.StringData, error) {
		return nil, errors.New("DB error")
	}

	_, err = getAllString(
		logger,
		getAllStringInDB,
	)
	if err == nil {
		t.Errorf("DB error should error")
	}
}
