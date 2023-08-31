package validation

import (
	"github.com/go-playground/validator/v10"
)

func MustValid(fl validator.FieldLevel) bool {
	f := fl.Field()
	meth := f.MethodByName("Valid")
	if meth.IsZero() {
		panic("missing method Valid with return bool")
	}
	returns := meth.Call(nil)
	for _, v := range returns {
		return v.Bool()
	}
	return false
}
