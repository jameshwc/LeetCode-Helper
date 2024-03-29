package helper

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type trendCSV struct {
	trends []trend
}
type trend struct {
	date                      string
	total, easy, medium, hard int
}

func (t *trendCSV) write(u leetCodeUser) bool {
	var isModify = true
	var trendFileName string
	if u.Language != "all" {
		trendFileName = fmt.Sprintf("%s-trend.csv", u.Language)
	} else {
		trendFileName = "trend.csv"
	}
	csvfile, err := os.OpenFile(trendFileName, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Couldn't open the trend csv file", err)
	}
	defer csvfile.Close()
	r := csv.NewReader(csvfile)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal("Error when reading trend file", err)
	}
	if len(rows) > 2 {
		for _, row := range rows[1:] {
			var tr trend
			tr.date = row[0]
			tr.total, _ = strconv.Atoi(row[1])
			tr.easy, _ = strconv.Atoi(row[2])
			tr.medium, _ = strconv.Atoi(row[3])
			tr.hard, _ = strconv.Atoi(row[4])
			t.trends = append(t.trends, tr)
		}
		if t.trends[len(t.trends)-1].total >= len(u.ACproblems) {
			isModify = false
		}
	}
	if isModify {
		csvfile.Seek(0, io.SeekEnd)
		w := csv.NewWriter(csvfile)
		if len(rows) < 1 {
			w.Write([]string{"date", "total", "easy", "medium", "hard"})
		}
		var tr trend
		tr.date = time.Now().Format("06/01/02")
		parseProblems(&tr, u.ACproblems)
		t.trends = append(t.trends, tr)
		wstr := []string{tr.date, strconv.Itoa(tr.total), strconv.Itoa(tr.easy), strconv.Itoa(tr.medium), strconv.Itoa(tr.hard)}
		w.Write(wstr)
		w.Flush()
	}
	return isModify
}

func parseProblems(tr *trend, ACproblems []problem) {
	for p := range ACproblems {
		switch ACproblems[p].Difficulty {
		case "Easy":
			tr.easy++
		case "Medium":
			tr.medium++
		case "Hard":
			tr.hard++
		}
		tr.total++
	}
}
