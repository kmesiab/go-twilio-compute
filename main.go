package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/mkideal/cli"
)

type argT struct {
	URL    string `cli:"*url" usage:"the url twilio will call"`
	ApiKey string `cli:"*key" usage:"your twilio api key used to hash the message"`
	Params string `cli:"*params" usage:"the URL encoded body of the message. ex: Body=sample%20message%3F&From=%2B12533"`
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		signatureWithPort := ComputeHash(addPort(argv.URL), argv.ApiKey, argv.Params)
		signatureWithoutPort := ComputeHash(removePort(argv.URL), argv.ApiKey, argv.Params)
		fmt.Println("Computed Signature (with port):", signatureWithPort)
		fmt.Println("Computed Signature (without port):", signatureWithoutPort)
		return nil
	}))
}

func ComputeHash(URL, key, params string) string {
	hmacData := hmac.New(sha1.New, []byte(key))
	hmacData.Write([]byte(URL + params))
	return base64.StdEncoding.EncodeToString(hmacData.Sum(nil))
}

func addPort(urlStr string) string {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil || parsedUrl.Port() != "" {
		return urlStr
	}

	var defaultPort string
	switch parsedUrl.Scheme {
	case "https":
		defaultPort = "443"
	case "http":
		defaultPort = "80"
	default:
		return urlStr
	}

	// Check if the URL has a path (and optionally a query string) and include it in the final URL
	pathAndQuery := ""
	if parsedUrl.Path != "" || parsedUrl.RawQuery != "" {
		pathAndQuery = parsedUrl.RequestURI()
	}

	// Rebuild the URL with the port added to the host
	return fmt.Sprintf("%s://%s:%s%s", parsedUrl.Scheme, parsedUrl.Host, defaultPort, pathAndQuery)
}

func removePort(urlStr string) string {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil || parsedUrl.Port() == "" {
		return urlStr
	}
	return strings.Replace(urlStr, ":"+parsedUrl.Port(), "", 1)
}
