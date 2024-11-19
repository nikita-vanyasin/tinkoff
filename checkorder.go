package tinkoff

import (
	"context"
)

type CheckOrderRequest struct {
	BaseRequest
	OrderID string `json:"OrderId"` // Номер заказа в системе Продавца
}

type CheckOrderResponse struct {
	BaseResponse
	OrderID  string `json:"OrderId"` // Номер заказа в системе Продавца
	Payments []Payment
}

type Payment struct {
	PaymentID    string `json:"PaymentId"`
	Amount       uint64 `json:"Amount,omitempty"`            // Стоимость товара в копейках
	Status       string `json:"Status"`            // Статус платежа
	RRN          string `json:"RRN,omitempty"`               // Внутренний номер операции в платежной системе — кроме операций по СБП.
	Success      bool   `json:"Success"`           // Успешность операции
	ErrorCode    string `json:"ErrorCode"`         // Код ошибки, «0» - если успешно
	ErrorMessage string `json:"Message,omitempty"` // Краткое описание ошибки
}

func (i *CheckOrderRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"OrderID": i.OrderID,
	}
}

func (c *Client) CheckOrderWithContext(ctx context.Context, request *CheckOrderRequest) (*CheckOrderResponse, error) {
	response, err := c.PostRequestWithContext(ctx, "/CheckOrder", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res CheckOrderResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	return &res, res.Error()
}
