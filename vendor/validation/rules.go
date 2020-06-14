package validation

import (
	"fmt"
)

func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))

	if l == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	if l < v.Min {
		return false, fmt.Errorf("should be at least %v chars long", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("should be less than %v chars long", v.Max)
	}

	return true, nil
}

func (v NumberValidator) Validate(val interface{}) (bool, error) {
	num := val.(int)
	if num < v.Min {
		return false, fmt.Errorf("should be greater than %v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("should be less than %v", v.Max)
	}

	return true, nil
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
	if !mailRe.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid email address")
	}
	return true, nil
}

func (v PhoneValidator) Validate(val interface{}) (bool, error) {
	if !phoneRe.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid phone")
	}
	return true, nil
}
