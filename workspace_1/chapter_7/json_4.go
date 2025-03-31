// jsonの生成
package main

import (
	"encoding/json"
	"fmt"
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
	post := Post{
		Id:      1,
		Content: "Hello, world!",
		Author: Author{
			Id:   2,
			Name: "John Doe",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "Great post!",
				Author:  "Alice",
			},
			Comment{
				Id:      4,
				Content: "Thanks for sharing!",
				Author:  "Bob",
			},
		},
	}

	jsonFile, err := os.Create("post_4.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonFile.Close()
}
