package app

import (
	logg "log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/oatsaysai/simple-go-http-server/db"
	e "github.com/oatsaysai/simple-go-http-server/error"
	log "github.com/oatsaysai/simple-go-http-server/log"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	trans    ut.Translator
	validate *validator.Validate
)

func init() {
	en := en.New()
	uni = ut.New(en, en)

	// this is usually known or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	translator, found := uni.GetTranslator("en")
	if !found {
		logg.Fatal("translator not found")
	}

	validate = validator.New()

	if err := en_translations.RegisterDefaultTranslations(validate, translator); err != nil {
		logg.Fatalf("register default translations error : %v", err)
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	trans = translator
}

func validateInput(input interface{}) *e.ValidationError {
	err := validate.Struct(input)
	if err != nil {
		messages := make([]string, 0)
		for _, e := range err.(validator.ValidationErrors) {
			// fmt.Println(e.Translate(trans))
			messages = append(messages, e.Translate(trans))
		}
		errMessage := strings.Join(messages, ", ")
		return &e.ValidationError{
			Code:    e.InputValidationError,
			Message: errMessage,
		}
	}
	return nil
}

type App struct {
	Logger log.Logger
	Config *Config
	DB     *db.DB
}

func New(logger log.Logger) (app *App, err error) {
	app = &App{
		Logger: logger,
	}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	dbConfig, err := db.InitConfig()
	if err != nil {
		return nil, err
	}

	app.DB, err = db.New(dbConfig, logger)
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *App) Close() error {
	if err := app.DB.Close(); err != nil {
		return err
	}
	return nil
}
