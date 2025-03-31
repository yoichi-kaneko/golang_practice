// xmlの解析
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
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var post Post
	err = xml.Unmarshal(xmlData, &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}
