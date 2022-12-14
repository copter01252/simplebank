package api

import (
	"github.com/copter01252/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fidleLevel validator.FieldLevel) bool {
	if currency, ok := fidleLevel.Field().Interface().(string); ok {
		//check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
