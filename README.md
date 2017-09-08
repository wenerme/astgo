# Asterisk to Go

Libs for Go to work with Asterisk 

## Features
* AMI
    * Async/Sync Command
    * Event Handler
    * Generated Command with document
    * Generated Event with document
    * Generated Command response with document

## AMI

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

	con, err := ami.Dial("192.168.1.2:5038", ami.Config{
		Username:  "admin",
		Secret:    "admin",
		Listeners: []chan<- *ami.Command{ch},
	})
	if err != nil {
		panic(err)
	}

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