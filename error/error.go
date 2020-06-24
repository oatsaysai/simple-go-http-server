package error

type ValidationError struct {
	Code    uint64      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

type UserError struct {
	Code       uint64      `json:"code"`
	Message    string      `json:"message"`
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
}

func (e *UserError) Error() string {
	return e.Message
}

type InternalError struct {
	Code    uint64      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *InternalError) Error() string {
	return e.Message
}
