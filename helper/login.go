package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mozillazg/request"
)

type myRequest request.Request
type myRequestPtr *request.Request
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
func (u *leetCodeUser) login() myRequestPtr {
	req := request.NewRequest(new(http.Client))
	req.Headers = map[string]string{
		"Accept-Encoding": "",
		"Referer":         "https://leetcode.com/",
	}
	return myRequestPtr(req)
}
func (r *myRequest) getCSRFToken() string {
	return ""
}

func LoginMain() {
	var u leetCodeUser
	u.init()
	fmt.Println(u)
}
