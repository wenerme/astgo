package xmlgen

//go:generate gomodifytags -file=doc_test.go -w -all -add-tags json -transform camelcase --skip-unexported -add-options json=omitempty
import (
	"context"
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
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

	err := ge.Generate(context.Background(), d)
	if err != nil {
		panic(err)
	}
	err = ge.Write(WriteConfig{
		Target: "./../../../agi/agimodels",
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
