package module

import (
	"fmt"
	"testing"
)

const (
	tmpURL = "https://nchc.dl.sourceforge.net/project/evolution-x/raphael/EvolutionX_4.4_raphael-10.0-20200602-1022-OFFICIAL.zip"
)

func TestFileSize(t *testing.T) {
	size, err := FileSize(tmpURL)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("size: %dB\n", size)
	fmt.Printf("size: %.3fGB\n", float64(size)/1024/1024/1024)
}
