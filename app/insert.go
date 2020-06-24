package app

import (
	"github.com/oatsaysai/simple-go-http-server/db"
	e "github.com/oatsaysai/simple-go-http-server/error"
	log "github.com/oatsaysai/simple-go-http-server/log"
)

type InsertStringParams struct {
	Data string `json:"data" validate:"required"`
}

type InsertStringResponse struct {
	ID int64 `json:"id"`
}

func (ctx *Context) InsertString(params *InsertStringParams) (*InsertStringResponse, error) {
	return insertString(
		ctx.Config,
		ctx.Logger,
		params,
		ctx.DB.InsertString,
	)
}

func insertString(
	appConfig *Config,
	logger log.Logger,
	params *InsertStringParams,
	insertStringToDB db.InsertString,
) (*InsertStringResponse, error) {

	if err := validateInput(params); err != nil {
		logger.Errorf("validateInput error Code : %d, Message : %s", e.InputValidationError, err.Error())
		return nil, &e.UserError{
			Code:    e.InputValidationError,
			Message: err.Error(),
		}
	}

	// check params with some business requirement
	if len(params.Data) > appConfig.MaxStringLength {
		return nil, &e.UserError{
			Code:    e.DataIsTooLong,
			Message: "Data is too long",
		}
	}

	stringID, err := insertStringToDB(params.Data)
	if err != nil {
		logger.Errorf("insert string to DB error: %s", err.Error())
		return nil, &e.InternalError{
			Code:    e.InsertDBError,
			Message: "Data is already exist",
		}
	}

	return &InsertStringResponse{
		ID: stringID,
	}, nil
}
