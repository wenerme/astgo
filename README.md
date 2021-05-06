# Asterisk to Go

Libs for Go to work with Asterisk

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

---

# Old

## Features

* AMI
  * Async/Sync Command
  * Event Handler
  * Auto reconnect
  * Generated Command with document
  * Generated Event with document
  * Generated Command response with document
* ADB
  * Go ORM for Asterisk Realtime Database
  * Based on gorm

## AMI

### Basic

```go
package main

import (
  "fmt"
  "github.com/wenerme/astgo/ami"
  "time"
)

func main() {
  ch := make(chan *ami.Command, 1024)

  go func() {
    // Event handler
    for {
      c, ok := <-ch
      if !ok {
        break
      }
      fmt.Println("Recv Event Name: " + c.Name())
    }
  }()

  con, err := ami.Dial("192.168.1.2:5038", ami.DialConf{
    Username:  "admin",
    Secret:    "admin",
    Listeners: []chan<- *ami.Command{ch},
  })
  if err != nil {
    panic(err)
  }

  // Sync, generated action type, generated response type
  fmt.Println("PING: ", con.WriteCommandResponse(ami.PingAction{}).(*ami.PingResponse).Ping)

  // Change to you own number
  res, err := con.WriteCommandSync(ami.OriginateAction{
    Channel:  "sip/9001!9001@wener.me!9002@wener.me",
    CallerID: "wener <9001>",
    Exten:    "9002",
    Context:  "public",
    Priority: "1",
    Async:    "true",
  })

  if err != nil {
    // Response Error
    panic(err)
  }
  fmt.Println("Originate ", res.Message())

  // Will log some event
  <-time.After(time.Second * 30)
}
```
