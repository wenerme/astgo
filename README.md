# Asterisk to Go [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]  [![Go Report Card][report-card-img]][report-card]

Libs for Golang to work with Asterisk

* AMI
* AGI

[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://pkg.go.dev/github.com/wenerme/astgo

[ci-img]: https://github.com/wenerme/astgo/actions/workflows/ci.yml/badge.svg
[ci]: https://github.com/wenerme/astgo/actions/workflows/ci.yml

[cov-img]: https://codecov.io/gh/wenerme/astgo/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/wenerme/astgo/branch/master

[report-card-img]: https://goreportcard.com/badge/github.com/wenerme/astgo
[report-card]: https://goreportcard.com/report/github.com/wenerme/astgo

## AMI

* Async Action
* Sync Action
* Event Subscribe
* Auto Reconnect
* Generated Action with document
* Generated Event with document
* Generated Client - Action -> Response

```go
package main

import (
  "context"
  "fmt"
  "github.com/wenerme/astgo/ami"
)

func main() {
  boot := make(chan *ami.Message, 1)

  conn, err := ami.Connect(
    "192.168.1.1:5038",
    ami.WithAuth("admin", "admin"), // AMI auth
    // add predefined subscriber
    ami.WithSubscribe(ami.SubscribeFullyBootedChanOnce(boot)),
    ami.WithSubscribe(func(ctx context.Context, msg *ami.Message) bool {
      fmt.Println(msg.Format()) // log everything
      return true               // keep subscribe
    }, ami.SubscribeSend(), // subscribe send message - default recv only
    ))
  if err != nil {
    panic(err)
  }
  <-boot
  // AMI now FullyBooted
  _ = conn
}
```

## AGI

* FastAGI
* AGI Bin
* Generated Command
* Generated Client

```go
package main

import (
	"github.com/wenerme/astgo/agi"
)

func main() {
	// agi.Listen for FastAGI
	agi.Run(func(session *agi.Session) {
		client := session.Client()
		client.Answer()
		client.StreamFile("activated", "#")
		client.SetVariable("AGISTATUS", "SUCCESS")
		client.Hangup()
	})
}
```

## Asterisk Database Model
* GORM Based Model

## Roadmap
* [-] Asterisk Database Model
* [ ] Stasis
* [ ] ARI
