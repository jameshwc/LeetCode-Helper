package helper

import (
	"fmt"
	"log"
	"os"
)

const readmeFileName = "README.md"

func (u *leetCodeUser) makeReadMe() error {
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
	for i := range u.ACproblems {
		s := fmt.Sprintf("|%s|%s|%.2f%%|%s|%s\n", u.ACproblems[i].NO, u.ACproblems[i].Title, u.ACproblems[i].Acceptance, u.ACproblems[i].Difficulty, u.ACproblems[i].Language)
		f.WriteString(s)
	}
	return nil
}
