package calc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	BYTE     = 1.0
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
	TERABYTE = 1024 * GIGABYTE
)

func ByteSize(bytes uint64) string {
	uint := ""
	value := float32(bytes)

	switch {
	case bytes >= TERABYTE:
		uint = "TB"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		uint = "GB"
		value = value / GIGABYTE
	case bytes >= KILOBYTE:
		uint = "KB"
		value = value / KILOBYTE
	case bytes >= BYTE:
		uint = "B"
	case bytes == 0:
		return "0"
	}
	stringValue := fmt.Sprintf("%.1f", value)
	stringValue = strings.TrimSuffix(stringValue, ".0")
	return fmt.Sprintf("%s%s", stringValue, uint)
}

func CalcSize(path string) [5]string {
	var max uint64 = 0
	var max_jpg string
	var min uint64 = 1<<63 - 1
	var min_jpg string
	var avg uint64 = 0

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		stat, err := os.Stat(path + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		if max < uint64(stat.Size()) {
			max = uint64(stat.Size())
			max_jpg = f.Name()
		}
		if min > uint64(stat.Size()) {
			min = uint64(stat.Size())
			min_jpg = f.Name()
		}
		avg += uint64(stat.Size())
	}
	var resp = [5]string{ByteSize(max), max_jpg, ByteSize(min), min_jpg, ByteSize(avg / 30)}
	return resp
}
