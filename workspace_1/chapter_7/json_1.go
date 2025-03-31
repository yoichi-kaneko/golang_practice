// jsonの解析
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post_1.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var post Post
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}
