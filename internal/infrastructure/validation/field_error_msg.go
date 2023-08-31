package validation

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/go-playground/validator/v10"
	"github.com/rtanx/gostarter/internal/infrastructure/logger"
)

type templateArgs struct {
	Param      []string
	FieldValue any
}
type FieldValidationErr struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

type FieldValidationErrs struct {
	Errors []*FieldValidationErr `json:"field_errors"`
}

var msgByTag map[string]string = map[string]string{
	"required":    `Required`,
	"required_if": `Required`,
	"email":       `Please enter a valid email address`,
	"min":         `must be at least {{.Param}}`,
	"oneof":       `one of{{range .Param}} {{.}}{{end}}`,
	// other validation message goes here
}

var msgByDataType map[string]map[string]string = map[string]map[string]string{
	"string": {
		"min": fmt.Sprintf(`%s %s`, msgByTag["min"], `characters`),
	},
	// "int": {},
	// "struct":{},
	// etc...
}

var msgByField map[string]map[string]string = map[string]map[string]string{
	// other validation message goes here
	// e.g
	// "Login.Email":{
	// "max": `some useful message...`
	// }
}

func generateMsg(fe validator.FieldError) string {
	tag := fe.Tag()
	ftype := fe.Type().String()
	fnamespce := fe.StructNamespace()

	msg, ok := msgByField[fnamespce][tag]
	if !ok {
		msg, ok = msgByDataType[ftype][tag]
		if !ok {
			msg, ok = msgByTag[tag]
			if !ok {
				// return default error message provided by go-validator if message by field, data type, and tags does not exists
				return fe.Error()
			}
		}
	}
	args := templateArgs{
		Param:      strings.Split(fe.Param(), " "),
		FieldValue: fe.Value(),
	}
	t := template.Must(template.New("msg").Parse(msg))
	out := new(strings.Builder)
	if err := t.Execute(out, args); err != nil {
		logger.Error("error while generating message from template", logger.Any("error", err))
		return fe.Error()
	}
	return out.String()

}

func FailedValidationMapper(ve validator.ValidationErrors) FieldValidationErrs {
	resp := make([]*FieldValidationErr, len(ve))
	for i, fe := range ve {
		namespce := strings.SplitN(fe.Namespace(), ".", 2)
		var fname string
		if len(namespce) <= 1 {
			fname = namespce[0]
		} else {
			fname = namespce[1]
		}
		msg := generateMsg(fe)

		resp[i] = &FieldValidationErr{
			Field: fname,
			Msg:   msg,
		}
	}
	return FieldValidationErrs{resp}
}
