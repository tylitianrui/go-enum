package enum

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const tagName = "enum"
const Scrambled = "qw@s5sadv+_)IJHUwe3e67ytressaqwef@$^%$#WDFTY&*IJHBVCazxcv3%bv&cde*ryuhb"

const (
	eScrambled = iota
	eInt
	eUint
	eFloat
	eString
)

type Enum struct {
	data interface{}
}

func New(data interface{}) *Enum {
	return &Enum{
		data: data,
	}
}

func (e *Enum) Verify() error {
	return vetTag(e.data)
}

func vetTag(data interface{}) error {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Struct {
		return errors.New("ParamTypeError")
	}

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)

		if content, err := tagSyntax1(tag); err != nil {
			return err
		} else {
			_, c := convert(reflect.ValueOf(data).Field(i))
			if !strings.Contains(content, c) {
				return NewEnumError(fmt.Sprintf("%s not in %s", c, tag))
			}
		}

	}
	return nil
}

func tagSyntax1(tag string) (string, error) {
	if !strings.HasPrefix(tag, "[") {
		return "", errors.New("语法错误")
	}
	if !strings.HasSuffix(tag, "]") {
		return "", errors.New("语法错误")
	}
	content := tag[1 : len(tag)-1]
	if strings.Contains(content, "[") || strings.Contains(content, "]") {
		return "", errors.New("语法错误")
	}
	return content, nil
}

func convert(data reflect.Value) (int, string) {
	switch data.Kind() {
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:

		return eInt, fmt.Sprintf("%d", data.Int())

	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return eUint, fmt.Sprintf("%d", data.Uint())

	case reflect.String:
		return eString, data.String()
	case reflect.Float32, reflect.Float64:
		return eFloat, fmt.Sprintf("%f", data.Float())

	default:
		return eScrambled, Scrambled

	}

}

type EnumError struct {
	err string
}

func NewEnumError(err string) *EnumError {
	return &EnumError{
		err: err,
	}

}

func (e *EnumError) Error() string {
	return e.err
}
