# Asterisk to Go

Libs for Golang to work with Asterisk

* AMI
* AGI

## AMI

* Async Action
* Sync Action
* Event Subscribe
* Auto Reconnect
* Generated Action with document
* Generated Event with document
* [ ] Generated action Response with document
* [ ] Generated client - Request -> Response
* Improved doc
  * Proper handle para to doc

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
