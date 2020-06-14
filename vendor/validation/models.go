package validation

import (
	"regexp"
)

var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
var phoneRe = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-9]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

type Validator interface {
	Validate(interface{}) (bool, error)
}

type DefaultValidator struct {
}

type StringValidator struct {
	Min int
	Max int
}

type NumberValidator struct {
	Min int
	Max int
}

type EmailValidator struct {
}

type PhoneValidator struct {
}
