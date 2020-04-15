package tinkoff

import (
	"errors"
	"fmt"
	"strconv"
)

type CancelRequest struct {
	BaseRequest

	PaymentID uint64   `json:"PaymentId"`
	ClientIP  string   `json:"IP,omitempty"`
	Amount    uint64   `json:"Amount,omitempty"`
	Receipt   *Receipt `json:"Receipt,omitempty"`
}

func (i *CancelRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"Amount":    strconv.FormatUint(i.Amount, 10),
		"IP":        i.ClientIP,
		"PaymentId": strconv.FormatUint(i.PaymentID, 10),
	}
}

type CancelResponse struct {
	TerminalKey    string `json:"TerminalKey"`       // Идентификатор терминала, выдается Продавцу Банком
	OriginalAmount uint64 `json:"OriginalAmount"`    // Сумма в копейках до операции отмены
	NewAmount      uint64 `json:"NewAmount"`         // Сумма в копейках после операции отмены
	OrderID        string `json:"OrderId"`           // Номер заказа в системе Продавца
	Success        bool   `json:"Success"`           // Успешность операции
	Status         string `json:"Status"`            // Статус транзакции
	PaymentID      string `json:"PaymentId"`         // Уникальный идентификатор транзакции в системе Банка
	ErrorCode      string `json:"ErrorCode"`         // Код ошибки, «0» - если успешно
	ErrorMessage   string `json:"Message,omitempty"` // Краткое описание ошибки
	ErrorDetails   string `json:"Details,omitempty"` // Подробное описание ошибки
}

func (c *Client) Cancel(request *CancelRequest) (status string, originalAmount uint64, newAmount uint64, err error) {
	response, err := c.postRequest("/Cancel", request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var res CancelResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return
	}

	status = res.Status
	originalAmount = res.OriginalAmount
	newAmount = res.NewAmount

	if !res.Success || res.ErrorCode != "0" {
		err = errors.New(fmt.Sprintf("while Cancel request: code %s - %s. %s", res.ErrorCode, res.ErrorMessage, res.ErrorDetails))
	}

	return
}
