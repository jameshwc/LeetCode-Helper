package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

const (
	loginURL  = "https://leetcode.com/accounts/login"
	apiURL    = "https://leetcode.com/api/problems/algorithms/"
	userAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"
)

type leetCodeUser struct {
	User       userConfig
	Name       string `json:"user_name"`
	AC         int    `json:"num_solved"`
	ACeasy     int    `json:"ac_easy"`
	ACmedium   int    `json:"ac_medium"`
	AChard     int    `json:"ac_hard"`
	ACproblems []problem
}
type userConfig struct {
	Session, Csrftoken string
}
type problem struct {
}
type rawProblem struct {
	Stat struct {
		ID          int    `json:"frontend_question_id"`
		Title       string `json:"question__title"`
		AC          int    `json:"total_acs"`
		TotalSubmit int    `json:"total_submitted"`
	} `json:"stat"`
	Status     string `json:"status"`
	Difficulty struct {
		Level int `json:"level"`
	}
}

func (u *leetCodeUser) init() {
	if _, err := toml.DecodeFile("config.toml", u); err != nil {
		log.Fatal(err)
	}
}
func (u *leetCodeUser) saveJSON() ([]byte, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: u.User.Csrftoken})
	req.AddCookie(&http.Cookie{Name: "LEETCODE_SESSION", Value: u.User.Session})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("request not successfully")
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile("leetcode.json", bodyBytes, 0644)
	if err != nil {
		return nil, err
	}
	problemBytes := []byte("[")
	rawproblemBytes := bytes.Split(bytes.Split(bodyBytes, []byte("["))[1], []byte("]"))[0]
	problemBytes = append(problemBytes, rawproblemBytes...)
	problemBytes = append(problemBytes, []byte("]")...)
	err = ioutil.WriteFile("problems.json", problemBytes, 0644)
	return problemBytes, nil
}
func (u *leetCodeUser) parseJSON(b []byte) {
	json.Unmarshal(b, u)
}

func parseProblems(b []byte) {
	problem := []rawProblem{}
	json.Unmarshal(b, &problem)
	for i := range problem {
		if problem[i].Status == "ac" {
			fmt.Println(problem[i].Stat.Title)
		}
	}
}
func SaveJSON() {
	var u leetCodeUser
	u.init()
	data, err := u.saveJSON()
	if err != nil {
		log.Fatal(err)
	}
	parseProblems(data)
	// u.parseJSON(data)
}
