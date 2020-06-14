package validation

import (
	"fmt"
	"strings"
)

type fn func(args []string) Validator

var rules = map[string]fn{
	"number": func(args []string) Validator {
		validator := NumberValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	},
	"string": func(args []string) Validator {
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	},
	"email": func(args []string) Validator {
		return EmailValidator{}
	},
	"phone": func(args []string) Validator {
		return PhoneValidator{}
	},
}
