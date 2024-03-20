package errors

type Error interface {
	error
	GetCode() int
}

type AppMessage struct {
	code    int
	message string
	options Options
}

type Options struct {
	Headers []string
}

func NewMessage(code int, message string, options ...Options) error {
	if len(options) == 1 {
		return &AppMessage{code: code, message: message, options: options[0]}
	}

	return &AppMessage{code: code, message: message}
}

func NewClientError(message string) error {
	return NewMessage(400, message)
}

func NewServerError(message string) error {
	return NewMessage(500, message)
}

var EmptyEmail = NewClientError("empty email")
var EmptyFirstName = NewClientError("empty first name")
var EmptyLastName = NewClientError("empty last name")
var InvalidEmail = NewClientError("invalid email")
var EmptyPassword = NewClientError("empty password")
var UserNotFound = NewClientError("user not found")
var WrongPassword = NewClientError("wrong password")
var NoRowsAffected = NewClientError("no rows affected")
var EmailIsAlreadyTaken = NewClientError("email is already taken")
var PayloadWrongLength = NewClientError("payload shall neither be empty nor exceed 255 KB")
var StringShallContainOnlyPrintableCharacters = NewClientError("string shall contain only printable characrters")
var EmptyJobID = NewClientError("empty job id")
var JobProcessingInProgress = NewClientError("your job is being processed")
var JobProcessingHasFailed = NewClientError("job processing has failed")
var JobNotFound = NewClientError("job not found")
var NotAuthorized = NewMessage(403, "not authorized")
var EmptyAmount = NewClientError("empty amount")
var WrongAmount = NewClientError("wrong amount")
var MinimumPayment10Dollars = NewClientError("minimum payment is 10 dollars")
var YourBalanceIsNotEnoughForNewJobs = NewClientError("your balance is nor enough for new jobs")

func (ce AppMessage) GetCode() int {
	return ce.code
}

func (ce AppMessage) Error() string {
	return ce.message
}
