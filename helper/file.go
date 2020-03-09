package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func MakeFolder(id int, language string) {
	var problems []rawProblem
	extension := map[string]string{
		"golang":     ".go",
		"kotlin":     ".kt",
		"javascript": ".js",
	}
	dat, err := ioutil.ReadFile("problems.json")
	if err != nil {
		log.Fatal("Error when reading leetcode.json", err)
	}
	json.Unmarshal(dat, &problems)
	for i := len(problems) - 1; i >= 0; i-- {
		if problems[i].Stat.ID == id {
			rootpath, _ := os.Getwd()
			algoPath := filepath.Join(rootpath, "algorithm")
			folderName := fmt.Sprintf("%.4d.%s", id, problems[i].Stat.TitleSlug)
			idFolderPath := filepath.Join(algoPath, folderName)
			os.Mkdir(idFolderPath, 0755)
			file := filepath.Join(idFolderPath, problems[i].Stat.TitleSlug+extension[language])
			os.Create(file)
			tagfile := filepath.Join(idFolderPath, ".tag")
			if _, err := os.Stat(tagfile); err != nil {
				os.Create(tagfile)
			}
			break
		}
	}
}
