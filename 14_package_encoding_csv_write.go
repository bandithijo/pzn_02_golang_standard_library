package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"rizqi", "nur", "assyaufi"})
	_ = writer.Write([]string{"pratama", "kedua", "ketiga"})
	_ = writer.Write([]string{"eka", "dwi", "tri"})

	writer.Flush()
	// rizqi,nur,assyaufi
	// pratama,kedua,ketiga
	// eka,dwi,tri
}
