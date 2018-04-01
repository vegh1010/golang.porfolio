package validator

import (
	"math"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"errors"
	"gopkg.in/validator.v2"
	"strings"
	"github.com/vegh1010/golang.porfolio/library/utilities"
)

var ErrUnsupported = errors.New("Error Unsupported")
var ErrBadParameter = errors.New("Error Bad Parameter")
var ErrRequired = errors.New("This field is required")
var ErrMinimum = errors.New("This field reached minimum value")
var ErrMaximum = errors.New("This field reached maximum value")
var ErrFormat = errors.New("Invalid Format")
var ErrEmail = errors.New("Invalid Email")
var ErrRestricted = errors.New("Not Allowed Field")

func NewValidator() (*validator.Validator)  {
	customerValidator := validator.NewValidator()
	//set validator function
	customerValidator.SetValidationFunc("vrequire", requiredValidation)

	customerValidator.SetValidationFunc("vmin", minValidation)
	customerValidator.SetValidationFunc("vmax", maxValidation)

	customerValidator.SetValidationFunc("vminlength", minLengthValidation)
	customerValidator.SetValidationFunc("vmaxlength", maxLengthValidation)

	customerValidator.SetValidationFunc("vminbits", minBitsValidation)
	customerValidator.SetValidationFunc("vmaxbits", maxBitsValidation)

	customerValidator.SetValidationFunc("vrestricted", restrictedValidation)
	customerValidator.SetValidationFunc("vemail", emailValidation)
	customerValidator.SetValidationFunc("vregex", regexValidation)

	return customerValidator
}

// Validate validates struct attribute and generate error messages for each attribute
//  Params:
//  @data Struct interface{}
//  Returns bool and interface
func Validate(dataStruct interface{}) (bool, interface{}) {
	customerValidator := NewValidator()

	if vErr := customerValidator.Validate(dataStruct); vErr != nil {
		result := make(map[string]interface{})
		//generate json name
		resultMap, err := utilities.GetStructTags(dataStruct, "json", true)
		if err != nil {
			result["failed"] = err.Error()
			return false, result
		}
		for key, value := range resultMap {
			if value == "" {
				resultMap[key] = key
			}
		}
		errs, _ := vErr.(validator.ErrorMap)
		for key, value := range errs {
			for fieldKey, fieldValue := range resultMap {
				if strings.ToLower(key) == strings.ToLower(fieldKey) {
					result[fieldValue] = value.Error()
					break
				}
			}
		}
		return false, result
	}
	return true, nil
}

// Builtin validator functions
// Here is the list of validator functions builtin in the package.
//  vmax
//  For numeric numbers, max will simply make sure that the value is
//  lesser or equal to the parameter given. For strings, it checks that
//  the string length is at most that number of characters. For slices,
//  arrays, and maps, validates the number of items. (Usage: max=10)
//
//  vmin
//  For numeric numbers, min will simply make sure that the value is
//  greater or equal to the parameter given. For strings, it checks that
//  the string length is at least that number of characters. For slices,
//  arrays, and maps, validates the number of items. (Usage: min=10)
//
//  vrequire
//  This validates that the value is not zero. The appropriate zero value
//  is given by the Go spec (e.g. for int it's 0, for string it's "", for
//  pointers is nil, etc.) Usage: nonzero
//
//  vregex
//  Only valid for string types, it will validate that the value matches
//  the regular expression provided as parameter. (Usage: regexp=^a.*b$)

// required validator func
func requiredValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	valid := true
	switch st.Kind() {
	case reflect.String:
		valid = len(st.String()) != 0
	case reflect.Ptr, reflect.Interface:
		valid = !st.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		valid = st.Len() != 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = st.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		valid = st.Uint() != 0
	case reflect.Float32, reflect.Float64:
		valid = st.Float() != 0
	case reflect.Bool:
		valid = st.Bool()
	case reflect.Invalid:
		valid = false // always invalid
	case reflect.Struct:
		valid = true // always valid since only nil pointers are empty
	default:
		return ErrUnsupported
	}

	if !valid {
		return ErrRequired
	}
	return nil
}

// min length validator func
func minLengthValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	invalid := false

	p, err := asInt(param)
	if err != nil {
		return ErrBadParameter
	}
	pInt := int(p)

	switch st.Kind() {
	case reflect.String:
		invalid = len(fmt.Sprint(st.String())) < pInt
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		invalid = len(fmt.Sprint(st.Int())) < pInt
	case reflect.Ptr:
		if st.IsNil() {
			invalid = false
		} else {
			return minLengthValidation(st.Elem(), param)
		}
	default:
		return ErrUnsupported
	}
	if invalid {
		msg := fmt.Sprint("This field reached minimum length:", param)
		return errors.New(msg)
	}
	return nil
}

// max length validator func
func maxLengthValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	invalid := false

	p, err := asInt(param)
	if err != nil {
		return ErrBadParameter
	}
	pInt := int(p)

	switch st.Kind() {
	case reflect.String:
		invalid = len(fmt.Sprint(st.String())) > pInt
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		invalid = len(fmt.Sprint(st)) > pInt // evaluate the length of the text value
	case reflect.Ptr:
		if st.IsNil() {
			invalid = false
		} else {
			return maxLengthValidation(st.Elem(), param)
		}
	default:
		return ErrUnsupported
	}

	if invalid {
		msg := fmt.Sprint("This field reached maximum length:", param)
		return errors.New(msg)
	}
	return nil
}

// min length validator func
func minBitsValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	invalid := false

	var powOf float64 = 2
	p, err := asFloat(param)
	if err != nil {
		return ErrBadParameter
	}
	bitMinFloat := math.Pow(p, powOf)
	bitMin := int64(bitMinFloat)

	switch st.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		invalid = st.Int() < bitMin
	case reflect.Ptr:
		if st.IsNil() {
			invalid = false
		} else {
			return minBitsValidation(st.Elem(), param)
		}
	default:
		return ErrUnsupported
	}
	if invalid {
		msg := fmt.Sprint("This field reached minimum bits:", param)
		return errors.New(msg)
	}
	return nil
}

// max length validator func
func maxBitsValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	invalid := false

	var powOf float64 = 2
	p, err := asFloat(param)
	if err != nil {
		return ErrBadParameter
	}
	bitMaxFloat := math.Pow(p, powOf)
	bitMax := int64(bitMaxFloat)

	switch st.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		invalid = st.Int() > bitMax
	case reflect.Ptr:
		if st.IsNil() {
			invalid = false
		} else {
			return maxBitsValidation(st.Elem(), param)
		}
	default:
		return ErrUnsupported
	}

	if invalid {
		msg := fmt.Sprint("This field reached maximum bits:", param)
		return errors.New(msg)
	}
	return nil
}

// min validator func
func minValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	invalid := false
	switch st.Kind() {
	case reflect.String:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = int64(len(st.String())) < p
	case reflect.Slice, reflect.Map, reflect.Array:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = int64(st.Len()) < p
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = st.Int() < p
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p, err := asUint(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = st.Uint() < p
	case reflect.Float32, reflect.Float64:
		p, err := asFloat(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = st.Float() < p
	case reflect.Ptr:
		if st.IsNil() {
			invalid = false
		} else {
			return minValidation(st.Elem(), param)
		}
	default:
		return ErrUnsupported
	}
	if invalid {
		return ErrMinimum
	}
	return nil
}

// max validator func
func maxValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	var invalid bool
	switch st.Kind() {
	case reflect.String:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = int64(len(st.String())) > p
	case reflect.Slice, reflect.Map, reflect.Array:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = int64(st.Len()) > p
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = st.Int() > p
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p, err := asUint(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = st.Uint() > p
	case reflect.Float32, reflect.Float64:
		p, err := asFloat(param)
		if err != nil {
			return ErrBadParameter
		}
		invalid = st.Float() > p
	case reflect.Ptr:
		if st.IsNil() {
			invalid = false
		} else {
			return maxValidation(st.Elem(), param)
		}
	default:
		return ErrUnsupported
	}
	if invalid {
		return ErrMaximum
	}
	return nil
}

// regex validator func
func regexValidation(v interface{}, param string) error {
	re, err := regexp.Compile(param)
	if err != nil {
		return ErrBadParameter
	}
	st := reflect.ValueOf(v)
	switch st.Kind() {
	case reflect.String:
		if !re.MatchString(st.String()) {
			return ErrFormat
		}
	case reflect.Ptr:
		if st.IsNil() {
			return nil
		} else {
			return regexValidation(st.Elem(), param)
		}
	default:
		return ErrUnsupported
	}
	return nil
}

// email validator func
func emailValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() == reflect.Ptr {
		if st.IsNil() {
			return nil
		} else {
			return emailValidation(st.Elem(), param)
		}
	}
	if st.Kind() != reflect.String {
		return ErrUnsupported
	}
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	exp, err := regexp.Compile(regex)
	if err != nil {
		return ErrBadParameter
	}
	if !exp.MatchString(st.String()) {
		return ErrEmail
	}
	return nil
}

// email validator func
func restrictedValidation(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() == reflect.Ptr {
		if st.IsNil() {
			return ErrUnsupported
		} else {
			return restrictedValidation(st.Elem(), param)
		}
	}
	if st.Kind() != reflect.String {
		return ErrUnsupported
	}
	valid := len(st.String()) == 0
	if !valid {
		return ErrRestricted
	}
	return nil
}

// asInt returns the parameter as a int64
// or panics if it can't convert
func asInt(param string) (int64, error) {
	i, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		return 0, ErrBadParameter
	}
	return i, nil
}

// asUint returns the parameter as a uint64
// or panics if it can't convert
func asUint(param string) (uint64, error) {
	i, err := strconv.ParseUint(param, 0, 64)
	if err != nil {
		return 0, ErrBadParameter
	}
	return i, nil
}

// asFloat returns the parameter as a float64
// or panics if it can't convert
func asFloat(param string) (float64, error) {
	i, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0.0, ErrBadParameter
	}
	return i, nil
}
