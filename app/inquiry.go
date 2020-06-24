package app

import (
	"github.com/oatsaysai/simple-go-http-server/db"
	e "github.com/oatsaysai/simple-go-http-server/error"
	log "github.com/oatsaysai/simple-go-http-server/log"
)

type CheckStringIsAlreadyExistParams struct {
	Data string `json:"data" validate:"required"`
}

type CheckStringIsAlreadyExistResponse struct {
	Exist bool `json:"exist"`
}

func (ctx *Context) CheckStringIsAlreadyExist(params *CheckStringIsAlreadyExistParams) (*CheckStringIsAlreadyExistResponse, error) {
	return checkStringIsAlreadyExist(
		ctx.Logger,
		params,
		ctx.DB.CheckStringIsAlreadyExist,
	)
}

func checkStringIsAlreadyExist(
	logger log.Logger,
	params *CheckStringIsAlreadyExistParams,
	checkStringIsAlreadyExist db.CheckStringIsAlreadyExistFunc,
) (*CheckStringIsAlreadyExistResponse, error) {

	if err := validateInput(params); err != nil {
		logger.Errorf("validateInput error Code : %d, Message : %s", e.InputValidationError, err.Error())
		return nil, &e.UserError{
			Code:    e.InputValidationError,
			Message: err.Error(),
		}
	}

	return &CheckStringIsAlreadyExistResponse{
		Exist: checkStringIsAlreadyExist(params.Data),
	}, nil

}

type GetAllStringParams struct{}

type GetAllStringResponse struct {
	StringList []db.StringData `json:"string_list"`
}

func (ctx *Context) GetAllString(params *GetAllStringParams) (*GetAllStringResponse, error) {
	return getAllString(
		ctx.Logger,
		ctx.DB.GetAllString,
	)
}

func getAllString(
	logger log.Logger,
	getAllString db.GetAllStringFunc,
) (*GetAllStringResponse, error) {

	strList, err := getAllString()
	if err != nil {
		logger.Errorf("Get string list from DB error: %s", err.Error())
		return nil, &e.InternalError{
			Code:    e.InquiryDBError,
			Message: "Inquiry database error",
		}
	}

	return &GetAllStringResponse{
		StringList: strList,
	}, nil

}
