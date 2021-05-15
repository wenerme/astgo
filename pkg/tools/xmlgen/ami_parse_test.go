package xmlgen

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, _ := os.Open("out.xml")
	defer f.Close()
	dec := xml.NewDecoder(f)
	var doc Docs
	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	ge := NewAMIGenerator()
	ge.Debug = true
	d := &AstDoc{}

	BuildAstDoc(&doc, d)

	err := ge.Generate(context.Background(), d)
	if err != nil {
		panic(err)
	}
	err = ge.Write(WriteConfig{
		Target: "./../../../ami/models",
		//DryRun: true,
	})
	if err != nil {
		panic(err)
	}
}

func BuildModel(in *DocModel, out *Model) {
	for _, v := range in.Agi {
		out.AGICommands = append(out.AGICommands, in.AgiModel(v))
	}
}
