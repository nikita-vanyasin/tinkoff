package tinkoff

import (
	"fmt"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t Time) String() string {
	original := time.Time(t)
	if original.IsZero() {
		return ""
	}
	return original.Format(time.RFC3339)
}

type ErrorInfo struct {
	ErrorCode    string `json:"ErrorCode"`         // Код ошибки, «0» - если успешно
	ErrorMessage string `json:"Message,omitempty"` // Краткое описание ошибки
	ErrorDetails string `json:"Details,omitempty"` // Подробное описание ошибки
}

func (e *ErrorInfo) FormatErrorInfo() string {
	return fmt.Sprintf("error code %s - %s. %s", e.ErrorCode, e.ErrorMessage, e.ErrorDetails)
}

func serializeBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
