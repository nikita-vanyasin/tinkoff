package tinkoff

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	PayTypeOneStep  = "O"
	PayTypeTwoSteps = "T"
)

type InitRequest struct {
	BaseRequest

	Amount          uint64            `json:"Amount,omitempty"`          // Сумма в копейках
	OrderID         string            `json:"OrderId"`                   // Идентификатор заказа в системе продавца
	ClientIP        string            `json:"IP,omitempty"`              // IP-адрес покупателя
	Description     string            `json:"Description,omitempty"`     // Описание заказа
	Language        string            `json:"Language,omitempty"`        // Язык платежной формы: ru или en
	Recurrent       string            `json:"Recurrent,omitempty"`       // Y для регистрации автоплатежа. Можно использовать SetIsRecurrent(true)
	CustomerKey     string            `json:"CustomerKey,omitempty"`     // Идентификатор покупателя в системе продавца. Передается вместе с параметром CardId. См. метод GetCardList
	Data            map[string]string `json:"DATA"`                      // Дополнительные параметры платежа
	Receipt         *Receipt          `json:"Receipt,omitempty"`         // Чек
	RedirectDueDate Time              `json:"RedirectDueDate,omitempty"` // Срок жизни ссылки
	NotificationURL string            `json:"NotificationURL,omitempty"` // Адрес для получения http нотификаций
	SuccessURL      string            `json:"SuccessURL,omitempty"`      // Страница успеха
	FailURL         string            `json:"FailURL,omitempty"`         // Страница ошибки
	PayType         string            `json:"PayType,omitempty"`         // Тип оплаты. см. PayType*
}

func (i *InitRequest) SetIsRecurrent(r bool) {
	if r {
		i.Recurrent = "Y"
	} else {
		i.Recurrent = ""
	}
}

func (i *InitRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"Amount":          strconv.FormatUint(i.Amount, 10),
		"OrderId":         i.OrderID,
		"IP":              i.ClientIP,
		"Description":     i.Description,
		"Language":        i.Language,
		"CustomerKey":     i.CustomerKey,
		"RedirectDueDate": i.RedirectDueDate.String(),
		"NotificationURL": i.NotificationURL,
		"SuccessURL":      i.SuccessURL,
		"FailURL":         i.FailURL,
	}
}

type InitResponse struct {
	TerminalKey string `json:"TerminalKey"`          // Идентификатор терминала, выдается Продавцу Банком
	Amount      uint64 `json:"Amount"`               // Сумма в копейках
	OrderID     string `json:"OrderId"`              // Номер заказа в системе Продавца
	Success     bool   `json:"Success"`              // Успешность операции
	Status      string `json:"Status"`               // Статус транзакции
	PaymentID   string `json:"PaymentId"`            // Уникальный идентификатор транзакции в системе Банка. По офф. документации это number(20), но фактически значение передается в виде строки.
	PaymentURL  string `json:"PaymentURL,omitempty"` // Ссылка на страницу оплаты. По умолчанию ссылка доступна в течении 24 часов.
	ErrorInfo
}

func (c *Client) Init(request *InitRequest) (*InitResponse, error) {
	response, err := c.PostRequest("/Init", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res InitResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	errMsg := ""
	if !res.Success || res.ErrorCode != "0" {
		errMsg = res.FormatErrorInfo()
	}
	if res.Status != StatusNew {
		errMsg = fmt.Sprintf("unexpected payment status: %s. %s", res.Status, errMsg)
	}
	if errMsg != "" {
		err = errors.New(errMsg)
	}
	return &res, err
}
