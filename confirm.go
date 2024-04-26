package tinkoff

import "context"

type ConfirmRequest struct {
	BaseRequest
	PaymentID string   `json:"PaymentId"`         // Идентификатор платежа в системе банка
	Amount    uint64   `json:"Amount,omitempty"`  // Сумма в копейках
	ClientIP  string   `json:"IP,omitempty"`      // IP-адрес покупателя
	Receipt   *Receipt `json:"Receipt,omitempty"` // Чек
}

func (i *ConfirmRequest) GetValuesForToken() map[string]string {
	v := map[string]string{
		"PaymentId": i.PaymentID,
		"IP":        i.ClientIP,
	}
	serializeUintToMapIfNonEmpty(&v, "Amount", i.Amount)
	return v
}

type ConfirmResponse struct {
	BaseResponse
	OrderID   string `json:"OrderId"`   // Номер заказа в системе Продавца
	Status    string `json:"Status"`    // Статус транзакции
	PaymentID string `json:"PaymentId"` // Идентификатор платежа в системе банка.
}

// Confirm finalizes the payment
// Deprecated: use ConfirmWithContext instead
func (c *Client) Confirm(request *ConfirmRequest) (*ConfirmResponse, error) {
	return c.ConfirmWithContext(context.Background(), request)
}

// ConfirmWithContext finalizes the payment
func (c *Client) ConfirmWithContext(ctx context.Context, request *ConfirmRequest) (*ConfirmResponse, error) {
	response, err := c.PostRequestWithContext(ctx, "/Confirm", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res ConfirmResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	return &res, res.Error()
}
