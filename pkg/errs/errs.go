package errs

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

// Error is the type that implements the error interface.
// It contains a number of fields, each of different type.
// An Error value may leave some values unset.
type Error struct {
	// or "Other" if its class is unknown or irrelevant.
	Kind Kind
	// Code is a human-readable, short representation of the error
	Code Code
	// The underlying error that triggered this one, if any.
	Err error
}

// Error func to implement error interface
func (e *Error) Error() string {
	return e.Err.Error()
}

// Returns original error (if any)
func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) isZero() bool {
	return e.Kind == 0 && e.Code == "" && e.Err == nil
}

// Kind defines the kind of error this is, mostly for use by systems
type Kind uint8

// Code is a human-readable, short representation of the error
type Code string

// HTTP status code as registered with IANA.
type Status uint8

const (
	Other          Kind = iota // Unclassified error. This value is not printed in the error message.
	Invalid                    // Invalid operation for this type of item.
	IO                         // External I/O error such as network failure.
	Exist                      // Item already exists.
	NotExist                   // Item does not exist.
	Private                    // Information withheld.
	Internal                   // Internal error or inconsistency.
	BrokenLink                 // Link target does not exist.
	Database                   // Error from database.
	Validation                 // Input validation error.
	Unanticipated              // Unanticipated error.
	InvalidRequest             // Invalid Request
	// Unauthenticated is used when a request lacks valid authentication credentials.
	//
	// For Unauthenticated errors, the response body will be empty.
	// The error is logged and http.StatusUnauthorized (401) is sent.
	Unauthenticated // Unauthenticated Request
	// Unauthorized is used when a user is authenticated, but is not authorized
	// to access the resource.
	//
	// For Unauthorized errors, the response body should be empty.
	// The error is logged and http.StatusForbidden (403) is sent.
	Unauthorized
)

func (k Kind) String() string {
	switch k {
	case Other:
		return "other_error"
	case Invalid:
		return "invalid_operation"
	case IO:
		return "I/O_error"
	case Exist:
		return "item_already_exists"
	case NotExist:
		return "item_does_not_exist"
	case BrokenLink:
		return "link_target_does_not_exist"
	case Private:
		return "information_withheld"
	case Internal:
		return "internal_error"
	case Database:
		return "database_error"
	case Validation:
		return "input_validation_error"
	case Unanticipated:
		return "unanticipated_error"
	case InvalidRequest:
		return "invalid_request_error"
	case Unauthenticated:
		return "unauthenticated_request"
	case Unauthorized:
		return "unauthorized_request"
	}

	return "unknown_error_kind"
}

// E builds an error value from its arguments.
// There must be at least one argument or E panics.
// The type of each argument determines its meaning.
// If more than one argument of a given type is presented,
// only the last one is recorded.
//
// The types are:
//
//		string
//			Treated as an error message and assigned to the
//			Err field after a call to errs.New.
//		errs.Kind
//			The class of error, such as permission failure.
//	 	errs.Code
//			The readable code
//		error
//			The underlying error that triggered this one.
//
// If the error is printed, only those items that have been
// set to non-zero values will appear in the result.
//
// If Kind is not specified or Other, we set it to the Kind of
// the underlying error.
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errs.E with no arguments")
	}

	e := &Error{}

	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			e.Err = errors.New(arg)
		case error:
			e.Err = arg
		case Kind:
			e.Kind = arg
		case Code:
			e.Code = arg
		default:
			_, file, line, _ := runtime.Caller(1)
			return fmt.Errorf("errs.E: bad call from %s:%d: %v, unknown type %T, value %v in error call", file, line, args, arg, arg)
		}
	}

	return e
}

// Match compares its two error arguments. It can be used to check
// for expected errors in tests. Both arguments must have underlying
// type *Error or Match will return false. Otherwise it returns true
// if every non-zero element of the first error is equal to the
// corresponding element of the second.
// If the Err field is a *Error, Match recurs on that field;
// otherwise it compares the strings returned by the Error methods.
// Elements that are in the second argument but not present in
// the first are ignored.
func Match(err1, err2 error) bool {
	e1, ok := err1.(*Error)
	if !ok {
		return false
	}

	e2, ok := err2.(*Error)
	if !ok {
		return false
	}

	if e1.Kind != Other && e2.Kind != e1.Kind {
		return false
	}

	if e1.Code != "" && e2.Code != e1.Code {
		return false
	}

	if e1.Err != nil {
		if _, ok := e1.Err.(*Error); ok {
			return Match(e1.Err, e2.Err)
		}

		if e2.Err == nil || e2.Err.Error() != e1.Err.Error() {
			return false
		}
	}

	return true
}

// KindIs reports whether err is an *Error of the given Kind.
// If err is nil then KindIs returns false.
func KindIs(kind Kind, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}

	if e.Kind != Other {
		return e.Kind == kind
	}

	if e.Err != nil {
		return KindIs(kind, e.Err)
	}

	return false
}

// httpErrorStatusCode maps an error Kind to an HTTP Status Code
func httpErrorStatusCode(k Kind) int {
	switch k {
	case Invalid, NotExist, Private, BrokenLink, Validation, InvalidRequest:
		return http.StatusBadRequest
	case Exist:
		return http.StatusConflict
	case Unauthenticated:
		return http.StatusUnauthorized
	case Unauthorized:
		return http.StatusForbidden
	// the zero value of Kind is Other, so if no Kind is present
	// in the error, Other is used. Errors should always have a
	// Kind set, otherwise, a 500 will be returned and no
	// error message will be sent to the caller
	case Other, IO, Internal, Database, Unanticipated:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
