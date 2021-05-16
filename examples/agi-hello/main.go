package main

import (
	"github.com/wenerme/astgo/agi"
)

func main() {
	agi.Run(func(session *agi.Session) {
		client := session.Client()
		client.Answer()
		client.StreamFile("activated", "#")
		client.SetVariable("AGISTATUS", "SUCCESS")
		client.Hangup()
	})
}
