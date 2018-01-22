package main

import (
	"flag"
	"log"

	"github.com/JinHyukParkk/resolution_tileSize_comparison/txt"
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
	log.Println("--------Start--------")
	txt.CreateTxt(*site, *location, "100")
	txt.CreateTxt(*site, *location, "90")
	txt.CreateTxt(*site, *location, "80")
	txt.CreateTxt(*site, *location, "70")
	log.Println("-------- End --------")

}
