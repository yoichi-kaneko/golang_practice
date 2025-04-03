package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	// データベースに接続する
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	// データベースからデータを取得する
	post = Post{}
	err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		return Post{}, err
	}
	return
}

func (post *Post) create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"

	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("UPDATE posts SET content = $1, author = $2 WHERE id = $3", post.Content, post.Author, post.Id)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("DELETE FROM posts WHERE id = $1", post.Id)
	return
}
