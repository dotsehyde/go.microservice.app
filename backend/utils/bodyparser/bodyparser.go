package bodyparser

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

func Parse(r *http.Request, v any) error {
	defer r.Body.Close()
	if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") || strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		return readForm(r, v)
	}
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		return readJSON(r, v)
	}
	if strings.HasPrefix(r.Header.Get("Content-Type"), "text/xml") || strings.HasPrefix(r.Header.Get("Content-Type"), "application/xml") {
		return readXML(r, v)
	}
	return nil
}
func readJSON(r *http.Request, v any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(v)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf("body contains incorrect JSON type for field %q at offset %d", unmarshalTypeError.Field, unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case errors.As(err, &invalidUnmarshalError):
			return fmt.Errorf("error unmarshalling json: %s", err.Error())
		default:
			return err
		}
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}
	return nil
}

func readForm(r *http.Request, v any) error {
	elType := reflect.TypeOf(v)
	elValue := reflect.ValueOf(v)
	if elType.Kind() != reflect.Pointer || elValue.Kind() != reflect.Pointer {
		panic("error: a pointer is expected as an argument")
	}
	elemType := elType.Elem()
	elemValue := elValue.Elem()
	r.ParseMultipartForm(0)
	wg := &sync.WaitGroup{}
	for i := 0; i < elemType.NumField(); i++ {
		wg.Add(1)
		go func(i int, r *http.Request, v any) {
			defer wg.Done()
			val, _ := elemType.Field(i).Tag.Lookup("json")
			switch elemValue.Field(i).Kind() {
			case reflect.String:
				elemValue.Field(i).Set(reflect.ValueOf(r.FormValue(val)))

			case reflect.Int:
				formVal, _ := strconv.Atoi(r.FormValue(val))
				elemValue.Field(i).Set(reflect.ValueOf(formVal))
			case reflect.Int8:
				formVal, _ := strconv.ParseInt(r.FormValue(val), 10, 8)
				elemValue.Field(i).Set(reflect.ValueOf(int8(formVal)))
			case reflect.Int16:
				formVal, _ := strconv.ParseInt(r.FormValue(val), 10, 16)
				elemValue.Field(i).Set(reflect.ValueOf(int16(formVal)))
			case reflect.Int32:
				formVal, _ := strconv.ParseInt(r.FormValue(val), 10, 32)
				elemValue.Field(i).Set(reflect.ValueOf(int32(formVal)))
			case reflect.Int64:
				formVal, _ := strconv.ParseInt(r.FormValue(val), 10, 64)
				elemValue.Field(i).Set(reflect.ValueOf(int64(formVal)))

			case reflect.Uint:
				formVal, _ := strconv.ParseUint(r.FormValue(val), 10, 64)
				elemValue.Field(i).Set(reflect.ValueOf(uint64(formVal)))
			case reflect.Uint8:
				formVal, _ := strconv.ParseUint(r.FormValue(val), 10, 8)
				elemValue.Field(i).Set(reflect.ValueOf(uint8(formVal)))
			case reflect.Uint16:
				formVal, _ := strconv.ParseUint(r.FormValue(val), 10, 16)
				elemValue.Field(i).Set(reflect.ValueOf(uint16(formVal)))
			case reflect.Uint32:
				formVal, _ := strconv.ParseUint(r.FormValue(val), 10, 32)
				elemValue.Field(i).Set(reflect.ValueOf(uint32(formVal)))
			case reflect.Uint64:
				formVal, _ := strconv.ParseUint(r.FormValue(val), 10, 64)
				elemValue.Field(i).Set(reflect.ValueOf(uint64(formVal)))

			case reflect.Float32:
				formVal, _ := strconv.ParseFloat(r.FormValue(val), 32)
				elemValue.Field(i).Set(reflect.ValueOf(float32(formVal)))
			case reflect.Float64:
				formVal, _ := strconv.ParseFloat(r.FormValue(val), 64)
				elemValue.Field(i).Set(reflect.ValueOf(float64(formVal)))

			case reflect.Bool:
				formVal, _ := strconv.ParseBool(r.FormValue(val))
				elemValue.Field(i).Set(reflect.ValueOf(formVal))

			case reflect.Slice, reflect.Array:

				if _, ok := elemValue.Field(i).Interface().([]*multipart.FileHeader); ok {
					elemValue.Field(i).Set(reflect.ValueOf(r.MultipartForm.File[val]))
				} else {

					switch elemValue.Field(i).Type().Elem().Kind() {
					case reflect.Int:
						sliceVal := make([]int, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.Atoi(sv)
							sliceVal = append(sliceVal, formVal)
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Int8:
						sliceVal := make([]int8, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseInt(r.FormValue(sv), 10, 8)
							sliceVal = append(sliceVal, int8(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Int16:
						sliceVal := make([]int16, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseInt(r.FormValue(sv), 10, 16)
							sliceVal = append(sliceVal, int16(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Int32:
						sliceVal := make([]int32, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseInt(r.FormValue(sv), 10, 32)
							sliceVal = append(sliceVal, int32(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Int64:
						sliceVal := make([]int64, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseInt(r.FormValue(sv), 10, 64)
							sliceVal = append(sliceVal, int64(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))

					case reflect.Uint:
						sliceVal := make([]uint, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseUint(r.FormValue(sv), 10, 64)
							sliceVal = append(sliceVal, uint(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Uint8:
						sliceVal := make([]uint8, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseUint(r.FormValue(sv), 10, 8)
							sliceVal = append(sliceVal, uint8(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Uint16:
						sliceVal := make([]uint16, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseUint(r.FormValue(sv), 10, 16)
							sliceVal = append(sliceVal, uint16(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Uint32:
						sliceVal := make([]uint32, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseUint(r.FormValue(sv), 10, 32)
							sliceVal = append(sliceVal, uint32(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Uint64:
						sliceVal := make([]uint64, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseUint(r.FormValue(sv), 10, 64)
							sliceVal = append(sliceVal, uint64(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))

					case reflect.Float32:
						sliceVal := make([]float32, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseFloat(r.FormValue(sv), 32)
							sliceVal = append(sliceVal, float32(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					case reflect.Float64:
						sliceVal := make([]float64, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseFloat(r.FormValue(sv), 64)
							sliceVal = append(sliceVal, float64(formVal))
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))

					case reflect.Bool:
						sliceVal := make([]bool, 0, len(r.PostForm[val]))
						for _, sv := range r.PostForm[val] {
							formVal, _ := strconv.ParseBool(r.FormValue(sv))
							sliceVal = append(sliceVal, formVal)
						}
						elemValue.Field(i).Set(reflect.ValueOf(sliceVal))
					default:
						elemValue.Field(i).Set(reflect.ValueOf(r.PostForm[val]))

					}
				}
			case reflect.Interface, reflect.Ptr:
				if _, ok := elemValue.Field(i).Interface().(*multipart.FileHeader); ok {
					_, fh, _ := r.FormFile(val)
					elemValue.Field(i).Set(reflect.ValueOf(fh))
				}
			}
		}(i, r, v)
	}
	wg.Wait()
	return nil
}
func readXML(r *http.Request, v any) error {
	dec := xml.NewDecoder(r.Body)
	err := dec.Decode(v)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single XML value")
	}
	return nil
}
