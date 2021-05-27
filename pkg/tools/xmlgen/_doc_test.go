package xmlgen

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestDocParse(t *testing.T) {
	f, _ := os.Open("out.xml")
	defer f.Close()
	dec := xml.NewDecoder(f)
	var doc DocModel
	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}
	//for _, v := range doc.Agi {
	//	if v.Name == "control stream file" {
	//		indent, err := json.MarshalIndent(v, "", "  ")
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Println(string(indent))
	//		return
	//	}
	//}

	ge := NewAGIGenerator()
	ge.Debug = true
	d := &Model{}
	BuildModel(&doc, d)

	{
		b, err := json.MarshalIndent(doc, "", "  ")
		assert.NoError(t, err)
		err = os.WriteFile("doc.json", b, 0644)
		assert.NoError(t, err)
	}
	{
		b, err := json.MarshalIndent(d, "", "  ")
		assert.NoError(t, err)
		err = os.WriteFile("gen.json", b, 0644)
		assert.NoError(t, err)
	}

	err := ge.Generate(context.Background(), d)
	if err != nil {
		panic(err)
	}
	err = ge.Write(WriteConfig{
		Target: "./../../../",
		//DryRun: true,
	})
	if err != nil {
		panic(err)
	}
	log.Println("AGI Commands", len(d.AGICommands))
}

//func TestName(t *testing.T) {
//	println(fieldName("Variablename"))
//	println(fieldName("Keytree"))
//}
