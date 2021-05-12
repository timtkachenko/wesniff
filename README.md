# wesniff - simple IDV

## build
`$ go get github.com/gobuffalo/packr/v2/packr2`

`$ packr2 -v && go build -o ./build/wesniff cmd/wesniff-server/main.go`

## run

`./build/wesniff --port 33111`


### listen webhooks locally
using **ngrock** set up proxy to https://wesniff.example.com/api/events

```yaml
wesniff:
      proto: http
      addr: "33111"
      subdomain: wesniff
      host_header: example.com
```
