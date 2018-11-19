package tinkoff

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
)

type Notification struct {
	TerminalKey    string            `json:"TerminalKey"` // Идентификатор магазина
	OrderID        string            `json:"OrderId"`     // Номер заказа в системе Продавца
	Success        bool              `json:"Success"`     // Успешность операции
	Status         string            `json:"Status"`      // Статус платежа (см. описание статусов операций)
	PaymentID      uint64            `json:"PaymentId"`   // Уникальный идентификатор платежа
	ErrorCode      string            `json:"ErrorCode"`   // Код ошибки, если произошла ошибка
	Amount         uint64            `json:"Amount"`      // Текущая сумма транзакции в копейках
	RebillID       string            `json:"RebillId"`    // Идентификатор рекуррентного платежа
	CardID         uint64            `json:"CardId"`      // Идентификатор привязанной карты
	PAN            string            `json:"Pan"`         // Маскированный номер карты
	DataStr        string            `json:"DATA"`
	Data           map[string]string `json:"-"`       // Дополнительные параметры платежа, переданные при создании заказа
	Token          string            `json:"Token"`   // Подпись запроса
	ExpirationDate string            `json:"ExpDate"` // Срок действия карты
}

func (n *Notification) GetValuesForToken() map[string]string {
	var result = map[string]string{
		"TerminalKey": n.TerminalKey,
		"OrderId":     n.OrderID,
		"Success":     serializeBool(n.Success),
		"Status":      n.Status,
		"PaymentId":   strconv.FormatUint(n.PaymentID, 10),
		"ErrorCode":   n.ErrorCode,
		"Amount":      strconv.FormatUint(n.Amount, 10),
		"CardId":      strconv.FormatUint(n.CardID, 10),
		"Pan":         n.PAN,
		"ExpDate":     n.ExpirationDate,
	}

	if n.DataStr != "" {
		result["DATA"] = n.DataStr
	}

	if n.RebillID != "" {
		result["RebillId"] = n.RebillID
	}

	return result
}

func (c *Client) ParseNotification(requestBody io.Reader) (*Notification, error) {
	bytes, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}
	var notification Notification
	err = json.Unmarshal(bytes, &notification)
	if err != nil {
		return nil, err
	}

	if c.terminalKey != notification.TerminalKey {
		return nil, errors.New("invalid terminal key")
	}

	valuesForTokenGen := notification.GetValuesForToken()
	valuesForTokenGen["Password"] = c.password
	token := generateToken(valuesForTokenGen)
	if token != notification.Token {
		return nil, errors.New("invalid token")
	}

	if notification.DataStr != "" {
		err = json.Unmarshal([]byte(notification.DataStr), &notification.Data)
		if err != nil {
			return nil, errors.New("can't unserialize DATA field: " + err.Error())
		}
	}

	return &notification, nil
}

func (c *Client) GetNotificationSuccessResponse() string {
	return "OK"
}
