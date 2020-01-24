package helper

import (
	"fmt"
	"log"
)

func Helper() {
	var u leetCodeUser
	var t trendCSV
	u.init()
	data, err := u.saveJSON()
	if err != nil {
		log.Fatal(err)
	}
	u.parseProblems(data)
	if t.write(u) {
		fmt.Println("You have accomplished more problems!")
	}
	makeReadMe(u, t)
}
