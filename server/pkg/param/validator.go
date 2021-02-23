package param

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type R map[string]map[string]string

type Params interface {
	Rule() R
}

func Validate(errs validator.ValidationErrors, params Params) error {
	for _, err := range errs {
		p, ok := params.Rule()[err.Field()][err.ActualTag()]
		if ok && p != "" {
			return errors.New(p)
		} else {
			return errors.New(err.Error())
		}
	}
	return nil
}