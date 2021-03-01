package utils

import (
	"errors"
	"strings"
)

func DataValidation(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("Data Input Cannot Empty")
		case 0:
			return errors.New("Data Cannot 0")
		case nil:
			return errors.New("Data Cannot be nil")
		}
	}
	return nil
}

func ValidateInputNotSymbol(data ...interface{}) error {
	for _, value := range data {
		if strings.ContainsAny(value.(string), "!@#$%^&*(){}:><\"") {
			return errors.New("Input cannot contains symbol")
		}
	}
	return nil
}

func ValidateInputLenCharacter(val int, data ...interface{}) error {
	for _, value := range data {
		if len(value.(string)) >= val {
			return errors.New("Input cannot more sesuai kebutuhan")
		}
	}
	return nil
}
