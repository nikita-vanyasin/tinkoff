package tinkoff

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	RedirectDueDateFormat = time.RFC3339
)

type InitRequest struct {
	BaseRequest

	Amount          uint64            `json:"Amount,omitempty"`          // Сумма в копейках
	OrderID         string            `json:"OrderId"`                   // Идентификатор заказа в системе продавца
	ClientIP        string            `json:"IP,omitempty"`              // IP-адрес покупателя
	Description     string            `json:"Description,omitempty"`     // Описание заказа
	Language        string            `json:"Language,omitempty"`        // Язык платежной формы: ru или en
	CustomerKey     string            `json:"CustomerKey,omitempty"`     // Идентификатор покупателя в системе продавца. Передается вместе с параметром CardId. См. метод GetCardList
	Data            map[string]string `json:"DATA"`                      // Дополнительные параметры платежа
	Receipt         *Receipt          `json:"Receipt,omitempty"`         // Чек
	RedirectDueDate string            `json:"RedirectDueDate,omitempty"` // Срок жизни ссылки
	NotificationURL string            `json:"NotificationURL,omitempty"` // Адрес для получения http нотификаций
	SuccessURL      string            `json:"SuccessURL,omitempty"`      // Страница успеха
	FailURL         string            `json:"FailURL,omitempty"`         // Страница ошибки

	// Not implemented yet:
	// Recurrent
	// PayType
}

func (i *InitRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"Amount":          strconv.FormatUint(i.Amount, 10),
		"OrderId":         i.OrderID,
		"IP":              i.ClientIP,
		"Description":     i.Description,
		"Language":        i.Language,
		"CustomerKey":     i.CustomerKey,
		"RedirectDueDate": i.RedirectDueDate,
		"NotificationURL": i.NotificationURL,
		"SuccessURL":      i.SuccessURL,
		"FailURL":         i.FailURL,
	}
}

type InitResponse struct {
	TerminalKey  string `json:"TerminalKey"`          // Идентификатор терминала, выдается Продавцу Банком
	Amount       uint64 `json:"Amount"`               // Сумма в копейках
	OrderID      string `json:"OrderId"`              // Номер заказа в системе Продавца
	Success      bool   `json:"Success"`              // Успешность операции
	Status       string `json:"Status"`               // Статус транзакции
	PaymentID    string `json:"PaymentId"`            // Уникальный идентификатор транзакции в системе Банка. По офф. документации это number(20), но фактически значение передается в виде строки.
	PaymentURL   string `json:"PaymentURL,omitempty"` // Ссылка на страницу оплаты. По умолчанию ссылка доступна в течении 24 часов.
	ErrorCode    string `json:"ErrorCode"`            // Код ошибки, «0» - если успешно
	ErrorMessage string `json:"Message,omitempty"`    // Краткое описание ошибки
	ErrorDetails string `json:"Details,omitempty"`    // Подробное описание ошибки
}

func validateDateFormat(dateStr string) error {
	if dateStr == "" {
		return nil
	}
	_, err := time.Parse(RedirectDueDateFormat, dateStr)
	return err
}

func (c *Client) Init(request *InitRequest) (*InitResponse, error) {
	err := validateDateFormat(request.RedirectDueDate)
	if err != nil {
		err = errors.New(fmt.Sprintf("while RedirectDueDate validation: %v", err))
		return nil, err
	}

	response, err := c.postRequest("/Init", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res InitResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	if !res.Success || res.ErrorCode != "0" || res.Status == StatusRejected {
		errMsg := fmt.Sprintf(
			"while init request: code %s - %s. %s",
			res.ErrorCode,
			res.ErrorMessage,
			res.ErrorDetails,
		)
		err = errors.New(errMsg)
	}

	return &res, err
}
