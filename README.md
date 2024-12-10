# Go SDK for (most) of [Buy Me A Coffee](https://www.buymeacoffee.com) [API](https://developers.buymeacoffee.com/#/apireference)

[![build-containers](https://github.com/DazWilkin/go-buymeacoffee/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/go-buymeacoffee/actions/workflows/build.yml)

Image: `ghcr.io/dazwilkin/go-buymeacoffee-server:ee566d6a23dde36052aa72bb93ccf2901539568c`

You'll need a token that you can generate in the [Developer Dashboard](https://developers.buymeacoffee.com/dashboard)

See [`cmd/main.go`](/cmd/main.go) or [`client/client_test.go`](/client/client_test.go) for examples

> **NOTE** The `types` aren't fully implemented.

## [Sigstore](https://www.sigstore.dev/)

`go-buymeacoffee` container images are being signed by Sigstore and may be verified:

```bash
cosign verify \
--key=./cosign.pub \
ghcr.io/dazwilkin/go-buymeacoffee:1234567890123456789012345678901234567890
```

NOTE `cosign.pub` may be downloaded [here](./cosign.pub)

To install cosign, e.g.:

```bash
go install github.com/sigstore/cosign/cmd/cosign@latest
```

## Test API Server

> **NOTE** `types/examples.go` revised slightly to replace `next_page_url` with `null`. Because BuyMeACoffee API uses URLs (`next_page_url`), it's more challenging to rewrite these in the client.

The SDK includes a test API server (`/cmd/server`) that implements the API and return the examples values from the [API documentation](https://developers.buymeacoffee.com/#/apireference)

```bash
PORT="8080"
ENDPOINT="0.0.0.0:${PORT}

go run ./cmd/server \
--endpoint="${ENDPOINT}"
```

Then:

```bash
PURCHASE="2621"
curl \
--silent \
--request GET \
http://${ENDPOINT}/extras

curl \
--silent \
--request GET \
http://${ENDPOINT}/extras/${PURCHASE}

SUBSCRIPTION="7979"
curl \
--silent \
--request GET \
http://${ENDPOINT}/subscriptions

curl \
--silent \
--request GET \
http://${ENDPOINT}/subscriptions/${SUBSCRIPTION}

SUPPORTER="245731"
curl \
--silent \
--request GET \
http://${ENDPOINT}/supporters

curl \
--silent \
--request GET \
http://${ENDPOINT}/supporters/${SUPPORTER}
```

<hr/>
<br/>
<a href="https://www.buymeacoffee.com/dazwilkin" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>
