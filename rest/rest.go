package rest

import (
	"context"
	"crypto/tls"
	"net/http"

	"gopkg.in/resty.v1"
)

type RestClient interface {
	Post(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	Get(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error)
	Put(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	Delete(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error)
	PostFormData(ctx context.Context, path string, headers http.Header, payload map[string]string) (body []byte, statusCode int, err error)
}

type client struct {
	options    Options
	httpClient *resty.Client
}

func New(options Options) RestClient {
	httpClient := resty.New()

	if options.SkipTLS {
		httpClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if options.WithProxy {
		httpClient.SetProxy(options.ProxyAddress)
	} else {
		httpClient.RemoveProxy()
	}

	httpClient.SetTimeout(options.Timeout)
	httpClient.SetDebug(options.DebugMode)

	return &client{
		options:    options,
		httpClient: httpClient,
	}
}

func (c *client) Post(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetBody(payload)
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}

	httpResp, httpErr := request.Post(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) PostFormData(ctx context.Context, path string, headers http.Header, payload map[string]string) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetFormData(payload)
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}

	httpResp, httpErr := request.Post(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}


func (c *client) Get(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}

	httpResp, httpErr := request.Get(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) Put(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetBody(payload)
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}

	httpResp, httpErr := request.Put(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) Delete(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}

	httpResp, httpErr := request.Delete(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}
