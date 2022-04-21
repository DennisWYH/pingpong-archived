package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"pingpong2/database"
	"pingpong2/ent"
	"pingpong2/ent/sentense"
	"pingpong2/util"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("welcome to index page.")
}
func Tokens_to_pinyins(tokens []string) []string {
	var pinyins []string
	for _, val := range tokens {
		pinyin := util.HanziToPinyins(val)
		pinyins = append(pinyins, pinyin)
	}
	return pinyins
}

func addSentence(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var s ent.Sentense
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestChinese := s.Chinese
	requestEnglish := s.English
	// for each article content, we first tokenize it
	tokens, err := util.Tokenizer(requestChinese)
	pinyins := Tokens_to_pinyins(tokens)
	pinyin := strings.Join(pinyins, "")

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	client.Sentense.Create().SetChinese(requestChinese).SetPinyin(pinyin).SetEnglish(requestEnglish).Save(ctx)
}

func displaySentence(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	sentences, err := client.Sentense.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
	t, _ := template.ParseFiles("static/display-sentence-all.html")
	t.Execute(w, sentences)
}

func addOne(value string) int {
	return 1
}

func displaySentenceCardViewByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	requestID, err := strconv.ParseInt(params.ByName("id"),10,64)
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	sentence, err := client.Sentense.Query().Where(sentense.ID(int(requestID))).Only(ctx)
	if err != nil {
		log.Fatalf("failed while querying the database: %v", err)
	}

	funcs := template.FuncMap{"addOne": addOne}
	t, _ := template.New("display-sentence-card-view.html").Funcs(funcs).ParseFiles("static/display-sentence-card-view.html")
	t.Execute(w, sentence)
}

func displaySentenceCardViewNext(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	currentID, err := strconv.ParseInt(params.ByName("id"),10,64)
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	nextID := int(currentID) + 1
	sentence, err := client.Sentense.Query().Where(sentense.ID(nextID)).Only(ctx)
	if err != nil {
		log.Fatalf("failed while querying the database: %v", err)
		sentence, _ = client.Sentense.Query().Where(sentense.ID(1)).Only(ctx)
	}
	t, _ := template.ParseFiles("static/display-sentence-card-view.html")
	t.Execute(w, sentence)
}

func main() {
	database.MigrateTablesWithDrop()
	database.CreateTestGraph()
	database.AddTenSentences()

	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("static"))
	router.GET("/", index)
	// private/internal
	router.POST("/addSentence", addSentence)
	// public/external
	router.GET("/displaySentence", displaySentence)
	router.GET("/displaySentenceCardView/:id/", displaySentenceCardViewByID)
	router.GET("/displaySentenceCardView/:id/next/", displaySentenceCardViewNext)
	http.ListenAndServe(":8080", router)
}
