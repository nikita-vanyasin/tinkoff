package tinkoff

import (
	"errors"
	"fmt"
	"strconv"
)

type InitRequest struct {
	BaseRequest

	Amount      uint64            `json:"Amount"`
	OrderID     string            `json:"OrderId"`
	ClientIP    string            `json:"IP"`
	Description string            `json:"Description"`
	CustomerKey string            `json:"CustomerKey"`
	Data        map[string]string `json:"DATA"`
	Receipt     *Receipt          `json:"Receipt"`

	// Not implemented:
	// Recurrent
	// RedirectDueDate
}

func (i *InitRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"Amount":      strconv.FormatUint(i.Amount, 10),
		"OrderId":     i.OrderID,
		"IP":          i.ClientIP,
		"Description": i.Description,
		"CustomerKey": i.CustomerKey,
	}
}

type InitResponse struct {
	TerminalKey  string `json:"TerminalKey"`          // Идентификатор терминала, выдается Продавцу Банком
	Amount       uint64 `json:"Amount"`               // Сумма в копейках
	OrderID      string `json:"OrderId"`              // Номер заказа в системе Продавца
	Success      bool   `json:"Success"`              // Успешность операции
	Status       string `json:"Status"`               // Статус транзакции
	PaymentID    string `json:"PaymentId"`            // Уникальный идентификатор транзакции в системе Банка
	PaymentURL   string `json:"PaymentURL,omitempty"` // Ссылка на страницу оплаты. По умолчанию ссылка доступна в течении 24 часов.
	ErrorCode    string `json:"ErrorCode"`            // Код ошибки, «0» - если успешно
	ErrorMessage string `json:"Message,omitempty"`    // Краткое описание ошибки
	ErrorDetails string `json:"Details,omitempty"`    // Подробное описание ошибки
}

func (c *Client) Init(request *InitRequest) (status string, paymentID uint64, paymentURL string, err error) {
	response, err := c.postRequest("/Init", request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var res InitResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return
	}

	status = res.Status
	paymentURL = res.PaymentURL

	paymentID, err = strconv.ParseUint(res.PaymentID, 10, 0)
	additionalErrInfo := ""
	if err != nil {
		additionalErrInfo = err.Error()
	}

	if !res.Success || res.ErrorCode != "0" || res.Status == StatusRejected {
		errMsg := fmt.Sprintf(
			"while init request: code %s - %s. %s",
			res.ErrorCode,
			res.ErrorMessage,
			res.ErrorDetails,
		)
		if additionalErrInfo != "" {
			errMsg += ". also there was error while parsing PaymentID: " + additionalErrInfo
		}
		err = errors.New(errMsg)
	}

	return
}
