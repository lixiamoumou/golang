package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.Marshal(coffee)
	fmt.Println(string(out))
	fmt.Println()

	out1, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out1))
	fmt.Println()

	fmt.Println(xml.Header + string(out1))
	fmt.Println()

	var p Plant
	if err := xml.Unmarshal(out1, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Println()

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		// xml和后面的内容没有空格
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))

}
