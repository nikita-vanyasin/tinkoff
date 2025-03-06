package tinkoff

import (
	"context"
	"fmt"
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
	Data            map[string]string `json:"DATA,omitempty"`            // Дополнительные параметры платежа
	Receipt         *Receipt          `json:"Receipt,omitempty"`         // Чек
	RedirectDueDate Time              `json:"RedirectDueDate,omitempty"` // Срок жизни ссылки
	NotificationURL string            `json:"NotificationURL,omitempty"` // Адрес для получения http нотификаций
	SuccessURL      string            `json:"SuccessURL,omitempty"`      // Страница успеха
	FailURL         string            `json:"FailURL,omitempty"`         // Страница ошибки
	PayType         string            `json:"PayType,omitempty"`         // Тип оплаты. см. PayType*
	Shops           *[]Shop           `json:"Shops,omitempty"`           // Объект с данными партнера
	Descriptor      string            `json:"Descriptor,omitempty"`      // Динамический дескриптор точки
}

func (i *InitRequest) SetIsRecurrent(r bool) {
	if r {
		i.Recurrent = "Y"
	} else {
		i.Recurrent = ""
	}
}

func (i *InitRequest) GetValuesForToken() map[string]string {
	v := map[string]string{
		"OrderId":         i.OrderID,
		"IP":              i.ClientIP,
		"Description":     i.Description,
		"Language":        i.Language,
		"CustomerKey":     i.CustomerKey,
		"RedirectDueDate": i.RedirectDueDate.String(),
		"NotificationURL": i.NotificationURL,
		"SuccessURL":      i.SuccessURL,
		"FailURL":         i.FailURL,
		"Recurrent":       i.Recurrent,
	}
	serializeUintToMapIfNonEmpty(&v, "Amount", i.Amount)
	return v
}

type InitResponse struct {
	BaseResponse
	Amount     uint64 `json:"Amount"`               // Сумма в копейках
	OrderID    string `json:"OrderId"`              // Номер заказа в системе Продавца
	Status     string `json:"Status"`               // Статус транзакции
	PaymentID  string `json:"PaymentId"`            // Уникальный идентификатор транзакции в системе Банка. По офф. документации это number(20), но фактически значение передается в виде строки.
	PaymentURL string `json:"PaymentURL,omitempty"` // Ссылка на страницу оплаты. По умолчанию ссылка доступна в течении 24 часов.
}

// Init prepares new payment transaction
// Deprecated: use InitWithContext instead
func (c *Client) Init(request *InitRequest) (*InitResponse, error) {
	return c.InitWithContext(context.Background(), request)
}

// InitWithContext prepares new payment transaction
func (c *Client) InitWithContext(ctx context.Context, request *InitRequest) (*InitResponse, error) {
	response, err := c.PostRequestWithContext(ctx, "/Init", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res InitResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	err = res.Error()
	if res.Status != StatusNew {
		err = errorConcat(err, fmt.Errorf("unexpected payment status: %s", res.Status))
	}

	return &res, err
}
