package formdata

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
)

func Decode(r *http.Request, v any) error {
	err := r.ParseMultipartForm(10 << 32)
	if err != nil {
		return err
	}

	reflectionTypePtr := reflect.TypeOf(v)
	reflectionValuePtr := reflect.ValueOf(v)

	if reflectionTypePtr.Kind() != reflect.Pointer {
		return fmt.Errorf("Expected kind 'pointer', got kind %s", reflectionTypePtr.Kind().String())
	}

	reflectionType := reflectionTypePtr.Elem()
	reflectionValue := reflectionValuePtr.Elem()

	if reflectionType.Kind() != reflect.Struct {
		return fmt.Errorf("Expected kind 'struct', got kind %s", reflectionType.Kind().String())
	}

	for i := 0; i < reflectionType.NumField(); i++ {
		field := reflectionType.Field(i)
		fieldValue := reflectionValue.Field(i)
		var formValue string
		var formFile []*multipart.FileHeader
		if len(r.MultipartForm.Value[field.Tag.Get("form")]) != 0 {
			formValue = r.MultipartForm.Value[field.Tag.Get("form")][0]
		}
		formFile = r.MultipartForm.File[field.Tag.Get("form")]

		switch field.Type.Kind() {
		case reflect.String:
			fieldValue.SetString(formValue)
		case reflect.Int, reflect.Int32, reflect.Int64:
			i, _ := strconv.Atoi(formValue)
			fieldValue.SetInt(int64(i))
		case reflect.Float32, reflect.Float64:
			f, _ := strconv.ParseFloat(formValue, 64)
			fieldValue.SetFloat(f)
		case reflect.Slice:
			if reflect.TypeOf([]*multipart.FileHeader{}) == field.Type {
				fieldValue.Set(reflect.ValueOf(formFile))
			}
		}
	}

	return nil
}
