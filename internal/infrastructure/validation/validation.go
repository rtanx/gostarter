package validation

import (
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type mainValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = NewValidator()

func NewValidator() *mainValidator {
	return new(mainValidator)
}
func (v *mainValidator) Engine() any {
	v.lazyInit()
	return v.validate
}
func (v *mainValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		v.validate.RegisterTagNameFunc(jsonKeyAsKeyTag)
		RegisterValidations(v.validate)
	})
}

func RegisterValidations(v *validator.Validate) {
	// Register custom validator here
	MustRegisterValidation(v, "must_valid", MustValid)
}

func MustRegisterValidation(v *validator.Validate, alias string, handler validator.Func) {
	if err := v.RegisterValidation(alias, handler); err != nil {
		panic(err)
	}
}

func (v *mainValidator) ValidateStruct(obj any) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

func kindOfData(data any) reflect.Kind {
	value := reflect.ValueOf(data)
	valType := value.Kind()

	if valType == reflect.Ptr {
		valType = value.Elem().Kind()
	}
	return valType
}

func jsonKeyAsKeyTag(fl reflect.StructField) string {
	if name := strings.SplitN(fl.Tag.Get("json"), ",", 2)[0]; name != "-" {
		return name
	}
	return ""
}
