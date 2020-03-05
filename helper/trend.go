package helper

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const trendFileName = "trend.csv"

type trendCSV struct {
	trends []trend
}
type trend struct {
	date                      string
	total, easy, medium, hard int
}

func (t *trendCSV) write(u leetCodeUser) bool {
	var isModify = true
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
		if t.trends[len(t.trends)-1].total >= u.AC {
			isModify = false
		}
	}
	if isModify {
		csvfile.Seek(0, io.SeekEnd)
		w := csv.NewWriter(csvfile)
		if len(rows) < 1 {
			w.Write([]string{"date", "total", "easy", "medium", "hard"})
		}
		wstr := []string{time.Now().Format("06/01/02"), strconv.Itoa(u.AC), strconv.Itoa(u.ACeasy), strconv.Itoa(u.ACmedium), strconv.Itoa(u.AChard)}
		// TODO: refactor the following code, it has smell.
		var tr trend
		tr.date = time.Now().Format("06/01/02")
		tr.easy = u.ACeasy
		tr.medium = u.ACmedium
		tr.hard = u.AChard
		tr.total = u.AC
		t.trends = append(t.trends, tr)
		w.Write(wstr)
		w.Flush()
	}
	return isModify
}
