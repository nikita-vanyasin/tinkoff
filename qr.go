package tinkoff

type GetQRRequest struct {
	BaseRequest

	PaymentID string `json:"PaymentId"` // Идентификатор платежа в системе банка. По офф. документации это number(20), но фактически значение передается в виде строки
	DataType  string `json:"DataType"`  //Тип возвращаемых данных. PAYLOAD – В ответе возвращается только Payload (по-умолчанию). IMAGE – В ответе возвращается SVG изображение QR
}

func (i *GetQRRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"PaymentId":   i.PaymentID,
		"TerminalKey": i.TerminalKey,
	}
}

type GetQRResponse struct {
	BaseResponse
	OrderID   string `json:"OrderId"`   // Номер заказа в системе Продавца
	Data      string `json:"Data"`      // Payload - или SVG
	PaymentID int    `json:"PaymentId"` // Идентификатор платежа в системе банка.
}

type GetQRTestRequest struct {
	BaseRequest

	PaymentID         string `json:"PaymentId"`         // Идентификатор платежа в системе банка. По офф. документации это number(20), но фактически значение передается в виде строки
	IsDeadlineExpired bool   `json:"IsDeadlineExpired"` // Признак эмуляции отказа проведения платежа Банком по таймауту. true – требуется эмуляция (не	может быть использован вместе с IsRejected = true)
	IsRejected        bool   `json:"IsRejected"`        // Признак эмуляции отказа Банка в проведении платежа. true – требуется эмуляция (не может быть использован вместе с IsDeadlineExpired = true)
}

func (i *GetQRTestRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"PaymentId":   i.PaymentID,
		"TerminalKey": i.TerminalKey,
	}
}

func (c *Client) GetQR(request *GetQRRequest) (*GetQRResponse, error) {
	response, err := c.PostRequest("/GetQr", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res GetQRResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	return &res, res.Error()
}

// SPBPayTest тестовый метод описанный в API.
// на рабочем терминале - функция не работает.
// тестовый терминал не работает у банка.
func (c *Client) SPBPayTest(request *GetQRTestRequest) (*GetQRResponse, error) {
	response, err := c.PostRequest("/SpbPayTest", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res GetQRResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	return &res, res.Error()
}
