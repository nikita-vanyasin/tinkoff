package tinkoff

import "context"

type GetStateRequest struct {
	BaseRequest

	PaymentID   string `json:"PaymentId"`    // Идентификатор платежа в системе банка. По офф. документации это number(20), но фактически значение передается в виде строки
	TerminalKey string `json:"TerminalKey"`  // Ключ терминала
	ClientIP    string `json:"IP,omitempty"` // IP-адрес покупателя
}

func (i *GetStateRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"IP":          i.ClientIP,
		"PaymentId":   i.PaymentID,
		"TerminalKey": i.TerminalKey,
	}
}

type GetStateResponse struct {
	BaseResponse
	OrderID   string `json:"OrderId"`   // Номер заказа в системе Продавца
	Status    string `json:"Status"`    // Статус платежа
	PaymentID string `json:"PaymentId"` // Уникальный идентификатор транзакции в системе Банка
}

// GetState returns info about payment
// Deprecated: use GetStateWithContext instead
func (c *Client) GetState(request *GetStateRequest) (*GetStateResponse, error) {
	return c.GetStateWithContext(context.Background(), request)
}

// GetStateWithContext returns info about payment
func (c *Client) GetStateWithContext(ctx context.Context, request *GetStateRequest) (*GetStateResponse, error) {
	response, err := c.PostRequestWithContext(ctx, "/GetState", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res GetStateResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	return &res, res.Error()
}
