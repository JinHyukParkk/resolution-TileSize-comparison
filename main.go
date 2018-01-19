package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/JinHyukParkk/resolutionOfTiles_Comparison/calculate"
)

func main() {

	site := flag.String("site", "naver", "Site Name")
	location := flag.String("name", "독도", "Location Name")
	flag.Parse()

	path := "./tileData/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		s := strings.Split(f.Name(), "_")
		path += f.Name() + "/" + *location
		if s[0] == *site {
			fmt.Println("######", s[0], *location, "Resolution", s[1]+"%")
			resp := calc.CalcSize(path)
			fmt.Println("MaxSize :", resp[0]+"  FileName :", resp[1])
			fmt.Println("MinSize :", resp[2]+"  FileName :", resp[3])
			fmt.Println("AvgSize :", resp[4])
		}
	}
}
