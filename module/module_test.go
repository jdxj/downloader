package module

import (
	"fmt"
	"os"
	"sync"
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

func TestOpenFile(t *testing.T) {
	file1, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		t.Fatalf("file1 err: %s\n", err)
	}
	defer file1.Close()

	file2, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		t.Fatalf("file2 err: %s\n", err)
	}
	defer file2.Close()

	num := 1024 * 1024 * 20 // 20M

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		_, err := file1.Seek(0, 0)
		if err != nil {
			fmt.Printf("go file1, err: %s\n", err)
		}
		buf := symbol(65, num)
		_, err = file1.Write(buf)
		if err != nil {
			fmt.Printf("go file1 write, err: %s\n", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_, err := file2.Seek(int64(num), 0)
		if err != nil {
			fmt.Printf("go file2, err: %s\n", err)
		}
		buf := symbol(97, num)
		_, err = file2.Write(buf)
		if err != nil {
			fmt.Printf("go file2 write, err: %s\n", err)
		}
	}()

	wg.Wait()
	return
}

func symbol(sb byte, num int) []byte {
	res := make([]byte, num)
	for i := range res {
		res[i] = sb
	}
	return res
}

func TestOpenFile2(t *testing.T) {
	file1, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		t.Fatalf("file1 err: %s\n", err)
	}
	defer file1.Close()

	num := 1024 * 1024 * 20 // 20M

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		_, err := file1.Seek(0, 0)
		if err != nil {
			fmt.Printf("go file1, err: %s\n", err)
		}
		buf := symbol(65, num)
		_, err = file1.Write(buf)
		if err != nil {
			fmt.Printf("go file1 write, err: %s\n", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_, err := file1.Seek(int64(num), 0)
		if err != nil {
			fmt.Printf("go file1, err: %s\n", err)
		}
		buf := symbol(97, num)
		_, err = file1.Write(buf)
		if err != nil {
			fmt.Printf("go file1 write, err: %s\n", err)
		}
	}()

	wg.Wait()
	return
}

func TestOpenFile3(t *testing.T) {
	file1, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		t.Fatalf("file1 err: %s\n", err)
	}
	defer file1.Close()

	num := 1024 * 1024 * 20 // 20M

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		_, err := file1.Seek(0, 0)
		if err != nil {
			fmt.Printf("go file1, err: %s\n", err)
		}
		buf := symbol(65, num)
		_, err = file1.Write(buf)
		if err != nil {
			fmt.Printf("go file1 write, err: %s\n", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_, err := file1.Seek(int64(num), 0)
		if err != nil {
			fmt.Printf("go file1, err: %s\n", err)
		}
		buf := symbol(97, num)
		_, err = file1.Write(buf)
		if err != nil {
			fmt.Printf("go file1 write, err: %s\n", err)
		}
	}()

	wg.Wait()
	return
}
