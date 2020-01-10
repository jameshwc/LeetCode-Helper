package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	leetCodeJSON = "leetcode.json"
)

type user struct {
	UserName  string `json:"user_name"`
	NumSolved int    `json:"num_solved"`
}

func leetcodeInit() {
	res, err := http.Get("https://leetcode.com/api/problems/algorithms/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	html, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(html))
	var u user
	err = json.Unmarshal(html, &u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
