package telegram

import (
	"awesomeProject3/lib/e"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
	// tg-bot.com/bot<token>
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBaseToken(token),
		client:   http.Client{},
	}
}

func newBaseToken(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Updates, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, e.Wrap("cant do request", err)
	}
	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("cant unmarshal response: %w", err)
	}
	return res.Result, nil
}

func (c *Client) SendMessage(ChatID int, Text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(ChatID))
	q.Add("text", Text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("cant send message", err)
	}
	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	const errMsg = "cant do request"
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	return body, nil
}
