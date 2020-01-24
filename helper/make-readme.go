package helper

import (
	"fmt"
	"log"
	"os"
)

const readmeFileName = "README.md"

func makeReadMe(u leetCodeUser, t trendCSV) {
	f, err := os.Create(readmeFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("LeetCode Ans\n")
	f.WriteString("===\n")
	f.WriteString("Currently only write in golang.\n")
	f.WriteString("Will support python, java, etc. in the future.\n")
	f.WriteString("\n## Status\n\n")
	f.WriteString("|Problem No.|Title|Acceptance|Difficulty|Language|\n")
	f.WriteString("|:-:|:-:|:-: | :-: | :-: |\n")
	for _, val := range u.ACproblems {
		s := fmt.Sprintf("|%.4d|%s|%.2f%%|%s|%s\n", val.NO, val.Title, val.Acceptance, val.Difficulty, val.Language)
		f.WriteString(s)
	}
	f.WriteString("\n|Date|total|easy|medium|hard\n") // TODO: Support multi language
	f.WriteString("|:-:|-:|-:|-:|-:|\n")
	for _, val := range t.trends {
		s := fmt.Sprintf("|%s|%2d|%2d|%2d|%2d\n", val.date, val.total, val.easy, val.medium, val.hard)
		f.WriteString(s)
	}
}
