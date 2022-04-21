package database

import (
	"context"
	"fmt"
	"log"
	"pingpong2/ent"
)

func CreateTestGraph() {
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
	_, err = client.Read.Create().SetUser(userA).SetSentence(sentenceA).SetResult(0).Save(ctx)
	if err != nil {
		fmt.Println("error while createing the read.", err)
	}
}

func AddTenSentences(){
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	client.Sentense.Create().SetChinese("我觉得天气很好").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("我可以出门吗？").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("这是我的苹果").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("你叫什么名字").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("我的名字是海").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("你怎么样？").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("我今年20岁").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("谢谢，不客气").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("今天的课很好").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
	client.Sentense.Create().SetChinese("明天我们还上课吗？").SetPinyin("wo jue de tian qi hen hao").SetEnglish("I think the weather is very good").Save(ctx)
}

