package error

const _internalCode = 100000

const (
	UnknownError         uint64 = _internalCode + 1
	PanicRecovery        uint64 = _internalCode + 2
	HttpCodeNotSet       uint64 = _internalCode + 3
	InputValidationError uint64 = _internalCode + 4
	InvalidJSONString    uint64 = _internalCode + 5
	InsertDBError        uint64 = _internalCode + 6
	InquiryDBError       uint64 = _internalCode + 7
)
