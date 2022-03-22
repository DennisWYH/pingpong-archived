package main

import (
	"fmt"
	"net/http"

	"entgo.io/ent/examples/fs/ent"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

type Sentence struct {
	chineseSentence []string
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "yunhaiwang"
	password = "koekje123"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("hello")
}

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang password=yes")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected")
	}

	router := httprouter.New()
	router.GET("/", index)
	http.ListenAndServe(":8080", router)
}
