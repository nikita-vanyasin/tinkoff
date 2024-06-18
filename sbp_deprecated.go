package tinkoff

import (
	"context"
	"errors"
)

// The whole file contains invalid signatures and URL. It should be removed in the next major release

// Deprecated: use GetSBPPayTestRequest instead
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

// Warning! This method never worked.
// Deprecated: use SBPPayTestWithContext instead
func (c *Client) SPBPayTest(_ *GetQRTestRequest) (*GetQRResponse, error) {
	return nil, errors.New("SPBPayTest method never worked. Use SBPPayTestWithContext instead")
}

// Warning! This method never worked.
// Deprecated: use SBPPayTestWithContext instead
func (c *Client) SPBPayTestWithContext(_ context.Context, _ *GetQRTestRequest) (*GetQRResponse, error) {
	return nil, errors.New("SPBPayTestWithContext method never worked. Use SBPPayTestWithContext instead")
}
