package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name "xml:plant"
	Id      int      "xml:id.attr"
	Name    string   "xml:name"
	Origin  []string "sml:origin"
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant Id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	cofee := &Plant{Id: 27, Name: "Cofee"}
	cofee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(cofee, " ", " ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{cofee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
