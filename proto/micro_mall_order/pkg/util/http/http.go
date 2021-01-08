package http

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitee.com/kelvins-io/common/gtool"
)

const (
	defaultTimeoutHTTP = 5
	maxTimeoutHTTP     = 30
	contentTypeForm    = "application/x-www-form-urlencoded"
	contentTypeJson    = "application/json"
)

type Client struct {
	tracingClient *gtool.TracingClient
}

func NewClient(timeOutCustom int) *Client {
	timeOut := defaultTimeoutHTTP
	if timeOutCustom > 0 && timeOutCustom < maxTimeoutHTTP {
		timeOut = timeOutCustom
	}
	return &Client{
		tracingClient: &gtool.TracingClient{
			HttpClient: &http.Client{Timeout: time.Duration(timeOut) * time.Second},
		},
	}
}

func (h *Client) PostForm(ctx context.Context, url string, data url.Values) ([]byte, error) {

	return h.Post(ctx, url, data, contentTypeForm)
}

func (h *Client) PostJson(ctx context.Context, url string, data url.Values) ([]byte, error) {

	return h.Post(ctx, url, data, contentTypeJson)
}

func (h *Client) Post(ctx context.Context, url string, data url.Values, contextType string) ([]byte, error) {

	resp, err := h.tracingClient.HttpPost(ctx, url, contextType, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
