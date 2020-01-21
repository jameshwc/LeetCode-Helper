package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	NO         string
	Title      string
	Acceptance float64
	Difficulty string
	Language   string
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

func num2String(n int) string {
	sn := strconv.Itoa(n)
	switch {
	case n < 10:
		return "000" + sn
	case n < 100:
		return "00" + sn
	case n < 1000:
		return "0" + sn
	}
	return sn
}
func level2String(n int) string {
	switch n {
	case 1:
		return "Easy"
	case 2:
		return "Medium"
	case 3:
		return "Hard"
	}
	return "Error"
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
	return problemBytes, nil
}
func (u *leetCodeUser) parseJSON(b []byte) {
	json.Unmarshal(b, u)
}
func (u *leetCodeUser) parseProblems(b []byte) {
	problems := []rawProblem{}
	json.Unmarshal(b, &problems)
	levelString := []string{"Easy", "Medium", "Hard"}
	for i := len(problems) - 1; i >= 0; i-- {
		if problems[i].Status == "ac" {
			var p problem
			p.Title = problems[i].Stat.Title
			p.NO = num2String(problems[i].Stat.ID)
			p.Difficulty = levelString[problems[i].Difficulty.Level-1]
			p.Language = "Golang" // TODO: Analyze code folder
			p.Acceptance = float64(problems[i].Stat.AC) / float64(problems[i].Stat.TotalSubmit) * 100
			u.ACproblems = append(u.ACproblems, p)
		}
	}
}

func SaveJSON() *leetCodeUser {
	var u leetCodeUser
	u.init()
	data, err := u.saveJSON()
	if err != nil {
		log.Fatal(err)
	}
	u.parseProblems(data)
	u.makeReadMe()
	return nil
	// u.parseJSON(data)
}
