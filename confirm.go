package tinkoff

import "strconv"

type ConfirmRequest struct {
	BaseRequest
	PaymentID string   `json:"PaymentId"`         // Идентификатор платежа в системе банка
	Amount    uint64   `json:"Amount,omitempty"`  // Сумма в копейках
	ClientIP  string   `json:"IP,omitempty"`      // IP-адрес покупателя
	Receipt   *Receipt `json:"Receipt,omitempty"` // Чек
}

func (i *ConfirmRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"Amount":    strconv.FormatUint(i.Amount, 10),
		"IP":        i.ClientIP,
		"PaymentId": i.PaymentID,
	}
}

type ConfirmResponse struct {
	BaseResponse
	OrderID   string `json:"OrderId"`   // Номер заказа в системе Продавца
	Status    string `json:"Status"`    // Статус транзакции
	PaymentID string `json:"PaymentId"` // Идентификатор платежа в системе банка.
}

func (c *Client) Confirm(request *ConfirmRequest) (*ConfirmResponse, error) {
	response, err := c.PostRequest("/Confirm", request)
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
