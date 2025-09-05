package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func processErr(err error, emp Employee) string {
	if err == nil {
		return ""
	}

	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("ID(%s) is invalid", emp.ID)
	}

	var errEmptyField ErrEmptyField
	if errors.As(err, &errEmptyField) {
		return fmt.Sprintf("%s field is empty", errEmptyField.FieldName)
	}

	return fmt.Sprintf("%v", err)
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}

		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d: %+v", count, emp)
		if err != nil {
			switch err := err.(type) {
			case interface{ Unwrap() []error }:
				var messages []string
				for _, e := range err.Unwrap() {
					messages = append(messages, processErr(e, emp))
				}
				message = fmt.Sprintf("%s errors: %s", message, strings.Join(messages, ", "))
			default:
				message = fmt.Sprintf("%s error: %v", message, processErr(err, emp))
			}
		}

		fmt.Println(message)
	}
}
