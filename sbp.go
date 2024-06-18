package tinkoff

import (
	"context"
)

type SBPPayTestRequest struct {
	BaseRequest

	PaymentID         string `json:"PaymentId"`         // Идентификатор платежа в системе банка. По офф. документации это number(20), но фактически значение передается в виде строки
	IsDeadlineExpired bool   `json:"IsDeadlineExpired"` // Признак эмуляции отказа проведения платежа Банком по таймауту. true – требуется эмуляция (не	может быть использован вместе с IsRejected = true)
	IsRejected        bool   `json:"IsRejected"`        // Признак эмуляции отказа Банка в проведении платежа. true – требуется эмуляция (не может быть использован вместе с IsDeadlineExpired = true)
}

func (i *SBPPayTestRequest) GetValuesForToken() map[string]string {
	return map[string]string{
		"PaymentId":         i.PaymentID,
		"TerminalKey":       i.TerminalKey,
		"IsDeadlineExpired": serializeBool(i.IsDeadlineExpired),
		"IsRejected":        serializeBool(i.IsRejected),
	}
}

type SBPPayTestResponse struct {
	BaseResponse
}

func (c *Client) SBPPayTestWithContext(ctx context.Context, request *SBPPayTestRequest) (*SBPPayTestResponse, error) {
	response, err := c.PostRequestWithContext(ctx, "/SbpPayTest", request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res SBPPayTestResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	return &res, res.Error()
}
