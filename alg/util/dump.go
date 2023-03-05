package util

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func Dump(slice []int64, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data, err := json.Marshal(slice)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(data)

}

func Restore1dimSlice(filename string) []int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		var data []int64
		err := json.Unmarshal([]byte(text), &data)
		if err != nil {
			log.Fatal(err)
		} else {
			return data
		}
	}
	return nil
}
