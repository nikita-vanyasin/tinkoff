package tinkoff

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

func (c *Client) Resend() (*ResendResponse, error) {
	response, err := c.PostRequest("/Resend", &ResendRequest{})
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
