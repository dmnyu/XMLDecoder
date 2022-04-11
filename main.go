package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Document struct {
	XMLName  xml.Name `xml:"doc"`
	Head     xml.Name `xml:"head"`
	Contents []Mixed  `xml:",any"`
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

	var doc Document
	if err := xml.Unmarshal([]byte(bytes), &doc); err != nil {
		panic(err)
	}

	fmt.Println(doc)
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
