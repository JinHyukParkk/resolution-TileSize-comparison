package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/JinHyukParkk/resolutionOfTiles_Comparison/calculate"
)

func main() {

	site := flag.String("site", "naver", "Site Name")
	location := flag.String("name", "독도", "Location Name")
	flag.Parse()

	path := "./tileData/" + *site + "/" + *location
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var max uint64 = 0
	var min uint64 = 1<<63 - 1
	var avg uint64 = 0
	for _, f := range files {
		stat, err := os.Stat(path + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		if max < uint64(stat.Size()) {
			max = uint64(stat.Size())
		}
		if min > uint64(stat.Size()) {
			min = uint64(stat.Size())
		}
		avg += uint64(stat.Size())
		fmt.Println(f.Name(), calc.ByteSize(uint64(stat.Size())))
	}
	fmt.Println("max :", calc.ByteSize(max))
	fmt.Println("min :", calc.ByteSize(min))
	fmt.Println("avg :", calc.ByteSize(avg/30))
}
