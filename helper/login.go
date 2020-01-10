package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mozillazg/request"
)

type myRequest request.Request

type user struct {
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
func (u *user) init() {
	u.config.get()
}
func (u *user) login() *myRequest {
	req := request.NewRequest(new(http.Client))
	req.Headers = map[string]string{
		"Accept-Encoding": "",
		"Referer":         "https://leetcode.com/",
	}

}
func (r *myRequest) getCSRFToken() string {
	return nil
}

func main() {
	var u user
	u.init()
	fmt.Println(u)
}
