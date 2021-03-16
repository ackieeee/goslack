package goslack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	httpCli    *http.Client
	webhookURL string
}

type reqData struct {
	Text string `json:"text"`
}

func NewClient(url string) *Client {
	return &Client{
		httpCli:    &http.Client{},
		webhookURL: url,
	}
}

func (c *Client) Post(msg string) error {
	data := reqData{Text: msg}
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	c.httpCli.Post(c.webhookURL, "application/json", bytes.NewBuffer(buf))
	return nil
}
