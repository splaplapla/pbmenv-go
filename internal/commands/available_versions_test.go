package commands

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestAvailableVersions(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(strings.NewReader(
					`[
					{ "name": "v1.0.0" },
					{ "name": "v1.0.1" },
					{ "name": "v2.0.0" }
					]`,
				)),
			}, nil
		},
	}

	expected := []string{"1.0.0", "1.0.1", "2.0.0"}
	versions, _ := AvailableVersions(mockClient)

	if len(versions) != len(expected) {
		t.Errorf("Expected %d versions, got %d", len(expected), len(versions))
	}

	for i, v := range versions {
		if v != expected[i] {
			t.Errorf("Expected '%s', got '%s'", expected[i], v)
		}
	}
}

func TestAvailableVersions_NotJson(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader("not json")),
			}, nil
		},
	}

	_, err := AvailableVersions(mockClient)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAvailableVersions_ErrorCase(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			}, nil
		},
	}

	_, err := AvailableVersions(mockClient)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
