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
	if site != nil && location != nil {
		txt.CreateTxt(*site, *location)
	} else {
		log.Println("site와 ln Flag를 입력해주세요.")
	}
	test()
	log.Println("-------- End --------")
}

func test() {
	txt.CreateTxt("vworld", "강남역")
	txt.CreateTxt("vworld", "여의도")
	txt.CreateTxt("vworld", "독도")
	txt.CreateTxt("vworld", "북한산")
	txt.CreateTxt("vworld", "설악산")
}
