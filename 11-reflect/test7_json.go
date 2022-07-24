package main

import (
	"encoding/json"
	"fmt"
)

/*
结构体json序列化  以及json 反序列化 结构体
*/

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title:  "喜剧之王",
		Year:   2000,
		Price:  10,
		Actors: []string{"xingye", "zhangbozhi"},
	}
	//编码的过程   结构体 ===> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr %s\n", jsonStr)

	//解码的过程		jsonStr ===> 结构体
	//jsonStr = jsonStr {"title":"喜剧之王","year":2000,"price":10,"actors":["xingye","zhangbozhi"]}
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("title = %s, year = %v, price = %v, actors = %v\n", my_movie.Title, my_movie.Year, my_movie.Price, my_movie.Actors)
}
