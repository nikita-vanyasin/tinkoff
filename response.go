package tinkoff

import "fmt"

type BaseResponse struct {
	TerminalKey  string `json:"TerminalKey"`       // Идентификатор терминала, выдается Продавцу Банком
	Success      bool   `json:"Success"`           // Успешность операции
	ErrorCode    string `json:"ErrorCode"`         // Код ошибки, «0» - если успешно
	ErrorMessage string `json:"Message,omitempty"` // Краткое описание ошибки
	ErrorDetails string `json:"Details,omitempty"` // Подробное описание ошибки
}

func (e *BaseResponse) Error() error {
	if !e.Success || e.ErrorCode != "0" {
		return fmt.Errorf("error code %s - %s. %s", e.ErrorCode, e.ErrorMessage, e.ErrorDetails)
	}
	return nil
}
