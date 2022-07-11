package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	validationErrorStrings := make([]string, 0)
	for _, err := range v {
		validationErrorStrings = append(
			validationErrorStrings,
			fmt.Sprintf("Validation error: %v, in property: %v", err.Err, err.Field),
		)
	}

	return strings.Join(validationErrorStrings, "\n")
}

var (
	minValueValidationRegexp = regexp.MustCompile(`min:(\d+)`)
	maxValueValidationRegexp = regexp.MustCompile(`max:(\d+)`)
)

var (
	ErrIntOutOfList    = errors.New("int out of list")
	ErrIntTooSmall     = errors.New("too small")
	ErrIntTooLarge     = errors.New("too big")
	ErrStringOutOfList = errors.New("string out of list")
	ErrStringLen       = errors.New("length is invalid")
	ErrStringRegexp    = errors.New("does not match the regexp pattern")
)

var (
	ErrUnsupportedPropertyType = errors.New("unsupported property type")
	ErrUnsupportedType         = errors.New("type must be struct")
)

func Validate(v interface{}) error {
	reflectV := reflect.ValueOf(v)

	if reflectV.Kind() != reflect.Struct {
		return ErrUnsupportedType
	}

	typeV := reflectV.Type()

	validationErrors := make(ValidationErrors, 0)
	for i := 0; i < typeV.NumField(); i++ {
		property := typeV.Field(i)
		propertyValidationRules := property.Tag.Get("validate")
		if len(propertyValidationRules) == 0 {
			continue
		}
		rules := strings.Split(propertyValidationRules, "|")
		validationErrorsForProperty, err := validateProperty(property, reflectV, i, rules)
		if err != nil {
			return err
		}

		validationErrors = append(validationErrors, validationErrorsForProperty...)
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func validateProperty(
	property reflect.StructField,
	reflectedStruct reflect.Value,
	i int,
	rules []string,
) (ValidationErrors, error) {
	var validationErrors ValidationErrors
	var err error

	// Без nolint получаю ошибку аля "не обработаны кейсы для: Array, Bool, Chan etc". Другое решение в первом коммите
	switch property.Type.Kind() { //nolint
	case reflect.Slice:
		propertyVal := reflectedStruct.Field(i)
		validationErrors, err = validateSlice(property.Name, propertyVal, rules)
	case reflect.Int:
		propertyVal := int(reflectedStruct.Field(i).Int())
		validationErrors, err = validateInt(property.Name, propertyVal, rules)
	case reflect.String:
		propertyVal := reflectedStruct.Field(i).String()
		validationErrors, err = validateString(property.Name, propertyVal, rules)
	default:
		err = ErrUnsupportedPropertyType
	}
	return validationErrors, err
}

func validateSlice(fieldName string, values reflect.Value, rules []string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)

	switch values.Interface().(type) {
	case []string:
		for _, val := range values.Interface().([]string) {
			validationErrorsForSliceItem, err := validateString(fieldName, val, rules)
			validationErrors = append(validationErrors, validationErrorsForSliceItem...)
			if err != nil {
				return validationErrors, err
			}
		}
	case []int:
		for _, val := range values.Interface().([]int) {
			validationErrorsForSliceItem, err := validateInt(fieldName, val, rules)
			validationErrors = append(validationErrors, validationErrorsForSliceItem...)
			if err != nil {
				return validationErrors, err
			}
		}
	}

	return validationErrors, nil
}

func validateInt(fieldName string, value int, rules []string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)
	for _, rule := range rules {
		validationErrorsIntIn, err := validateIntIn(fieldName, value, rule)
		validationErrors = append(validationErrors, validationErrorsIntIn...)
		if err != nil {
			return validationErrors, err
		}

		validationErrorsIntMin, err := validateIntMin(fieldName, value, rule)
		validationErrors = append(validationErrors, validationErrorsIntMin...)
		if err != nil {
			return validationErrors, err
		}

		validationErrorsInMax, err := validateIntMax(fieldName, value, rule)
		validationErrors = append(validationErrors, validationErrorsInMax...)
		if err != nil {
			return validationErrors, err
		}
	}

	return validationErrors, nil
}

func validateIntMax(fieldName string, value int, rule string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)
	isMaxValidation := maxValueValidationRegexp.Match([]byte(rule))
	if !isMaxValidation {
		return validationErrors, nil
	}

	pattern := regexp.MustCompile(`\\d+`)
	maxValue := pattern.FindAllString(rule, 1)
	if len(maxValue) == 1 {
		minLength, err := strconv.Atoi(maxValue[0])
		if err != nil {
			return validationErrors, err
		}
		if value > minLength {
			validationError := ValidationError{Field: fieldName, Err: ErrIntTooLarge}
			validationErrors = append(validationErrors, validationError)
		}
	}

	return validationErrors, nil
}

func validateIntMin(fieldName string, value int, rule string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)
	isMinValidation := minValueValidationRegexp.Match([]byte(rule))
	if !isMinValidation {
		return validationErrors, nil
	}

	pattern := regexp.MustCompile(`\\d+`)
	minValue := pattern.FindAllString(rule, 1)
	if len(minValue) == 1 {
		minLength, err := strconv.Atoi(minValue[0])
		if err != nil {
			return validationErrors, err
		}
		if value < minLength {
			validationError := ValidationError{Field: fieldName, Err: ErrIntTooSmall}
			validationErrors = append(validationErrors, validationError)
		}
	}

	return validationErrors, nil
}

func validateIntIn(fieldName string, value int, rule string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)
	isListValidation := strings.HasPrefix(rule, "in:")
	if !isListValidation {
		return validationErrors, nil
	}

	isValid := false
	rule = strings.TrimPrefix(rule, "in:")
	values := strings.Split(rule, ",")
	for _, v := range values {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			return validationErrors, err
		}
		if value == intValue {
			isValid = true
			break
		}
	}
	if !isValid {
		validationError := ValidationError{Field: fieldName, Err: ErrIntOutOfList}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors, nil
}

func validateString(fieldName string, value string, rules []string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)
	for _, rule := range rules {
		stringRegExpErrors, err := validateStringRegExp(fieldName, value, rule)
		validationErrors = append(validationErrors, stringRegExpErrors...)
		if err != nil {
			return validationErrors, err
		}

		stringInErrors := validateStringIn(fieldName, value, rule)
		validationErrors = append(validationErrors, stringInErrors...)

		stringLengthErrors, err := validateStringLength(fieldName, value, rule)
		validationErrors = append(validationErrors, stringLengthErrors...)
		if err != nil {
			return validationErrors, err
		}
	}

	return validationErrors, nil
}

func validateStringRegExp(fieldName string, value string, rule string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)
	isRegexpValidation := strings.HasPrefix(rule, "regexp:")
	if !isRegexpValidation {
		return validationErrors, nil
	}

	validationPattern := strings.TrimPrefix(rule, "regexp:")
	isValid, err := regexp.Match(validationPattern, []byte(value))
	if err != nil {
		return validationErrors, err
	}
	if !isValid {
		validationError := ValidationError{Field: fieldName, Err: ErrStringRegexp}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors, nil
}

func validateStringIn(fieldName string, value string, rule string) ValidationErrors {
	validationErrors := make(ValidationErrors, 0)
	isListValidation := strings.HasPrefix(rule, "in:")
	if !isListValidation {
		return validationErrors
	}

	rule = strings.TrimPrefix(rule, "in:")
	values := strings.Split(rule, ",")
	isValid := false
	for _, v := range values {
		if value == v {
			isValid = true
			break
		}
	}
	if !isValid {
		validationError := ValidationError{Field: fieldName, Err: ErrStringOutOfList}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}

func validateStringLength(fieldName string, value string, rule string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)
	isLengthValidation, _ := regexp.Match("^len:(\\d+)", []byte(rule))
	if !isLengthValidation {
		return validationErrors, nil
	}

	fieldLength, err := strconv.Atoi(strings.TrimPrefix(rule, "len:"))
	if err != nil {
		return validationErrors, err
	}
	isValid := len(value) == fieldLength
	if !isValid {
		validationError := ValidationError{Field: fieldName, Err: ErrStringLen}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors, nil
}
