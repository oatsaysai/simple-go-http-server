package http

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/oatsaysai/simple-go-http-server/app"
	e "github.com/oatsaysai/simple-go-http-server/error"
	"github.com/oatsaysai/simple-go-http-server/log"
)

type statusCodeRecorder struct {
	http.ResponseWriter
	http.Hijacker
	StatusCode int
}

type Response struct {
	Code    uint64      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *statusCodeRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

type API struct {
	App    *app.App
	Config *Config
}

func New(app *app.App) (api *API, err error) {
	api = &API{App: app}
	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (api *API) Init(router *mux.Router) {
	api.addRoutes(router, StringRoutes, "")
}

// addRoutes is an internal function to add routes to the api app.
func (api *API) addRoutes(router *mux.Router, routes Routes, prefix string) {
	if prefix == "" {
		for _, route := range routes {
			router.
				Handle(route.Path, api.handler(route.HandlerFunc)).
				Methods(route.Method).
				Name(route.Name)
		}
	} else {
		subRouter := router.PathPrefix(prefix).Subrouter()
		for _, route := range routes {
			subRouter.
				Handle(route.Path, api.handler(route.HandlerFunc)).
				Methods(route.Method).
				Name(route.Name)
		}
	}
}

func (api *API) RemoteAddressForRequest(r *http.Request) string {
	return r.RemoteAddr
}

func (api *API) handler(f func(*app.Context, http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)

		beginTime := time.Now()

		hijacker, _ := w.(http.Hijacker)
		w = &statusCodeRecorder{
			ResponseWriter: w,
			Hijacker:       hijacker,
		}

		ctx := api.App.NewContext().WithRemoteAddress(api.RemoteAddressForRequest(r))

		funcName := strings.Replace(r.URL.String(), "/api/", "", -1)

		ctx = ctx.WithLogger(ctx.Logger.WithFields(
			log.Fields{
				"action":         funcName,
				"remote_address": ctx.RemoteAddress,
			},
		))

		logger := ctx.Logger.WithFields(log.Fields{
			"component": "simple-go-http-server",
		})
		if api.Config.LogLevel == "debug" {
			requestDump, err := httputil.DumpRequest(r, true)
			if err != nil {
				logger.Errorf("%+v", err)
			}
			logger.Debugf("(Incoming HTTP)\n%s", string(requestDump))
		}

		logger.WithFields(log.Fields{
			"tag": "HTTP.Incoming",
		}).Infof("Incoming HTTP")

		w.Header().Set("Content-Type", "application/json")

		defer func() {
			statusCode := w.(*statusCodeRecorder).StatusCode
			duration := time.Since(beginTime)

			logger.WithFields(log.Fields{
				"duration":    duration.Milliseconds(),
				"status_code": statusCode,
				"tag":         "HTTP.Outgoing",
			}).Infof("%s %s (Outgoing)", r.Method, r.URL.RequestURI())
		}()

		defer func() {
			if r := recover(); r != nil {
				logger.Errorf("%v: %s", r, debug.Stack())
				resData, err := json.Marshal(&Response{
					Code:    e.PanicRecovery,
					Message: "Internal server error",
				})
				if err != nil {
					http.Error(w, "internal server error", http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				_, err = w.Write(resData)
			}
		}()

		if err := f(ctx, w, r); err != nil {
			if verr, ok := err.(*e.ValidationError); ok {
				data, err := json.Marshal(&Response{
					Code:    verr.Code,
					Message: verr.Message,
					Data:    verr.Data,
				})
				if err == nil {
					w.WriteHeader(http.StatusBadRequest)
					_, err = w.Write(data)
				}
				if err != nil {
					logger.Errorf("%+v", err)
				}
			} else if ierr, ok := err.(*e.InternalError); ok {
				data, err := json.Marshal(&Response{
					Code:    ierr.Code,
					Message: ierr.Message,
					Data:    ierr.Data,
				})
				if err == nil {
					w.WriteHeader(http.StatusInternalServerError)
					_, err = w.Write(data)
				}
				if err != nil {
					logger.Errorf("%#v", err)
				}
			} else if uerr, ok := err.(*e.UserError); ok {
				data, err := json.Marshal(&Response{
					Code:    uerr.Code,
					Message: uerr.Message,
					Data:    uerr.Data,
				})
				if err == nil {
					if uerr.StatusCode == 0 {
						w.WriteHeader(http.StatusBadRequest)
					} else {
						w.WriteHeader(uerr.StatusCode)
					}
					_, err = w.Write(data)
				}
				if err != nil {
					logger.Errorf("%#v", err)
				}
			} else {
				logger.Errorf("%#v", err)
			}
			logger.Errorf("Response failure: %#v", err)
		}

		statusCode := w.(*statusCodeRecorder).StatusCode
		if statusCode == 0 {
			logger.Errorf("return HTTP status code not set, responding with 500 internal server error")
			resData, err := json.Marshal(&Response{
				Code:    e.HttpCodeNotSet,
				Message: "Internal server error",
			})
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write(resData)
		}
	})
}
