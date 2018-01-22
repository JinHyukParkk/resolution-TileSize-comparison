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
	t := flag.String("type", "1", "Vworld map 2D or 3D")
	flag.Parse()
	log.Println("--------Start--------")
	if *site != "1" && *location != "1" {
		txt.CreateTxt(*site, *location, *t)
	} else {
		log.Println("site와 ln Flag를 입력해주세요.")
	}
	test()
	log.Println("-------- End --------")
}

func test() {
	txt.CreateTxt("vworld", "강남역", "2D")
	txt.CreateTxt("vworld", "여의도", "2D")
	txt.CreateTxt("vworld", "독도", "2D")
	txt.CreateTxt("vworld", "북한산", "2D")
	txt.CreateTxt("vworld", "설악산", "2D")

	txt.CreateTxt("vworld", "강남역", "3D")
	txt.CreateTxt("vworld", "여의도", "3D")
	txt.CreateTxt("vworld", "독도", "3D")
	txt.CreateTxt("vworld", "북한산", "3D")
	txt.CreateTxt("vworld", "설악산", "3D")

	txt.CreateTxt("naver", "강남역", "nil")
	txt.CreateTxt("naver", "여의도", "nil")
	txt.CreateTxt("naver", "독도", "nil")
	txt.CreateTxt("naver", "북한산", "nil")
	txt.CreateTxt("naver", "설악산", "nil")

	txt.CreateTxt("daum", "강남역", "nil")
	txt.CreateTxt("daum", "여의도", "nil")
	txt.CreateTxt("daum", "독도", "nil")
	txt.CreateTxt("daum", "북한산", "nil")
	txt.CreateTxt("daum", "설악산", "nil")
}
