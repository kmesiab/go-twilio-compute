package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"

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
		signature := ComputeHash(argv.URL, argv.ApiKey, argv.Params)
		fmt.Println("Computed Signature:", signature)
		return nil
	}))
}

func ComputeHash(URL, key, params string) string {
	hmacData := hmac.New(sha1.New, []byte(key))
	hmacData.Write([]byte(URL + params))
	return base64.StdEncoding.EncodeToString(hmacData.Sum(nil))
}
