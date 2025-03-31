// xmlの解析。Decoderを用いる
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post_2.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		switch se := t.(type) {
		case xml.StartElement:

			if se.Name.Local == "comment" {
				var comment Comment
				err := decoder.DecodeElement(&comment, &se)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(comment)
			}
		}
	}
}
