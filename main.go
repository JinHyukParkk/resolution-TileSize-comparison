package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/JinHyukParkk/resolution_tileSize_comparison/calculate"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func main() {
	site := flag.String("site", "naver", "Site Name")
	location := flag.String("ln", "독도", "Location Name")
	flag.Parse()

	path := "./tileData/"
	files, err := ioutil.ReadDir(path)
	check(err)
	createFilePath := "./" + *site + "Result.txt"
	result, err := os.Create(createFilePath)
	check(err)
	defer result.Close()
	w := bufio.NewWriter(result)
	for _, f := range files {
		s := strings.Split(f.Name(), "_")
		nPath := path + f.Name() + "/" + *location
		if s[0] == *site {
			out := "###### " + s[0] + " " + *location + " Resolution" + " " + s[1] + "%\r\n"
			w.WriteString(out)
			resp := calc.CalcSize(nPath)
			out = "MaxSize : " + resp[0] + "\tFileName : " + resp[1] + "\r\n"
			w.WriteString(out)
			out = "MinSize : " + resp[2] + "\tFileName : " + resp[3] + "\r\n"
			w.WriteString(out)
			out = "AvgSize : " + resp[4] + "\r\n"
			w.WriteString(out)
		}
	}
	w.Flush()
}
