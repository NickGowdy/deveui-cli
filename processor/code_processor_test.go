package processor

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	DoPost func(url string, contentType string, body io.Reader) (resp *http.Response, err error)
}

func TestCanProcessCodes(t *testing.T) {
	client := &MockClient{
		DoPost: func(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
			return &http.Response{}, nil
		},
	}

	CodeProcessor := &CodeProcessor{
		CodeRegistrationLimit: 10,
		MaxConcurrentJobs:     10,
		BaseUrl:               "http://www.mock-url.com",
		Client:                client,
	}

	registeredDevices := CodeProcessor.Process()

	if len(*registeredDevices) != 10 {
		t.Errorf("expecting 10 registered devices, but have: %d", len(*registeredDevices))
	}

}

func (m *MockClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(nil)),
			Status:     "200 OK"},
		nil
}
