package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	apiURL = "https://leetcode.com/api/problems/algorithms/"
)

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
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: u.Connection.Csrftoken})
	req.AddCookie(&http.Cookie{Name: "LEETCODE_SESSION", Value: u.Connection.Session})
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
	json.Unmarshal(bodyBytes, &u)
	err = ioutil.WriteFile("leetcode.json", bodyBytes, 0644)
	if err != nil {
		return nil, err
	}
	problemBytes := []byte("[")
	rawproblemBytes := bytes.Split(bytes.Split(bodyBytes, []byte("["))[1], []byte("]"))[0]
	problemBytes = append(problemBytes, rawproblemBytes...)
	problemBytes = append(problemBytes, []byte("]")...)
	err = ioutil.WriteFile("problems.json", problemBytes, 0644)
	if err != nil {
		return nil, err
	}
	return problemBytes, nil
}
func (u *leetCodeUser) parseProblems(b []byte) {
	problems := []rawProblem{}
	json.Unmarshal(b, &problems)
	levelString := []string{"Easy", "Medium", "Hard"}
	tags := parseTags()
	for i := len(problems) - 1; i >= 0; i-- {
		if problems[i].Status == "ac" &&
			(u.Language == "all" || strings.Contains(tags[problems[i].Stat.ID], u.Language)) {
			var p problem
			p.Title = problems[i].Stat.Title
			p.NO = problems[i].Stat.ID
			p.Difficulty = levelString[problems[i].Difficulty.Level-1]
			p.Acceptance = float64(problems[i].Stat.AC) / float64(problems[i].Stat.TotalSubmit) * 100
			if u.Language == "all" {
				p.Language = parseLanguage(tags[p.NO])
			} else {
				p.Language = u.Language
			}
			p.tags = strings.Split(tags[p.NO], ",")
			u.ACproblems = append(u.ACproblems, p)
		}
	}
}
