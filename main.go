package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type EAD struct {
	Head     string  `xml:"head" json:"head"`
	Contents []Mixed `xml:",any" json:"contents"`
}

type Mixed struct {
	Type  string
	Value interface{}
}

func main() {
	bytes, err := ioutil.ReadFile("example.xml")
	if err != nil {
		panic(err)
	}

	var doc EAD
	if err := xml.Unmarshal([]byte(bytes), &doc); err != nil {
		panic(err)
	}

	jdoc, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jdoc))
}

func (m *Mixed) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "head", "p", "list":
		var e string
		if err := d.DecodeElement(&e, &start); err != nil {
			return err
		}
		m.Value = e
		m.Type = start.Name.Local
	default:
		return fmt.Errorf("unknown element: %s", start)
	}
	return nil
}
