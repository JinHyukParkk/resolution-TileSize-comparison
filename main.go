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
	site := flag.String("site", "1", "Site Name")
	location := flag.String("ln", "1", "Location Name")
	flag.Parse()
	log.Println("--------Start--------")
	if *site != "1" && *location != "1" {
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

	txt.CreateTxt("naver", "강남역")
	txt.CreateTxt("naver", "여의도")
	txt.CreateTxt("naver", "독도")
	txt.CreateTxt("naver", "북한산")
	txt.CreateTxt("naver", "설악산")

	txt.CreateTxt("daum", "강남역")
	txt.CreateTxt("daum", "여의도")
	txt.CreateTxt("daum", "독도")
	txt.CreateTxt("daum", "북한산")
	txt.CreateTxt("daum", "설악산")
}
