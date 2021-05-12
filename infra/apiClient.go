package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/go-cleanhttp"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	*http.Response
	Body []byte
}

func (r *Response) apiError() error {
	if r.StatusCode == 200 || r.StatusCode == 201 {
		if r.Header.Get("Content-Type") == "application/json" {
			var b interface{}
			err := json.Unmarshal(r.Body, &b)
			if err != nil {
				return fmt.Errorf("%s: %s", string(r.Body), err.Error())
			}
		}
	}
	if r.StatusCode > 299 {
		return fmt.Errorf("%s (%d) %s", http.StatusText(r.StatusCode), r.StatusCode, string(r.Body))
	}
	return nil
}

func (r *Response) unpackBody() (err error) {
	if len(r.Body) != 0 {
		return
	}
	reader := r.Response.Body

	defer func() { _ = r.Response.Body.Close() }()

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	r.Body = buf
	return
}

type ApiClient interface {
	Execute(options ApiClientOptions, args ...interface{}) (*Response, error)
	ExecuteWithContext(ctx context.Context, options ApiClientOptions, obj interface{}) (*Response, error)
}
type apiClient struct {
	Url     string
	timeout time.Duration
}
type ApiClientOptions struct {
	UrlSuffix  string
	Method     string
	Headers    map[string]string
	EncodeJson bool
}

func NewApiClient(url string, timeout time.Duration) ApiClient {
	return &apiClient{url, timeout}
}

func (ac *apiClient) Execute(options ApiClientOptions, args ...interface{}) (res *Response, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), ac.timeout)
	defer func() {
		if err == nil {
			err = ctx.Err()
		}
		cancel()
	}()

	if len(args) > 1 {
		return nil, fmt.Errorf("unsupported extra args, %x", args)
	}
	var obj interface{}
	if len(args) == 1 {
		obj = args[0]
	}
	return ac.ExecuteWithContext(ctx, options, obj)
}
func (ac *apiClient) ExecuteWithContext(ctx context.Context, options ApiClientOptions, obj interface{}) (*Response, error) {
	var buf bytes.Buffer
	var body []byte
	var err error
	if obj != nil {
		if options.EncodeJson {
			body, err = json.Marshal(obj)
			if err != nil {
				return nil, err
			}
		} else {
			if objStr, ok := obj.(string); ok {
				body = []byte(objStr)
			}
		}
		_, err := buf.Write(body)
		if err != nil {
			return nil, err
		}
	}
	if options.Method == "" {
		options.Method = http.MethodGet
	}
	url := ac.Url + options.UrlSuffix
	req, err := http.NewRequest(options.Method, url, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")
	if len(options.Headers) != 0 {
		for k, v := range options.Headers {
			req.Header.Set(k, v)
		}
	}
	req = req.WithContext(ctx)

	httpClient := cleanhttp.DefaultClient()
	httpClient.Timeout = ac.timeout
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("no response")
	}

	res := &Response{Response: resp}
	err = res.unpackBody()
	if err != nil {
		return nil, err
	}

	err = res.apiError()
	if err != nil {
		return nil, err
	}
	return res, nil
}
