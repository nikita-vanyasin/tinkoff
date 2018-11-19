package tinkoff

type BaseRequest struct {
	TerminalKey string `json:"TerminalKey"`
	Password    string `json:"Password"`
	Token       string `json:"Token"`
}

type RequestInterface interface {
	GetValuesForToken() map[string]string
	SetTerminalKey(key string)
	SetPassword(password string)
	SetToken(token string)
}

func (r *BaseRequest) SetTerminalKey(key string) {
	r.TerminalKey = key
}

func (r *BaseRequest) SetPassword(password string) {
	r.Password = password
}

func (r *BaseRequest) SetToken(token string) {
	r.Token = token
}
