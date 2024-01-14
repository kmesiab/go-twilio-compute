package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHash(t *testing.T) {
	testCases := []struct {
		url, key, params, expectedHash string
	}{
		{"https://example.com/api", "testkey1", "param1=value1&param2=value2", "kbLovgPV3UKQnVNhgGIy5Yxp95g="},
		{"https://example.com/api/path", "testkey2", "param1=value1", "PF0HbycGbiprxUu4YKsKkBSekyI="},
		{"http://example.com/api", "testkey3", "param1=value1&param2=value2&param3=value3", "mRleMeU7j4eBEGgGB+l6Vcut8+0="},
		{"https://example.com/api", "testkey4", "", "nVBfJfYaf9oXq28/ggQq6J3XIbo="},
		{"https://example.com/api/path/segment", "testkey5", "param1=value1&param2=value2&param3=value3", "qrI/fD6EdICvz1xnEP1ZJiuKSVo="},
	}

	for _, tc := range testCases {
		t.Run("TestComputeHash", func(t *testing.T) {
			computedHash := ComputeHash(tc.url, tc.key, tc.params)
			assert.Equal(t, tc.expectedHash, computedHash, "Computed hash should match the expected hash")
		})
	}
}

func TestAddPort(t *testing.T) {
	testCases := []struct {
		url, expectedURL string
	}{
		{"https://example.com", "https://example.com:443"},
		{"http://example.com", "http://example.com:80"},
		{"https://example.com:1234", "https://example.com:1234"},
		{"http://example.com/path", "http://example.com:80/path"},
		{"https://example.com/api?query=param", "https://example.com:443/api?query=param"},
	}

	for _, tc := range testCases {
		t.Run("TestAddPort", func(t *testing.T) {
			resultURL := addPort(tc.url)
			assert.Equal(t, tc.expectedURL, resultURL, "URL with added port should match the expected URL")
		})
	}
}

func TestRemovePort(t *testing.T) {
	testCases := []struct {
		url, expectedURL string
	}{
		{"https://example.com:443", "https://example.com"},
		{"http://example.com:80", "http://example.com"},
		{"https://example.com:1234", "https://example.com"},
		{"http://example.com:8080/path", "http://example.com/path"},
		{"https://example.com:443/api?query=param", "https://example.com/api?query=param"},
	}

	for _, tc := range testCases {
		t.Run("TestRemovePort", func(t *testing.T) {
			resultURL := removePort(tc.url)
			assert.Equal(t, tc.expectedURL, resultURL, "URL with removed port should match the expected URL")
		})
	}
}
