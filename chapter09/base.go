package main

import (
	"errors"
	"fmt"
	"regexp"
)

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID      = regexp.MustCompile(`\w{4}-\d{3}`)
	ErrInvalidID = errors.New("invalid ID")
)

type ErrEmptyField struct {
	FieldName string
}

func (efe ErrEmptyField) Error() string {
	return fmt.Sprintf("missing %s field", efe.FieldName)
}

func ValidateEmployee(e Employee) error {
	var errs []error
	if len(e.ID) == 0 {
		errs = append(errs, ErrEmptyField{FieldName: "ID"})
	}

	if !validID.MatchString(e.ID) {
		errs = append(errs, ErrInvalidID)
	}

	if len(e.FirstName) == 0 {
		errs = append(errs, ErrEmptyField{FieldName: "FirstName"})
	}

	if len(e.LastName) == 0 {
		errs = append(errs, ErrEmptyField{FieldName: "LastName"})
	}

	if len(e.Title) == 0 {
		errs = append(errs, ErrEmptyField{FieldName: "Title"})
	}

	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errors.Join(errs...)
	}
}
