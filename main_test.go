package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHash(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.url, func(t *testing.T) {
			hash := ComputeHash(tc.url, tc.key, tc.params)
			assert.Equal(t, tc.expectedHash, hash, "Hash should match the expected value")
		})
	}
}

var cases = []struct {
	url, key, params, expectedHash string
}{
	// Test with basic URL and parameters
	{"https://example.com/api", "mysecretkey", "param1=value1&param2=value2", "tYQQ58Oi6gab/WpmurMuMHXayq4="},

	// Test with different key
	{"https://example.com/api", "anotherkey", "param1=value1&param2=value2", "q0c+MOY2Kvzi2LNPK+x8/v9qfEg="},

	// Test with URL having additional path segments
	{"https://example.com/api/path/segment", "mysecretkey", "param1=value1&param2=value2", "eIesI/BxOCA7OddDitgisj4cO1I="},

	// Test with empty parameters
	{"https://example.com/api", "mysecretkey", "", "bYqj3xMDrpzjGDUhXSW0OQL/tvg="},

	// Test with no parameters and trailing slash in URL
	{"https://example.com/api/", "mysecretkey", "", "rg45cX+fAbzs7Xsc8X3gSJSnR7w="},

	// Test with a longer key
	{"https://example.com/api", "aVeryLongSecretKeyThatIsComplex", "param1=value1&param2=value2", "Ov5Q6q4ZqwC44eeu+VfLig4Q03I="},

	// Test with special characters in the key
	{"https://example.com/api", "key$pecial*Char!", "param1=value1&param2=value2", "Onrw0fZd5QjFJCSQdd68j2yHsA4="},

	// Test with special characters in the parameters
	{"https://example.com/api", "mysecretkey", "param1=special!Value&param2=another@Value", "8i7v6MKse3TFpK+7PDRz3XPPOk4="},

	// Test with a single parameter
	{"https://example.com/api", "mysecretkey", "singleParam=value", "BCKsEQwQKJbkyvSPS2dZiJITaMg="},

	// Test with URL query parameters
	{"https://example.com/api?query=param", "mysecretkey", "param1=value1&param2=value2", "RI+ct9b4q/73Gzz1muTE9RD3QW8="},
}
