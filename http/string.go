package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/oatsaysai/simple-go-http-server/app"
	e "github.com/oatsaysai/simple-go-http-server/error"
)

var StringRoutes = Routes{
	Route{
		Name:        "Insert",
		Path:        "/insert",
		Method:      "POST",
		HandlerFunc: (&API{}).InsertString,
	},
	Route{
		Name:        "CheckStringIsAlreadyExist",
		Path:        "/check_exist",
		Method:      "POST",
		HandlerFunc: (&API{}).CheckStringIsAlreadyExist,
	},
	Route{
		Name:        "ListString",
		Path:        "/list_string",
		Method:      "GET",
		HandlerFunc: (&API{}).ListString,
	},
}

func (api *API) InsertString(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input app.InsertStringParams

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return &e.UserError{
			Code:       e.InvalidJSONString,
			Message:    "Invalid JSON string",
			StatusCode: http.StatusBadRequest,
		}
	}

	resData, err := ctx.InsertString(&input)
	if err != nil {
		return err
	}

	resp := &Response{
		Code:    0,
		Message: "",
		Data:    resData,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	ctx.Logger.Debugf("Response Success : %+v", resp)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	return err
}

func (api *API) CheckStringIsAlreadyExist(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input app.CheckStringIsAlreadyExistParams

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return &e.UserError{
			Code:       e.InvalidJSONString,
			Message:    "Invalid JSON string",
			StatusCode: http.StatusBadRequest,
		}
	}

	resData, err := ctx.CheckStringIsAlreadyExist(&input)
	if err != nil {
		return err
	}

	resp := &Response{
		Code:    0,
		Message: "",
		Data:    resData,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	ctx.Logger.Debugf("Response Success : %+v", resp)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	return err
}

func (api *API) ListString(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input app.GetAllStringParams

	// defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	return err
	// }

	// if err := json.Unmarshal(body, &input); err != nil {
	// 	return &e.UserError{
	// 		Code:       e.InvalidJSONString,
	// 		Message:    "Invalid JSON string",
	// 		StatusCode: http.StatusBadRequest,
	// 	}
	// }

	resData, err := ctx.GetAllString(&input)
	if err != nil {
		return err
	}

	resp := &Response{
		Code:    0,
		Message: "",
		Data:    resData,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	ctx.Logger.Debugf("Response Success : %+v", resp)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	return err
}
