package errs

import (
	"fmt"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

type Errs struct {
	HTTPStatusCode int    `json:"-"`
	Code           string `json:"code"`
	Msg            string `json:"message"`
	Description    string `json:"description"`
}

var (
	SystemErr = "System Error"
)

func New(httpStatusCode int, errorCode, msg, description string) error {
	return &Errs{
		HTTPStatusCode: httpStatusCode,
		Code:           errorCode,
		Msg:            msg,
		Description:    description,
	}
}

func (e *Errs) Error() string {
	return fmt.Sprintf("code:%s, msg:%s, description:%s\n", e.Code, e.Msg, e.Description)
}
