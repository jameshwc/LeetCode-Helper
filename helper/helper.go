package helper

import (
	"fmt"
	"log"
)

type leetCodeUser struct {
	Connection struct {
		Session, Csrftoken string
	}
	Name       string `json:"user_name"`
	AC         int    `json:"num_solved"`
	ACeasy     int    `json:"ac_easy"`
	ACmedium   int    `json:"ac_medium"`
	AChard     int    `json:"ac_hard"`
	ACproblems []problem
}
type problem struct {
	NO         int
	Title      string
	Acceptance float64
	Difficulty string
	Language   []string
	tags       []string
}
type rawProblem struct {
	Stat struct {
		ID          int    `json:"frontend_question_id"`
		Title       string `json:"question__title"`
		TitleSlug   string `json:"question__title_slug"`
		AC          int    `json:"total_acs"`
		TotalSubmit int    `json:"total_submitted"`
	}
	Status     string
	Difficulty struct {
		Level int
	}
}

func ReadMeHelper() {
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
func Test() {
	fmt.Println(parseTags())
}
