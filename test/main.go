package main

import "fmt"

func main() {
	//infra.Init()
	//collection := infra.Mongodb.Database("bk").Collection("article")
	//data := model.Article{AuthorId: 332, Title: "wzhwzhwzh"}
	//res, err := collection.InsertOne(context.Background(), data)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("res ID:", res.InsertedID)
	m := make(map[string]string)
	m["wzh"] = "666"
	m["xj"] = "888"
	fmt.Println(m)
}
