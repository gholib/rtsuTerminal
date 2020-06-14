package validation

import (
	"fmt"
	"reflect"
	"strings"
)

const tagName = "validate"

func Validate(object interface{}) []error {
	errs := []error{}
	value := reflect.ValueOf(object)
  
	for i := 0; i < value.NumField(); i++ {
	  tag := value.Type().Field(i).Tag.Get(tagName)
	  if tag == "" || tag == "-" {
		continue
	  }
	  validator := getValidatorFromTag(tag)
	  valid, err := validator.Validate(value.Field(i).Interface())
	  if !valid && err != nil {
		errs = append(errs, fmt.Errorf("%s %s", value.Type().Field(i).Name, err.Error()))
	  }
	}
  
	return errs
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	rule, ok := rules[args[0]]
	if !ok {
		return DefaultValidator{}
	}
	validator := rule(args)

	return validator
}