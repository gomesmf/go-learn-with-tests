package main

import (
	"log"
	"os"

	blogposts "github.com/gomesmf/go-learning/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", posts)
}
