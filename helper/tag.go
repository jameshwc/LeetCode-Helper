package helper

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseTags() map[int]string {
	problemsPath := "algorithm/"
	var files []string
	tags := make(map[int]string)
	err := filepath.Walk(problemsPath, func(path string, info os.FileInfo, err error) error {
		if info.Name() == ".tag" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		dat, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		folder := filepath.Base(filepath.Dir(file))
		v, err := strconv.Atoi(folder[:4])
		if err != nil {
			panic(err)
		}
		tags[v] = string(dat)
	}
	return tags
}

func parseLanguage(tag string) string {
	var language []string
	tags := strings.Split(tag, ",")
	for t := range tags {
		switch tags[t] {
		case "golang", "kotlin", "javascript":
			language = append(language, tags[t])
		}
	}
	return strings.Join(language, ",")
}
