package keysrequest

import (
	"fmt"
	"strings"
)

// ErrInvalid is an error that happens when an invalid
// keysrequest is build
type ErrInvalid struct {
	msg   string
	field string
	sub   *ErrInvalid
}

func (e ErrInvalid) Error() string {
	fields, last := e.fields()
	if fields == nil {
		return e.msg
	}
	return fmt.Sprintf("field \"%s\": %s", strings.Join(fields, "."), last.msg)
}

// Fields returns a list of field names from the parent to this error
func (e ErrInvalid) Fields() []string {
	fields, _ := e.fields()
	return fields
}

func (e ErrInvalid) fields() ([]string, *ErrInvalid) {
	if e.field == "" {
		return nil, nil
	}
	if e.sub == nil {
		return []string{e.field}, &e
	}
	var fields []string
	last := &e
	for last.sub != nil {
		fields = append(fields, last.field)
		last = last.sub
	}
	return fields, last
}

// ErrJSON is returned when invalid json is parsed or the json can not
// be decoded to the KeysRequest
type ErrJSON struct {
	msg string
	err error
}

func (e ErrJSON) Error() string {
	return e.msg + ": " + e.err.Error()
}

// Unwrap returns the thrown error
func (e ErrJSON) Unwrap() error {
	return e.err
}