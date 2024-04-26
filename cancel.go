package tinkoff

import (
	"context"
)

type CancelRequest struct {
	BaseRequest

	PaymentID string   `json:"PaymentId"`         // Идентификатор платежа в системе банка. По офф. документации это number(20), но фактически значение передается в виде строки
	ClientIP  string   `json:"IP,omitempty"`      // IP-адрес покупателя
	Amount    uint64   `json:"Amount,omitempty"`  // Сумма возврата в копейках
	Receipt   *Receipt `json:"Receipt,omitempty"` // Чек
}

func (i *CancelRequest) GetValuesForToken() map[string]string {
	v := map[string]string{
		"PaymentId": i.PaymentID,
		"IP":        i.ClientIP,
	}
	serializeUintToMapIfNonEmpty(&v, "Amount", i.Amount)
	return v
}

type CancelResponse struct {
	BaseResponse
	OriginalAmount uint64 `json:"OriginalAmount"` // Сумма в копейках до операции отмены
	NewAmount      uint64 `json:"NewAmount"`      // Сумма в копейках после операции отмены
	OrderID        string `json:"OrderId"`        // Номер заказа в системе Продавца
	Status         string `json:"Status"`         // Статус транзакции
	PaymentID      string `json:"PaymentId"`      // Уникальный идентификатор транзакции в системе Банка
}

// Cancel cancels the payment
// Deprecated: use CancelWithContext instead
func (c *Client) Cancel(request *CancelRequest) (*CancelResponse, error) {
	return c.CancelWithContext(context.Background(), request)
}

// CancelWithContext cancels the payment
func (c *Client) CancelWithContext(ctx context.Context, request *CancelRequest) (*CancelResponse, error) {
	response, err := c.PostRequestWithContext(ctx, "/Cancel", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res CancelResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	return &res, res.Error()
}
