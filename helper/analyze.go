package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	Easy   = 0
	Medium = 1
	Hard   = 2
)

type myproblem struct {
	problemID, difficulty int
	title, filename       string
}

func getFiles() []string {
	var files []string
	problemDir, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/../algorithm")
	if err != nil {
		log.Fatal(err)
	}
	fileinfo, err := ioutil.ReadDir(problemDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fileinfo {
		files = append(files, f.Name())
	}
	return files
}
func (p *myproblem) Init(filename string) {
	problemID, err := strconv.Atoi(strings.Split(filename, ".")[0])
	if err != nil {
		log.Fatal(err)
	}
	p.problemID = problemID
	p.filename = strings.Split(filename, ".")[1]
	p.title = strings.Join(strings.Split(strings.Title(p.filename), "-"), " ")
}
func ListAllProblems() {
	leetcodeInit()
	files := getFiles()
	var problems = make([]myproblem, len(files))
	for i := 0; i < len(files); i++ {
		problems[i].Init(files[i])
		fmt.Println(problems[i].problemID, problems[i].title)
	}
}
