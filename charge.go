package tinkoff

import (
	"context"
	"fmt"
)

type ChargeRequest struct {
	BaseRequest

	PaymentID string `json:"PaymentId,omitempty"` // Уникальный идентификатор транзакции в системе Т‑Кассы
	RebillID  string `json:"RebillId,omitempty"`  // Идентификатор рекуррентного платежа
	IP        string `json:"IP,omitempty"`        // IP-адрес клиента
	SendEmail bool   `json:"SendEmail,omitempty"` // true — если клиент хочет получать уведомления на почту.
	InfoEmail string `json:"InfoEmail,omitempty"` // Адрес почты клиента. Обязателен при передаче SendEmail
}

func (i *ChargeRequest) GetValuesForToken() map[string]string {
	v := map[string]string{
		"PaymentId": i.PaymentID,
		"RebillId":  i.RebillID,
	}
	if i.IP != "" {
		v["IP"] = i.IP
	}
	if i.SendEmail {
		v["SendEmail"] = serializeBool(i.SendEmail)
	}
	if i.InfoEmail != "" {
		v["InfoEmail"] = i.InfoEmail
	}
	return v
}

type ChargeResponse struct {
	BaseResponse
	Amount    uint64 `json:"Amount"`    // Сумма в копейках
	OrderID   string `json:"OrderId"`   // Номер заказа в системе Продавца
	Status    string `json:"Status"`    // Статус транзакции
	PaymentID string `json:"PaymentId"` // Уникальный идентификатор транзакции в системе Банка. По офф. документации это number(20), но фактически значение передается в виде строки.
}

// Charge проводит рекуррентный (повторный) платеж
func (c *Client) Charge(request *ChargeRequest) (*ChargeResponse, error) {
	return c.ChargeWithContext(context.Background(), request)
}

// ChargeWithContext проводит рекуррентный (повторный) платеж
func (c *Client) ChargeWithContext(ctx context.Context, request *ChargeRequest) (*ChargeResponse, error) {
	if request.SendEmail {
		if request.InfoEmail == "" {
			return nil, fmt.Errorf("InfoEmail is empty")
		}
	}
	response, err := c.PostRequestWithContext(ctx, "/Charge", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res ChargeResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	err = res.Error()
	if res.Status != StatusConfirmed {
		err = errorConcat(err, fmt.Errorf("unexpected payment status: %s", res.Status))
	}

	return &res, err
}
