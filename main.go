package main

import (
	"flag"
	"log"
	"os"

	"github.com/JinHyukParkk/resolution_tileSize_comparison/txt"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	site := flag.String("site", "naver", "Site Name")
	location := flag.String("ln", "독도", "Location Name")
	flag.Parse()
	log.Println("--------Start--------")
	txt.CreateTxt(*site, *location)
	log.Println("-------- End --------")

}
