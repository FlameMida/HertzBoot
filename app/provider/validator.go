package provider

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"regexp"
	"strings"
)

type BindError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *BindError) Error() string {
	if e.Msg != "" {
		e.Msg = strings.Replace(e.Msg, "\"", "`", -1)
		return fmt.Sprintf("%sparams: ${%s},message: %v", e.ErrType, e.FailField, e.Msg)
	}
	return fmt.Sprintf("%sparams: ${%s},invalid", e.ErrType, e.FailField)
}

type ValidateError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *ValidateError) Error() string {
	if e.Msg != "" {
		e.Msg = strings.Replace(e.Msg, "\"", "`", -1)
		return fmt.Sprintf("%sparams: ${%s},message: %v", e.ErrType, e.FailField, e.Msg)
	}
	return fmt.Sprintf("%sparams: ${%s},invalid", e.ErrType, e.FailField)
}

func init() {
	binding.MustRegValidateFunc("dateTimes", func(args ...interface{}) error {
		if len(args) > 0 {
			dateTime, _ := args[0].(string)
			var re = regexp.MustCompile(`(?m)^((?:(\d{4}-\d{2}-\d{2}) (\d{2}:\d{2}:\d{2}(?:\.\d+)?))?)$`)
			match := re.MatchString(dateTime)
			if !match {
				return fmt.Errorf("dateTime format invalid")
			}
		}
		return nil
	})
	binding.MustRegValidateFunc("dateTime", func(args ...interface{}) error {
		if len(args) > 0 {
			dateTime, _ := args[0].(string)
			var re = regexp.MustCompile(`(?m)^((?:(\d{4}-\d{2}-\d{2}) (\d{2}:\d{2}(?:\.\d+)?))?)$`)
			match := re.MatchString(dateTime)
			if !match {
				return fmt.Errorf("dateTime format invalid")
			}
		}
		return nil
	})

	CustomBindErrFunc := func(failField string, msg string) error {
		err := BindError{
			ErrType:   "BindingError-",
			FailField: failField,
			Msg:       msg,
		}

		return &err
	}

	CustomValidateErrFunc := func(failField string, msg string) error {
		err := ValidateError{
			ErrType:   "ValidateError-",
			FailField: failField,
			Msg:       msg,
		}

		return &err
	}

	binding.SetErrorFactory(CustomBindErrFunc, CustomValidateErrFunc)
}
