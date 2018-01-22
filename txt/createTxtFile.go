package txt

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	calc "github.com/JinHyukParkk/resolution_tileSize_comparison/calculate"
)

type txtFile struct {
	file os.File
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("###### Create result directory ")
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func CreateTxt(site string, location string, t string) {
	path := "./tileData/"
	files, err := ioutil.ReadDir(path)
	check(err)
	CreateDirIfNotExist("./result")
	var createFilePath string
	if site == "vworld" {
		createFilePath = "./result/" + site + "_" + t + location + "_Result.txt"
	} else {
		createFilePath = "./result/" + site + "_" + location + "_Result.txt"
	}
	log.Println("###### Create txt -", createFilePath)
	result, err := os.Create(createFilePath)
	check(err)
	defer result.Close()
	w := bufio.NewWriter(result)
	for _, f := range files {
		s := strings.Split(f.Name(), "_")
		if s[0] == site {
			if s[0] == "vworld" {
				lo := t + location
				nPath := path + f.Name() + "/" + lo
				Calc(w, site, s, nPath, lo)
			} else {
				nPath := path + f.Name() + "/" + location
				Calc(w, site, s, nPath, location)
			}
		}
	}
	w.Flush()
	log.Println("###### Complete")
}
func Calc(w *bufio.Writer, site string, s []string, nPath string, location string) {
	out := "###### " + s[0] + " " + location + " Resolution" + " " + s[1] + "%\r\n"
	w.WriteString(out)
	resp := calc.CalcSize(nPath)
	out = fmt.Sprintf("MaxSize : %-6s\tFileName : %-13s\r\n", resp[0], resp[1])
	w.WriteString(out)
	out = fmt.Sprintf("MinSize : %-6s\tFileName : %-13s\r\n", resp[2], resp[3])
	w.WriteString(out)
	out = fmt.Sprintf("AvgSize : %-6s\r\n", resp[4])
	w.WriteString(out)
}
