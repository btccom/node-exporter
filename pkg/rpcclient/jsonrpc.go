package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

type JSONRPCClient struct {
	client http.Client
	Debug  bool
}

// NewJSONRPCClient 初始化客户端
func NewJSONRPCClient() *JSONRPCClient {
	client := http.Client{Timeout: 10 * time.Second}
	return &JSONRPCClient{client: client}
}

// SetTimeout 重置超时时间
func (c *JSONRPCClient) SetTimeout(t time.Duration) {
	c.client.Timeout = t
}

type Request struct {
	JSONRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
	ID      interface{}       `json:"id"`
}

// NewRequest returns a new JSON-RPC 1.0 request object given the provided id,
// method, and parameters.  The parameters are marshaled into a json.RawMessage
// for the Params field of the returned request object.  This function is only
// provided in case the caller wants to construct raw requests for some reason.
//
// Typically callers will instead want to create a registered concrete command
// type with the NewCmd or New<Foo>Cmd functions and call the MarshalCmd
// function with that command to generate the marshaled JSON-RPC request.
func NewRequest(id interface{}, method string, params []interface{}) (*Request, error) {
	rawParams := make([]json.RawMessage, 0, len(params))
	for _, param := range params {
		marshalledParam, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		rawMessage := json.RawMessage(marshalledParam)
		rawParams = append(rawParams, rawMessage)
	}

	return &Request{
		JSONRPC: "1.0",
		ID:      id,
		Method:  method,
		Params:  rawParams,
	}, nil
}

// Post : post 方法
func (c *JSONRPCClient) Post(endpoint string, request *Request, username, password string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(data)
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, password)

	// 输出到 log
	if c.Debug {
		byts, _ := httputil.DumpRequest(req, true)
		fmt.Println(string(byts))
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.Debug {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}
