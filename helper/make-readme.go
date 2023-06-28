package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const readmeFileName = "README.md"

func makeReadMe(u leetCodeUser, t trendCSV) {
	fullLanguageName := map[string]string{
		"js":     "javascript",
		"go":     "golang",
		"kotlin": "kotlin",
	}
	f, err := os.Create(readmeFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	sampleBytes, err := ioutil.ReadFile("README.md.sample")
	if err != nil {
		log.Fatal("open README.md.sample error", err)
	}
	f.Write(sampleBytes)
	for _, val := range u.ACproblems {
		var s string
		if u.Language != "all" {
			s = fmt.Sprintf("|%.4d|%s|%.2f%%|%s|%s\n", val.NO, val.Title, val.Acceptance, val.Difficulty, fullLanguageName[u.Language])
		} else {
			s = fmt.Sprintf("|%.4d|%s|%.2f%%|%s|%s\n", val.NO, val.Title, val.Acceptance, val.Difficulty, val.Language)
		}
		f.WriteString(s)
	}
	f.WriteString("\n|Date|total|easy|medium|hard\n") // TODO: Support multi language
	f.WriteString("|:-:|-:|-:|-:|-:|\n")
	for _, val := range t.trends {
		s := fmt.Sprintf("|%s|%2d|%2d|%2d|%2d\n", val.date, val.total, val.easy, val.medium, val.hard)
		f.WriteString(s)
		fmt.Println(s)
	}
	f.Sync()
}
