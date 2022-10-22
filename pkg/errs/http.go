package errs

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

// ServiceError has fields for Service errors. All fields with no data will
// be omitted

type ErrResponse struct {
	Error ServiceError `json:"error"`
}

type ServiceError struct {
	Kind    string `json:"kind,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func HTTPErrorResponse(w http.ResponseWriter, log *zap.Logger, err error) {
	var e *Error
	if errors.As(err, &e) {
		httpStatusCode := httpErrorStatusCode(e.Kind)

		// zero error (just in case)
		if e.isZero() {
			log.Error("Response error sent",
				zap.Int("http_statuscode", httpStatusCode),
			)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		// typical error

		// log the error
		log.Error("Response error sent",
			zap.Int("http_statuscode", httpStatusCode),
			zap.String("Kind", e.Kind.String()),
			zap.String("Code", string(e.Code)),
		)

		// response error
		er := ErrResponse{
			Error: ServiceError{
				Kind:    e.Kind.String(),
				Code:    string(e.Code),
				Message: e.Err.Error(),
			},
		}

		// Marshal errResponse struct to JSON for the response body
		errJSON, _ := json.Marshal(er)
		ej := string(errJSON)

		// Write Content-Type headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Write HTTP Statuscode
		w.WriteHeader(httpStatusCode)

		// Write response body (json)
		_, err = w.Write([]byte(ej))
		if err != nil {
			log.Error("Unable to send response error",
				zap.Int("http_statuscode", httpStatusCode),
				zap.String("Kind", e.Kind.String()),
				zap.String("Code", string(e.Code)),
			)
		}

		return
	}

	er := ErrResponse{
		Error: ServiceError{
			Kind:    Unanticipated.String(),
			Code:    "Unanticipated",
			Message: "Unexpected error - contact support",
		},
	}

	log.Error("Unknown response error", zap.Error(err))

	// Marshal errResponse struct to JSON for the response body
	errJSON, _ := json.Marshal(er)
	ej := string(errJSON)

	// Write Content-Type headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Write HTTP Statuscode
	w.WriteHeader(http.StatusInternalServerError)

	// Write response body (json)
	_, err = w.Write([]byte(ej))
	if err != nil {
		log.Error("Unable to send response error",
			zap.Int("http_statuscode", http.StatusInternalServerError),
			zap.String("Kind", e.Kind.String()),
			zap.String("Code", string(e.Code)),
		)
	}
}
