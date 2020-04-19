package tinkoff

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type Client struct {
	terminalKey string
	password    string
	baseURL     string
}

func NewClient(terminalKey, password string) *Client {
	return &Client{
		terminalKey: terminalKey,
		password:    password,
		baseURL:     "https://securepay.tinkoff.ru/v2",
	}
}

// SetBaseURL allows to change default API endpoint
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) decodeResponse(response *http.Response, result interface{}) error {
	return json.NewDecoder(response.Body).Decode(result)
}

// PostRequest will automatically sign the request with token
// Use BaseRequest type to implement any API request
func (c *Client) PostRequest(url string, request RequestInterface) (*http.Response, error) {
	c.secureRequest(request)
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(c.baseURL+url, "application/json", bytes.NewReader(data))
	return resp, err
}

func (c *Client) secureRequest(request RequestInterface) {
	request.SetTerminalKey(c.terminalKey)

	v := request.GetValuesForToken()
	v["TerminalKey"] = c.terminalKey
	v["Password"] = c.password
	request.SetToken(generateToken(v))
}

func generateToken(v map[string]string) string {
	keys := make([]string, 0)
	for key := range v {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var b bytes.Buffer
	for _, key := range keys {
		b.WriteString(v[key])
	}
	sum := sha256.Sum256(b.Bytes())
	return fmt.Sprintf("%x", sum)
}
