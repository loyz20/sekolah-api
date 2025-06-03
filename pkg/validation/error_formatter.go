package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	res := map[string]string{}

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			field := e.Field()
			tag := e.Tag()

			var msg string
			switch tag {
			case "required":
				msg = fmt.Sprintf("%s wajib diisi", field)
			case "email":
				msg = fmt.Sprintf("%s harus berupa email yang valid", field)
			case "len":
				msg = fmt.Sprintf("%s harus sepanjang %s karakter", field, e.Param())
			case "numeric":
				msg = fmt.Sprintf("%s harus berupa angka", field)
			default:
				msg = fmt.Sprintf("%s tidak valid", field)
			}

			res[field] = msg
		}
	}

	return res
}
