package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Size() > 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		colnum, _ := strconv.ParseInt(os.Args[1], 0, 32)
		colFrom(string(bytes), colnum)
	} else {
		fmt.Println("| col [column-number]")
	}
}

func colFrom(in string, number int64) {
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, col := range records {
		fmt.Println(col[number])
	}

}

func intersect(sect []string, orig []string, tgt []string) []string {
	for _, i := range orig {
		for _, x := range tgt {
			if i == x {
				if !exists(x, sect) {
					sect = append(sect, x)
				}
				return intersect(sect, orig[1:], tgt[1:])
			}
		}
	}
	return sect
}

func exists(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
