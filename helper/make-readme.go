package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const readmeFileName = "README.md"

func makeReadMe(u leetCodeUser, t trendCSV) {
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
		language := strings.Join(val.Language[:], ",")
		s := fmt.Sprintf("|%.4d|%s|%.2f%%|%s|%s\n", val.NO, val.Title, val.Acceptance, val.Difficulty, language)
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
