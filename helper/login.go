package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/mozillazg/request"
)

const (
	loginURL  = "https://leetcode.com/accounts/login"
	userAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"
)

type leetCodeUser struct {
	config userConfig
}
type userConfig struct {
	username, password string
}

func (u *userConfig) get() {
	dat, _ := ioutil.ReadFile("config.yaml")
	u.username = strings.Split(string(dat), "\"")[1]
	u.password = strings.Split(string(dat), "\n")[1]
	u.password = strings.Split(u.password, "\"")[1]
}
func (u *leetCodeUser) init() {
	u.config.get()
}
func (u *leetCodeUser) login() error {
	req := request.NewRequest(new(http.Client))
	csrfToken, cfdUID := getCSRFToken(req)
	cookies := "__cfduid=" + cfdUID + 
				";csrftoken=" + csrfToken 
				+ ";"
	fmt.Println(cookies)
	req.Headers = map[string]string{
		"User-Agent": userAgent,
		"Connection": "keep-alive",
		"Referer":    "https://leetcode.com/accounts/login/",
		"origin":     "https://leetcode.com",
		// "Cookie":          cookies,
	}
	req.Data = map[string]string{
		"csrfmiddlewaretoken": csrfToken,
		"login":               u.config.username,
		"password":            u.config.password,
		"next":                "problems",
		// "__cfduid":            cfdUID,
	}
	// fmt.Printf("%s %s", u.config.username, u.config.password)
	response, err := req.Post(loginURL)
	defer response.Body.Close()
	fmt.Println(response.Status)
	return err
}
func getCSRFToken(r *request.Request) (string, string) {
	response, err := r.Get(loginURL)
	if err != nil {
		log.Panicf("Cannot get to %s: %s", response, err)
	}
	cookies := response.Cookies()
	var csrfToken, cfdUID string
	for _, cookie := range cookies {
		switch cookie.Name {
		case "csrftoken":
			csrfToken = cookie.Value
		case "__cfduid":
			cfdUID = cookie.Value
		}
	}
	if csrfToken != "" && cfdUID != "" {
		return csrfToken, cfdUID
	}
	panic("Can not find both csrftoken and cfduid in cookies!")
}

func LoginMain() {
	var u leetCodeUser
	u.init()
	u.login()
}
