package tinkoff

import "context"

type ResendRequest struct {
	BaseRequest
}

func (r *ResendRequest) GetValuesForToken() map[string]string {
	return map[string]string{}
}

type ResendResponse struct {
	BaseResponse
	Count int `json:"Count"` // Количество сообщений, отправляемых повторно
}

// Resend requests to send unacknowledged notifications again
// Deprecated: use ResendWithContext instead
func (c *Client) Resend() (*ResendResponse, error) {
	return c.ResendWithContext(context.Background())
}

// ResResendWithContextend requests to send unacknowledged notifications again
func (c *Client) ResendWithContext(ctx context.Context) (*ResendResponse, error) {
	response, err := c.PostRequestWithContext(ctx, "/Resend", &ResendRequest{})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res ResendResponse
	err = c.decodeResponse(response, &res)
	if err != nil {
		return nil, err
	}

	return &res, res.Error()
}
