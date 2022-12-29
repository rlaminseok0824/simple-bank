//if validator(화폐 단위)가 다양해질 때를 대비하여 만든 파일

package api

import (
	"example.com/db/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		//check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
