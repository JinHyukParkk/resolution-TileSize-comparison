package calc

import (
	"fmt"
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
