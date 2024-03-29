# Twilio Signature Calculator 🛠️

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)

![Build](https://github.com/kmesiab/go-twilio-compute/actions/workflows/go-build.yml/badge.svg)
![Build](https://github.com/kmesiab/go-twilio-compute/actions/workflows/go-lint.yml/badge.svg)
![Build](https://github.com/kmesiab/go-twilio-compute/actions/workflows/go-test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kmesiab/go-twilio-compute)](https://goreportcard.com/report/github.com/kmesiab/go-twilio-compute)

## Overview 🌐

This command-line tool is designed to compute the signature for Twilio API
requests. It helps ensure the integrity and authenticity of the requests
coming from Twilio by using HMAC-SHA1 hashing algorithm.

## Features ✨

- Compute the HMAC-SHA1 hash for Twilio API requests.
- Easy to use command-line interface.
- Supports different URL structures and parameter combinations.

## Installation 📥

To use this tool, you need to have Go installed on your system. Follow these
steps to set it up:

   ```bash
   go install github.com/kmesiab/go-twilio-compute@latest
   ```

or

1. Clone the repository:

   ```bash
   git clone https://github.com/kmesiab/go-twilio-compute.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-twilio-compute
   ```

3. Build the project:

   ```bash
   go build
   ```

## Usage 🚀

Run the tool with the required flags:

```bash
./go-twilio-compute

--url <url Twilio will call> 
--key <your Twilio API key>
--params <URL encoded body of the message>
```

Example:

```bash
./go-twilio-compute --url "https://example.com/api" --key "mysecretkey"
--params "Body=sample%20message%3F&From=%2B12533"
```

### Command-Line Switches 🎚️

- `--url`: The URL Twilio will call.
- `--key`: Your Twilio API key used to hash the message.
- `--params`: The URL-encoded body of the message.

### Example Outputs 📝

```bash
./go-twilio-compute --url http://foo.com --key MyFakeKey --params foo=bar

Computed Signature: w6A4pOIG/hOAQsWqwnZRblmo7vo=
```

## Running Tests 🧪

To run the unit tests:

```bash
go test
```

## Contributing 🤝

Contributions are welcome! Feel free to submit pull requests or open issues
to improve the tool or add new features.

## License 📜

This project is licensed under the MIT License - see the LICENSE file for
details.

## What This Is Not 🚫

This is not a replacement for
[Twilio validation](https://www.twilio.com/docs/usage/webhooks/webhooks-security#validating-signatures-from-twilio).
This tool enables you to generate valid signatures for testing your Twilio applications.

## Twilio Documentation 🔗

[Validating Signatures from Twilio](https://www.twilio.com/docs/usage/webhooks/webhooks-security#validating-signatures-from-twilio)
