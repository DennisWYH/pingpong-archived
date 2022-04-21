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


func createGraph() {
// create a new user
	var u ent.User
	u.Name = "testUser"
	u.Password = "123"

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	userA, _ := client.User.Create().SetName(u.Name).SetPassword(u.Password).Save(ctx)

// create a new sentence
	var s ent.Sentense
	s.Chinese = "今天天气真好"
	s.Pinyin = "jin tian tian qi zhen hao"
	s.English = "Today's weather is so good"
	sentenceA, _ := client.Sentense.Create().SetChinese(s.Chinese).SetPinyin(s.Pinyin).SetEnglish(s.English).Save(ctx)

// create a read assignment
	readA, err := client.Read.Create().SetUser(userA).SetSentence(sentenceA).SetResult(0).Save(ctx)
	if err != nil {
		fmt.Println("error while createing the read.", err)
	}
	fmt.Println("readA is, ", readA)
}

//func queryGraph() {
//	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
//	if err != nil {
//		log.Fatalf("failed opening connection to sqlite: %v", err)
//	}
//	defer client.Close()
//	ctx := context.Background()
//}

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//queryUserYunhai()
	//createGraph()

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
